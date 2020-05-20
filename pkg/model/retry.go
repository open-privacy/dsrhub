package model

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/spf13/cast"
)

type Retry struct {
	Count int `yaml:"count"`
	Delay int `yaml:"delay"` // delay in seconds
}

func (r *Retry) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	s := cast.ToString(value)
	if err := json.Unmarshal([]byte(s), r); err != nil {
		return err
	}
	return nil
}

func (r Retry) Value() (driver.Value, error) {
	bytes, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}
	return string(bytes), nil
}
