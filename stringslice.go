package flagutil

import (
	"strings"
)

// StringSliceFlag is a flag type which support comma separated values mapping string slices
type StringSliceFlag []string

func (s *StringSliceFlag) String() string {
	return strings.Join(*s, ",")
}

func (s *StringSliceFlag) Set(value string) error {
	*s = strings.Split(value, ",")
	return nil
}
