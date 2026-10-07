package lib

import (
	"errors"
	"fmt"
	"time"
)

const mailboxReportLayout = "2006-01-02"

type TimeMailboxReport struct {
	time.Time
}

func (ct *TimeMailboxReport) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		return nil
	}
	if len(b) < 2 || b[0] != '"' || b[len(b)-1] != '"' {
		return errors.New("invalid JSON string format for custom time")
	}
	str := string(b[1 : len(b)-1])

	t, err := time.Parse(mailboxReportLayout, str)
	if err != nil {
		return fmt.Errorf("failed to parse custom time: %w", err)
	}

	ct.Time = t
	return nil
}

func (cd TimeMailboxReport) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, cd.Format(mailboxReportLayout))), nil
}

// func (cd TimeMailboxReport) UnmarshalBSON() ([]byte, error) {
// 	// return []byte(fmt.Sprintf(`"%s"`, cd.Format(mailboxReportLayout))), nil
// 	    return bson.RawValue()
// }
