package filemapper

import (
	"fmt"
	"strings"

	"github.com/xuri/excelize/v2"
)

func ReadExcel(datafile string) (WorkBook, error) {
	wb, err := excelize.OpenFile(datafile)
	if err != nil {
		return nil, fmt.Errorf("can't read excel file '%s': %w", datafile, err)
	}

	data := []Sheet{}
	for _, sheetName := range wb.GetSheetList() {
		sheet := Sheet{
			Name:    sheetName,
			Fields:  map[string]int{},
			Records: []Record{},
		}

		cells, err := wb.GetRows(sheetName, excelize.Options{RawCellValue: false})
		if err != nil {
			return nil, err
		}

		for r := range cells {
			if len(sheet.Fields) == 0 {
				for c, title := range cells[r] {
					title = strings.TrimSpace(title)
					sheet.Fields[title] = c
					sheet.Header = append(sheet.Header, title)
				}
				continue
			}

			fs := []string{}
			for f := range cells[r] {
				fs = append(fs, cells[r][f])
			}
			sheet.Records = append(sheet.Records, Record{
				sheet:  &sheet,
				Fields: fs,
			})

		}
		if len(sheet.Records) > 0 {
			data = append(data, sheet)
		}
	}
	return data, nil
}
