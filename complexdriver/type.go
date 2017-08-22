package complexdriver

import (
	"fmt"
	"strconv"

	"github.com/davecgh/go-spew/spew"
)

type NullPercent struct {
	Valid   bool
	Percent float64
}

func (n *NullPercent) Scan(src interface{}) error {
	spew.Dump(src)
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
