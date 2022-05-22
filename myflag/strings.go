package myflag

import "strings"

type Strings []string

func (sl *Strings) Set(s string) error {
	l := strings.Split(s, ",")
	(*sl) = append((*sl), l...)
	return nil
}

func (sl Strings) String() string {
	return strings.Join(sl, ",")
}

type StringMaps map[string]interface{}

func (sm *StringMaps) Set(s string) error {
	if *sm == nil {
		*sm = make(StringMaps)
	}
	if strings.Contains(s, ",") {
		ss := strings.Split(s, ",")
		for _, s := range ss {
			(*sm)[strings.TrimSpace(s)] = nil
		}
	} else {
		(*sm)[strings.TrimSpace(s)] = nil
	}
	return nil
}

func (sm StringMaps) String() string {
	s := []string{}
	for k := range sm {
		s = append(s, k)
	}
	return strings.Join(s, ",")
}
