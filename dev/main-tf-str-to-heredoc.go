package main

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/jercle/cloudini/lib"
)

func main() {
	filePath := "/home/jercle/git/evan-tooling/terraform/apc/dtq/sentinel_rules/azurerm_sentinel_alert_rule_scheduled.tf.bak"
	// filePath := "/home/jercle/git/evan-tooling/terraform/apc/prod/sentinel_rules/azurerm_sentinel_alert_rule_scheduled.tf.bak"
	// filePath := "/home/jercle/git/cloudini/dev/file.tf"
	// filePath := "/home/jercle/git/cloudini/dev/file-rpl.tf"
	// filePath := "/home/jercle/git/cloudini/dev/file-rpl.tf"
	tmpFilePath := "/home/jercle/git/cloudini/dev/file-rpl.tf"
	// tmpFilePath := "/home/jercle/git/cloudini/dev/file-rpl2.tf"
	f, err := os.Open(filePath)
	lib.CheckFatalError(err)
	defer f.Close()

	// scanner := bufio.NewScanner(f)
	// re := regexp.MustCompile(`"(?:[^"\\]|\\.)*"`)

	// tmp, err := os.CreateTemp("/home/jercle/git/cloudini/dev", "replace-*")
	if _, err := os.Stat(tmpFilePath); err == nil {
		// fmt.Println("Deleting temp file")
		os.Remove(tmpFilePath)
	}
	tmp, err := os.Create(tmpFilePath)
	lib.CheckFatalError(err)
	defer tmp.Close()

	// replace while copying from f to tmp
	if err := replace(f, tmp); err != nil {
		log.Fatal(err)
	}

	// ruleDescriptions := GetAllSentinelRuleDescriptions("/home/jercle/git/cloudini/dev/prod.json")
	// _ = ruleDescriptions
	// jsonStr, _ := json.MarshalIndent(ruleDescriptions, "", "  ")
	// fmt.Println(string(jsonStr))

	// fmt.Println(ruleDescriptions["Email access via active sync"])
	// for key, _ := range ruleDescriptions {
	// 	fmt.Println(key)
	// }

	// make sure the tmp file was successfully written to
	if err := tmp.Close(); err != nil {
		log.Fatal(err)
	}

	// close the file we're reading from
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func replace(r io.Reader, w io.Writer) error {
	re := regexp.MustCompile(`"(?:[^"\\]|\\.)*"`)

	ruleDescriptions := GetAllSentinelRuleDescriptions("/home/jercle/git/cloudini/dev/prod.json")
	_ = ruleDescriptions

	// use scanner to read line by line
	scanner := bufio.NewScanner(r)
	prevLine := ""
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		// quoted := strconv.Quote(line)
		if prevLine == "" && line == "" {
		} else if strings.Contains(line, "  query               ") {
			leadingWhitespaceCount := len(line) - len(strings.TrimLeft(line, " "))
			leadingWhitespace := strings.Repeat(" ", leadingWhitespaceCount)
			start := leadingWhitespace + "query = <<QUERY\n"
			end := "\nQUERY\n"

			newStr := re.FindString(line)
			newStr, err := strconv.Unquote(newStr)
			lib.CheckFatalError(err)
			_, err = io.WriteString(w, start+newStr+end)
			lib.CheckFatalError(err)
		} else if strings.Contains(line, "  description    ") {
		} else if strings.Contains(line, "# Please review these resources and move them into your main configuration files.") {
		} else if strings.Contains(line, "# __generated__ by OpenTofu") {
			// fmt.Println("found")
			// fmt.Println(prevLine)
			if prevLine != "" {
				_, err := io.WriteString(w, "\n")
				_, err = io.WriteString(w, "\n")
				_, err = io.WriteString(w, "\n")
				_, err = io.WriteString(w, "\n")
				lib.CheckFatalError(err)
			}
		} else if strings.Contains(line, "  display_name  ") {
			leadingWhitespaceCount := len(line) - len(strings.TrimLeft(line, " "))
			leadingWhitespace := strings.Repeat(" ", leadingWhitespaceCount)

			displayName := re.FindString(line)
			displayName, err := strconv.Unquote(displayName)
			lib.CheckFatalError(err)
			description := ruleDescriptions[displayName]
			quotedDescription := strconv.Quote(description)
			newDescription := leadingWhitespace + "description = " + quotedDescription
			_, err = io.WriteString(w, line+"\n")
			lib.CheckFatalError(err)
			if description != "" {
				_, err = io.WriteString(w, newDescription+"\n")
				lib.CheckFatalError(err)
			}
			prevLine = line
		} else {
			_, err := io.WriteString(w, line+"\n")
			lib.CheckFatalError(err)
			if line != "" {
				prevLine = line
			}
		}
		// fmt.Println(strconv.Quote(prevLine))
		// os.Exit(0)
	}

	// fmt.Println(ruleDescriptions)
	return scanner.Err()
}

