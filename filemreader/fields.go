package filemapper

import (
	"github.com/xuri/excelize/v2"

	"fmt"
	"strings"
	"time"
)

type WorkBook []Sheet

type Sheet struct {
	Name    string
	Header  []string
	Fields  map[string]int
	Records []Record
}

type Record struct {
	sheet  *Sheet
	Fields []string
}

func (wb WorkBook) Records() []Record {
	r := []Record{}
	for _, s := range wb {
		r = append(r, s.Records...)
	}
	return r
}

func (wb WorkBook) Sheet(name string) (Sheet, error) {
	n := strings.TrimSpace(strings.ToLower(name))
	for _, s := range wb {
		if strings.TrimSpace(strings.ToLower(s.Name)) == n {
			return s, nil
		}
	}
	return Sheet{}, fmt.Errorf("can't find sheet '%s'", name)
}

func (r Record) FieldName(idx int) string {
	if idx > len(r.sheet.Header) {
		return ""
	}
	return r.sheet.Header[idx]
}

func (r Record) Field(name string) (string, error) {
	var err error
	i, ok := r.sheet.Fields[name]
	if !ok {
		// Maybe the this is the column name
		i, err = excelize.ColumnNameToNumber(name)
		if err != nil {
			return "", fmt.Errorf("unknown field '%s' in sheet '%s'", name, r.sheet.Name)
		}
		i--
	}
	if i > len(r.Fields)-1 {
		return "", nil
	}
	return r.Fields[i], nil
}
func (r Record) FieldOrEmpty(name string) string {
	s, _ := r.Field(name)
	return s
}

func TimeParse(layout string, f string) (time.Time, error) {
	t, err := time.Parse(layout, f)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}
