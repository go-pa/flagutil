package flagutil


import (
	"encoding/json"
)

type JSONStringMapFlag map[string]string

func (s *JSONStringMapFlag) String() string {
	data, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func (s *JSONStringMapFlag) Set(value string) error {
	var m map[string]string
	err := json.Unmarshal([]byte(value), &m)
	if err != nil {
		return err
	}
	*s = m
	return nil
}
