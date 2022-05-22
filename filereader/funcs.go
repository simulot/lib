package filemapper

import (
	"regexp"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/google/uuid"
	"github.com/xuri/excelize/v2"
)

var FuncsMap = template.FuncMap{
	"SQLText":         SQLText,
	"SysGuid":         SysGuid,
	"Coalesce":        Coalesce,
	"ValidEmail":      ValidEmail,
	"Nvl":             Nvl,
	"CleanString":     CleanString,
	"CDATA":           CData,
	"TrimSpace":       strings.TrimSpace,
	"TimeParse":       time.Parse,
	"ExcelDateToTime": ExcelDateToTime,
}

func SysGuid() string {
	return strings.Replace(uuid.NewString(), "-", "", -1)
}
func SQLText(s string) string {
	b := strings.Builder{}

	wasDash := false
	for _, c := range s {
		switch c {
		case '\'':
			b.WriteString("''")
			wasDash = false
		case '-':
			if wasDash {
				b.WriteString("'||'-")
				wasDash = false
				continue
			}
			b.WriteRune('-')
			wasDash = true
		default:
			wasDash = false
			b.WriteRune(c)
		}
	}
	return b.String()
}

func Coalesce(s ...string) string {
	for i := range s {
		if len(s[i]) > 0 {
			return s[i]
		}
	}
	return ""
}

var reEmailValid = regexp.MustCompile(`[^@]+@[^\.]+\..+`)

func ValidEmail(s string) string {
	if !reEmailValid.MatchString(s) {
		return ""
	}
	return s
}

func Nvl(s, null string) string {
	if s == "" {
		return null
	}
	return s
}

var reDeQuotes = regexp.MustCompile("''+")

func CleanString(s string) string {
	b := strings.Builder{}
	b.WriteString(reDeQuotes.ReplaceAllString(s, "'"))
	return b.String()
}

func CData(s string) string {
	b := strings.Builder{}
	b.WriteString("<![CDATA[")
	b.WriteString(s)
	b.WriteString("]]>")
	return b.String()
}

func ExcelDateToTime(excelDate string, format string) (string, error) {
	f, err := strconv.ParseFloat(excelDate, 64)
	if err != nil {
		return "", err
	}
	t, err := excelize.ExcelDateToTime(f, false)
	if err != nil {
		return "", err
	}
	s := t.Format(format)
	return s, nil
}
