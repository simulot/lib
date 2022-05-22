package myflag

import (
	"regexp"
)

type Regexp struct {
	*regexp.Regexp
}

func (re *Regexp) Set(s string) error {
	r, err := regexp.Compile(s)
	if err != nil {
		return err
	}

	(*re).Regexp = r
	return nil
}

func (re Regexp) String() string {
	if re.Regexp == nil {
		return ""
	}
	return re.Regexp.String()
}
