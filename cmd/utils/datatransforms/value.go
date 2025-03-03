package datatransforms

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/fatih/structtag"
	"golang.org/x/exp/maps"
)

// FIXME extract sub-structs

// An value describes an observed value.
type value struct {
	observations        int
	emptys              int
	arrays              int
	bools               int
	float64s            int
	ints                int
	nulls               int
	objects             int
	strings             int
	times               int // time.Time is an implicit more specific type than string.
	arrayElements       *value
	allObjectProperties *value
	objectProperties    map[string]*value
}

type generateOptions struct {
	exportNameFunc            ExportNameFunc
	imports                   map[string]struct{}
	intType                   string
	omitEmptyOption           OmitEmptyOption
	skipUnparseableProperties bool
	structTagNames            []string
	useJSONNumber             bool
}

// observe merges a into v.
func (v *value) observe(a any) *value {
	if v == nil {
		v = &value{}
	}
	v.observations++
	switch a := a.(type) {
	case []any:
		v.arrays++
		if len(a) == 0 {
			v.emptys++
		}
		if v.arrayElements == nil {
			v.arrayElements = &value{}
		}
		for _, e := range a {
			v.arrayElements = v.arrayElements.observe(e)
		}
	case bool:
		v.bools++
		if !a {
			v.emptys++
		}
	case float64:
		v.float64s++
		if a == 0 {
			v.emptys++
		}
	case int:
		v.ints++
		if a == 0 {
			v.emptys++
		}
	case nil:
		v.nulls++
	case map[string]any:
		v.objects++
		if len(a) == 0 {
			v.emptys++
		}
		if v.objectProperties == nil {
			v.objectProperties = make(map[string]*value)
		}
		for property, value := range a {
			v.allObjectProperties = v.allObjectProperties.observe(value)
			v.objectProperties[property] = v.objectProperties[property].observe(value)
		}
	case string:
		if a == "" {
			v.emptys++
		}
		if v.times == v.strings {
			if _, err := time.Parse(time.RFC3339Nano, a); err == nil {
				v.times++
			}
		}
		v.strings++
	case json.Number:
		if _, err := a.Int64(); err == nil {
			v.ints++
		} else {
			v.float64s++
		}
	}
	return v
}

// goType returns the Go type of v.
func (v *value) goType(observations int, options *generateOptions) (string, bool) {
	// Determine the number of distinct types observed.
	distinctTypes := 0
	if v.arrays > 0 {
		distinctTypes++
	}
	if v.bools > 0 {
		distinctTypes++
	}
	if v.float64s > 0 {
		distinctTypes++
	}
	if v.ints > 0 {
		distinctTypes++
	}
	if v.nulls > 0 {
		distinctTypes++
	}
	if v.objects > 0 {
		distinctTypes++
	}
	if v.strings > 0 {
		distinctTypes++
	}

	// Based on the observed distinct types, find the most specific Go type.
	switch {
	case distinctTypes == 1 && v.arrays > 0:
		fallthrough
	case distinctTypes == 2 && v.arrays > 0 && v.nulls > 0:
		elementGoType, _ := v.arrayElements.goType(0, options)
		return "[]" + elementGoType, v.arrays+v.nulls < observations && v.emptys == 0
	case distinctTypes == 1 && v.bools > 0:
		return "bool", v.bools < observations && v.emptys == 0
	case distinctTypes == 2 && v.bools > 0 && v.nulls > 0:
		return "*bool", false
	case distinctTypes == 1 && v.float64s > 0:
		return "float64", v.float64s < observations && v.emptys == 0
	case distinctTypes == 2 && v.float64s > 0 && v.nulls > 0:
		return "*float64", false
	case distinctTypes == 1 && v.ints > 0:
		return options.intType, v.ints < observations && v.emptys == 0
	case distinctTypes == 2 && v.ints > 0 && v.nulls > 0:
		return "*" + options.intType, false
	case distinctTypes == 2 && v.float64s > 0 && v.ints > 0:
		omitEmpty := v.float64s+v.ints < observations && v.emptys == 0
		if options.useJSONNumber {
			options.imports["encoding/json"] = struct{}{}
			return "json.Number", omitEmpty
		}
		return "float64", omitEmpty
	case distinctTypes == 3 && v.float64s > 0 && v.ints > 0 && v.nulls > 0:
		if options.useJSONNumber {
			options.imports["encoding/json"] = struct{}{}
			return "*json.Number", false
		}
		return "*float64", false
	case distinctTypes == 1 && v.objects > 0:
		fallthrough
	case distinctTypes == 2 && v.objects > 0 && v.nulls > 0:
		if len(v.objectProperties) == 0 {
			switch {
			case observations == 0 && v.nulls == 0:
				return "struct{}", false
			case v.nulls > 0:
				return "*struct{}", false
			case v.objects == observations:
				return "struct{}", false
			default:
				return "*struct{}", v.objects < observations
			}
		}
		hasUnparseableProperties := false
		for k := range v.objectProperties {
			if strings.ContainsRune(k, ' ') {
				hasUnparseableProperties = true
				break
			}
		}
		if hasUnparseableProperties && !options.skipUnparseableProperties {
			valueGoType, _ := v.allObjectProperties.goType(0, options)
			return "map[string]" + valueGoType, v.objects+v.nulls < observations
		}
		b := &bytes.Buffer{}
		properties := maps.Keys(v.objectProperties)
		sort.Strings(properties)
		fmt.Fprintf(b, "struct {\n")
		var unparseableProperties []string
		for _, property := range properties {
			if isUnparseableProperty(property) {
				unparseableProperties = append(unparseableProperties, property)
				continue
			}
			goType, observedEmpty := v.objectProperties[property].goType(v.objects, options)
			var omitEmpty bool
			switch {
			case options.omitEmptyOption == OmitEmptyNever:
				omitEmpty = false
			case options.omitEmptyOption == OmitEmptyAlways:
				omitEmpty = true
			case options.omitEmptyOption == OmitEmptyAuto:
				omitEmpty = observedEmpty
			}

			tags, _ := structtag.Parse("")
			var structTagOptions []string
			if omitEmpty {
				structTagOptions = append(structTagOptions, "omitempty")
			}
			for _, structTagName := range options.structTagNames {
				tag := &structtag.Tag{
					Key:     structTagName,
					Name:    property,
					Options: structTagOptions,
				}
				_ = tags.Set(tag)
			}

			fmt.Fprintf(b, "%s %s `%s`\n", options.exportNameFunc(property), goType, tags)
		}
		for _, property := range unparseableProperties {
			fmt.Fprintf(b, "// %q cannot be unmarshalled into a struct field by encoding/json.\n", property)
		}
		fmt.Fprintf(b, "}")
		switch {
		case observations == 0:
			return b.String(), false
		case v.objects == observations:
			return b.String(), false
		case v.objects < observations && v.nulls == 0:
			return "*" + b.String(), true
		default:
			return "*" + b.String(), v.objects+v.nulls < observations
		}
	case distinctTypes == 1 && v.strings > 0 && v.times == v.strings:
		options.imports["time"] = struct{}{}
		return "time.Time", v.times < observations
	case distinctTypes == 1 && v.strings > 0:
		return "string", v.strings < observations && v.emptys == 0
	case distinctTypes == 2 && v.strings > 0 && v.nulls > 0 && v.times == v.strings:
		options.imports["time"] = struct{}{}
		return "*time.Time", false
	case distinctTypes == 2 && v.strings > 0 && v.nulls > 0:
		return "*string", false
	default:
		return "any", v.arrays+v.bools+v.float64s+v.ints+v.nulls+v.objects+v.strings < observations
	}
}
