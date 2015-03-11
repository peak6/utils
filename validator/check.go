package validator

import (
	"time"
)

type Check struct {
	err error
}

func NewCheck() *Check {
	return &Check{}
}

func (c *Check) GetError() error {
	return c.err
}

func (c *Check) Clear() {
	c.err = nil
}

func (c *Check) CheckErr(err error) {
	if c.err != nil {
		return
	}

	if err != nil {
		c.err = err
	}
}

func (c *Check) RequiredString(val string, err error) {
	if c.err != nil {
		return
	}

	if requiredString(val) {
		c.err = err
	}
}

func (c *Check) RequiredInt(val int, err error) {
	if c.err != nil {
		return
	}

	if requiredInt(val) {
		c.err = err
	}
}

func (c *Check) RequiredFloat64(val float64, err error) {
	if c.err != nil {
		return
	}

	if requiredFloat64(val) {
		c.err = err
	}
}

func (c *Check) RequiredBool(val bool, err error) {
	if c.err != nil {
		return
	}

	if requiredBool(val) {
		c.err = err
	}
}

func (c *Check) NotNil(val interface{}, err error) {
	if c.err != nil {
		return
	}

	if notNil(val) {
		c.err = err
	}
}

func (c *Check) RequiredTime(val time.Time, err error) {
	if c.err != nil {
		return
	}

	if requiredTime(val) {
		c.err = err
	}
}

func (c *Check) MinInt(val int, n int, err error) {
	if c.err != nil {
		return
	}

	if minInt(val, n) {
		c.err = err
	}
}

func (c *Check) MaxInt(val int, n int, err error) {
	if c.err != nil {
		return
	}

	if maxInt(val, n) {
		c.err = err
	}
}

func (c *Check) MinFloat64(val float64, n float64, err error) {
	if c.err != nil {
		return
	}

	if minFloat64(val, n) {
		c.err = err
	}
}

func (c *Check) MaxFloat64(val float64, n float64, err error) {
	if c.err != nil {
		return
	}

	if maxFloat64(val, n) {
		c.err = err
	}
}

func (c *Check) MinChar(val string, n int, err error) {
	if c.err != nil {
		return
	}

	if minChar(val, n) {
		c.err = err
	}
}

func (c *Check) MaxChar(val string, n int, err error) {
	if c.err != nil {
		return
	}

	if maxChar(val, n) {
		c.err = err
	}
}

func (c *Check) Email(val string, err error) {
	if c.err != nil {
		return
	}

	if email(val) {
		c.err = err
	}
}

func (c *Check) Gender(val string, err error) {
	if c.err != nil {
		return
	}

	if gender(val) {
		c.err = err
	}
}

func (c *Check) Confirm(val, confirm string, err error) {
	if c.err != nil {
		return
	}

	if confirmVals(val, confirm) {
		c.err = err
	}
}

func (c *Check) ISO8601DataTime(val string, err error) {
	if c.err != nil {
		return
	}

	if iso8601DataTime(val) {
		c.err = err
	}
}

func (c *Check) InString(val string, in []string, err error) {
	if c.err != nil {
		return
	}

	if inString(val, in) {
		c.err = err
	}
}
