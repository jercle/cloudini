package lib

func UniqueNonEmptyElementsOf(s []string) []string {
	unique := make(map[string]bool, len(s))
	us := make([]string, len(unique))
	for _, elem := range s {
		if len(elem) != 0 {
			if !unique[elem] {
				us = append(us, elem)
				unique[elem] = true
			}
		}
	}
	return us
}

// func (tables *LogAnalyticsTables) printTable() {
// 	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
// 	columnFmt := color.New(color.FgYellow).SprintfFunc()
// 	tbl := table.New("Name", "Retention Period", "isDefault")
// 	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
// 	for _, record := range *tables {
// 		retentionPeriod := strconv.Itoa(record.Properties.RetentionInDays)
// 		var isDefault string
// 		if record.Properties.RetentionInDaysAsDefault {
// 			isDefault = "True"
// 		} else {
// 			isDefault = "False"
// 		}
// 		tbl.AddRow(record.Name, retentionPeriod, isDefault)
// 	}
// 	tbl.Print()
// }