func GetAllSentinelRuleDescriptions(filePath string) map[string]string {
	var sentinelRules []SentinelAlertRule
	f, err := os.ReadFile(filePath)
	lib.CheckFatalError(err)
	json.Unmarshal(f, &sentinelRules)
	// fmt.Println(sentinelRules)
	descriptionsMap := make(map[string]string)

	for _, rule := range sentinelRules {
		descriptionsMap[rule.Properties.DisplayName] = rule.Properties.Description
	}

	return descriptionsMap
}

type SentinelAlertRule struct {
	Etag       string `json:"etag"`
	ID         string `json:"id"`
	Kind       string `json:"kind"`
	Name       string `json:"name"`
	Properties struct {
		AlertDetailsOverride *struct {
			AlertDescriptionFormat string `json:"alertDescriptionFormat,omitempty"`
			AlertDisplayNameFormat string `json:"alertDisplayNameFormat,omitempty"`
			AlertDynamicProperties []any  `json:"alertDynamicProperties"`
		} `json:"alertDetailsOverride,omitempty"`
		AlertRuleTemplateName *string `json:"alertRuleTemplateName"`
		CustomDetails         *struct {
			DnsQueries        string `json:"DNSQueries,omitempty"`
			DnsQueryCount     string `json:"DNSQueryCount,omitempty"`
			DnsQueryThreshold string `json:"DNSQueryThreshold,omitempty"`
			DnsQuerythreshold string `json:"DNSQuerythreshold,omitempty"`
			NxdomainCount     string `json:"NXDOMAINCount,omitempty"`
			NxdomaiNthreshold string `json:"NXDOMAINthreshold,omitempty"`
			ProcessName       string `json:"ProcessName,omitempty"`
			SubjectUserName   string `json:"SubjectUserName,omitempty"`
			TimeEnabled       string `json:"TimeEnabled,omitempty"`
		} `json:"customDetails,omitempty"`
		Description               string `json:"description"`
		DisplayName               string `json:"displayName"`
		DisplayNamesExcludeFilter any    `json:"displayNamesExcludeFilter,omitempty"`
		DisplayNamesFilter        any    `json:"displayNamesFilter,omitempty"`
		Enabled                   bool   `json:"enabled"`
		EntityMappings            []struct {
			EntityType    string `json:"entityType"`
			FieldMappings []struct {
				ColumnName string `json:"columnName"`
				Identifier string `json:"identifier"`
			} `json:"fieldMappings"`
		} `json:"entityMappings,omitempty"`
		EventGroupingSettings *struct {
			AggregationKind string `json:"aggregationKind"`
		} `json:"eventGroupingSettings,omitempty"`
		IncidentConfiguration *struct {
			CreateIncident        bool `json:"createIncident"`
			GroupingConfiguration struct {
				Enabled              bool     `json:"enabled"`
				GroupByAlertDetails  []string `json:"groupByAlertDetails"`
				GroupByCustomDetails []any    `json:"groupByCustomDetails"`
				GroupByEntities      []string `json:"groupByEntities"`
				LookbackDuration     string   `json:"lookbackDuration"`
				MatchingMethod       string   `json:"matchingMethod"`
				ReopenClosedIncident bool     `json:"reopenClosedIncident"`
			} `json:"groupingConfiguration"`
		} `json:"incidentConfiguration,omitempty"`
		LastModifiedUtc     time.Time `json:"lastModifiedUtc"`
		ProductFilter       string    `json:"productFilter,omitempty"`
		Query               string    `json:"query,omitempty"`
		QueryFrequency      string    `json:"queryFrequency,omitempty"`
		QueryPeriod         string    `json:"queryPeriod,omitempty"`
		SeveritiesFilter    []string  `json:"severitiesFilter,omitempty"`
		Severity            string    `json:"severity,omitempty"`
		SuppressionDuration string    `json:"suppressionDuration,omitempty"`
		SuppressionEnabled  bool      `json:"suppressionEnabled"`
		Tactics             []string  `json:"tactics"`
		Techniques          []string  `json:"techniques"`
		TemplateVersion     string    `json:"templateVersion,omitempty"`
		TriggerOperator     string    `json:"triggerOperator,omitempty"`
		TriggerThreshold    float64   `json:"triggerThreshold"`
	} `json:"properties"`
	Type string `json:"type"`
}
