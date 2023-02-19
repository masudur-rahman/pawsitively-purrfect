package pkg

import (
	"encoding/json"

	_ "github.com/masudur-rahman/go-oneliners"
)

func ParseInto(src any, dst any) error {
	jsonByte, err := json.Marshal(src)
	if err != nil {
		return err
	}
	return json.Unmarshal(jsonByte, dst)
}
