package okxhelper

import (
	"bytes"
	"encoding/json"

	"github.com/google/go-querystring/query"
)

func ToQueryParameter(s interface{}) string {
	v, _ := query.Values(s)
	return v.Encode()
}

func StructToJson(s interface{}) (*bytes.Reader, error) {
	j, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(j), nil
}

func JsonTostruct()
