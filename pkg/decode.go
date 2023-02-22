package pkg

import (
	"encoding/json"

	_ "github.com/masudur-rahman/go-oneliners"

	"github.com/graphql-go/graphql"
)

func ParseInto(src any, dst any) error {
	jsonByte, err := json.Marshal(src)
	if err != nil {
		return err
	}
	return json.Unmarshal(jsonByte, dst)
}

func ParseGraphQLData(src *graphql.Result, dst any, key string) error {
	jsonByte, err := json.Marshal(src.Data)
	if err != nil {
		return err
	}

	var data map[string]json.RawMessage
	if err = json.Unmarshal(jsonByte, &data); err != nil {
		return err
	}

	return json.Unmarshal(data[key], dst)
}
