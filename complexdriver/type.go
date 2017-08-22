package complexdriver

import (
	"fmt"
	"strconv"
)

// NullPercent records the nil information in awql
type NullPercent struct {
	Valid   bool
	Percent float64
}

// Scan reads the value
func (n *NullPercent) Scan(src interface{}) error {
	if s, ok := src.(string); ok {
		if s == " --" {
			return nil
		}
		if s == "0.00%" {
			n.Valid = true
			n.Percent = 0
			return nil
		}

		if s[len(s)-1] != '%' {
			return fmt.Errorf("unknown value %q", s)
		}
		var err error
		n.Percent, err = strconv.ParseFloat(s[:len(s)-1], 64)
		if err != nil {
			return err
		}
		n.Valid = true
		return nil
	}
	return fmt.Errorf("unknow value %q", src)
}
