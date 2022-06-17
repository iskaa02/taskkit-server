package types

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type Subtask struct {
	IsCompleted bool   `json:"is_completed"`
	Name        string `json:"name"`
}
type Subtasks map[string]Subtask

func (s Subtasks) Value() (driver.Value, error) {
	return json.Marshal(s)
}

func (s *Subtasks) Scan(src interface{}) error {
	return scan(src, s)
}

func scan(src interface{}, f interface{}) error {
	var source []byte
	switch t := src.(type) {
	case string:
		source = []byte(t)
	case []byte:
		source = t
	case nil:
		return nil
	default:
		return fmt.Errorf("incompatible data type %T", src)
	}
	if len(source) == 0 {
		return nil
	}
	return json.Unmarshal(source, f)
}
