package robinhood

import (
	"fmt"
	"strings"
	"time"
)

type date struct {
	time.Time
}

const layout = "2006-01-02"

func (d *date) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		d.Time = time.Time{}
		return nil
	}
	var err error
	d.Time, err = time.Parse(layout, s)
	return err
}

func (d *date) MarshalJSON() ([]byte, error) {
	if d.Time.UnixNano() == (time.Time{}).UnixNano() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", d.Time.Format(layout))), nil
}
