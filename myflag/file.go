package myflag

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type FileName string

func (fn *FileName) Set(s string) error {
	_, err := os.Stat(s)
	if err != nil {
		return fmt.Errorf("cant't check parameter: %w", err)
	}
	*fn = FileName(s)
	return nil
}

func (fn FileName) String() string {
	return string(fn)
}

type GlobFiles map[string]interface{}

func (rgf *GlobFiles) Set(s string) error {
	if *rgf == nil {
		*rgf = make(GlobFiles)
	}
	m, err := filepath.Glob(s)
	if err != nil {
		return err
	}

	for _, f := range m {
		(*rgf)[f] = nil
	}
	return nil
}

func (rgf GlobFiles) Files() []string {
	ss := []string{}
	for k := range rgf {
		ss = append(ss, k)
	}
	sort.Strings(ss)
	return ss
}

func (rgf GlobFiles) String() string {
	ss := rgf.Files()
	return strings.Join(ss, ",")
}

type DirName string

func (fn *DirName) Set(s string) error {
	_, err := os.Stat(s)
	if err != nil {
		return fmt.Errorf("cant't check parameter: %w", err)
	}
	*fn = DirName(s)
	return nil
}

func (fn DirName) String() string {
	return string(fn)
}
