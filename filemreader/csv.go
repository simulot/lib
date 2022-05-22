package filemapper

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

func ReadCSV(datafile string, hasHeader bool, fieldSep rune) ([]Sheet, error) {
	if !hasHeader {
		return nil, fmt.Errorf("can't manage CSV files without header row")
	}

	f, err := os.Open(datafile)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)

	csvReader.Comma = fieldSep
	csvReader.TrimLeadingSpace = true

	fields, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}
	data := []Sheet{}
	sheet := Sheet{
		Name:    filepath.Base(datafile),
		Fields:  map[string]int{},
		Records: []Record{},
	}

	if len(fields) == 0 {
		return data, nil
	}

	for r := range fields {
		if r == 0 {
			for f := range fields[r] {
				sheet.Fields[fields[r][f]] = f
			}
			continue
		}

		fs := []string{}
		for f := range fields[r] {
			fs = append(fs, fields[r][f])
		}
		sheet.Records = append(sheet.Records, Record{
			sheet:  &sheet,
			Fields: fs,
		})
	}
	data = append(data, sheet)
	return data, nil
}
