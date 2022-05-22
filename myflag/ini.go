package myflag

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type IniValues map[string]string

var reSplitINI = regexp.MustCompile(`^([^=]*)(?: *= *)(.*)$`)

func readIniFile(name string) (IniValues, error) {
	ini := IniValues{}
	f, err := os.Open(name)
	if err != nil {
		return nil, fmt.Errorf("can't read ini file '%s': %w", name, err)
	}

	defer f.Close()
	s := bufio.NewScanner(f)
	for s.Scan() {
		m := reSplitINI.FindStringSubmatch(s.Text())
		if len(m) < 3 {
			continue
		}
		k, v := strings.TrimSpace(m[1]), strings.TrimSpace(m[2])
		if strings.HasPrefix(v, `"`) && strings.HasSuffix(v, `"`) {
			v = v[1 : len(v)-1]
		}
		if !strings.HasPrefix(k, "#") {
			ini[k] = v
		}
	}
	if s.Err() != nil {
		return nil, fmt.Errorf("can't read ini file '%s': %w", name, err)
	}
	return ini, nil
}

func Parse(exeName string) error {
	iniFile := filepath.Join(filepath.Dir(exeName), strings.TrimSuffix(filepath.Base(exeName), filepath.Ext(exeName))+".ini")
	ini, err := readIniFile(iniFile)
	if err != nil {
		if !errors.Is(err, fs.ErrNotExist) {
			return err
		}
	}
	if !flag.Parsed() {
		flag.Parse()
	}

	setFlags := map[string]interface{}{}

	flag.Visit(func(f *flag.Flag) {
		setFlags[f.Name] = nil
	})

	for k, v := range ini {
		if _, isSet := setFlags[k]; !isSet {
			err = flag.Set(k, v)
			if err != nil {
				return fmt.Errorf("can't set flag '%s' read from ini file '%s': %w", k, iniFile, err)
			}
		}
	}
	return nil
}
