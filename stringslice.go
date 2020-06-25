package flagutil

import (
	"strings"
)

// CommaStringSliceFlag is a flag type which support comma separated values mapping string slices
type CommaStringSliceFlag []string

func (s *CommaStringSliceFlag) String() string {

	return strings.Join(*s, ",")
}

func (s *CommaStringSliceFlag) Set(value string) error {
	*s = strings.Split(value, ",")
	return nil
}

// RepeatingStringSliceFlag is a flag type that requiress repeating the flag to add more entries.
type RepeatingStringSliceFlag []string

func (s *RepeatingStringSliceFlag) String() string {
	return strings.Join(*s, " AND ")
}

func (s *RepeatingStringSliceFlag) Set(value string) error {
	*s = append(*s, value)
	return nil
}

// StringSliceFlag is an alias of CommaStringSliceFlag to not break current usage.
// TODO(thomasf): remove before v1.0
type StringSliceFlag = CommaStringSliceFlag
