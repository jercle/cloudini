package datatransforms

import (
	"fmt"
	"log"
	"os"

	"golang.design/x/clipboard"
)

// var (
// 	abbreviations             = flag.String("abbreviations", "", "comma-separated list of extra abbreviations")
// 	format                    = flag.String("format", "json", "format (json or yaml)")
// 	uncompress                = flag.Bool("z", false, "decompress input with gzip")
// 	omitempty                 = flag.String("omitempty", "auto", "generate omitempty (never, always, or auto)")
// 	packageComment            = flag.String("packagecomment", "", "package comment")
// 	packageName               = flag.String("packagename", "main", "package name")
// 	skipUnparseableProperties = flag.Bool("skipunparseableproperties", true, "skip unparseable properties")
// 	structTagName             = flag.String("structtagname", "", "struct tag name")
// 	typeComment               = flag.String("typecomment", "", "type comment")
// 	typeName                  = flag.String("typename", "T", "type name")
// 	intType                   = flag.String("inttype", "", "integer type")
// 	useJSONNumber             = flag.Bool("usejsonnumber", false, "use json.Number")
// 	goFormat                  = flag.Bool("goformat", true, "format generated Go code")
// 	output                    = flag.String("o", "", "output filename")

// 	omitEmptyOption = map[string]json2struct.OmitEmptyOption{
// 		"never":  json2struct.OmitEmptyNever,
// 		"always": json2struct.OmitEmptyAlways,
// 		"auto":   json2struct.OmitEmptyAuto,
// 	}
// )

type GenerateOptions struct {
	JsonString       string
	JsonFile         string
	WebApi           string
	Stdin            bool
	GetFromClipboard bool
}

func Generate(opts GenerateOptions) (string, error) {
	// err := clipboard.Init()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	os.Exit(0)
	// }
	generator := NewGenerator()

	// if flag.NArg() == 0 {
	// var input io.Reader = os.Stdin

	if opts.Stdin {
		generator.ObserveJSONReader(os.Stdin)
	} else if opts.JsonFile != "" {
		// fmt.Println(opts.JsonFile)
		generator.ObserveJSONFile(opts.JsonFile)
	} else if opts.JsonString != "" {
		generator.ObserveJSONString(opts.JsonString)
	} else if opts.GetFromClipboard {
		err := clipboard.Init()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}
		generator.ObserveJSONString(string(clipboard.Read(clipboard.FmtText)))
		// } else if opts.WebApi != "" {
		// TODO, implement getting and parsing direct from an API
		// }

		// fmt.Println(cb)
	}

	goCode, err := generator.Generate()
	if err != nil {
		log.Fatal(err)
	}

	return string(goCode), err

}
