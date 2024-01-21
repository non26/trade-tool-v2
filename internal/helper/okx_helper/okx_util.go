package helper

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

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

func AddInstIdUSDTPostfix(instId string) string {
	return fmt.Sprintf("%v-USDT", instId)
}

func OkxConditionResponseError(httpCode int, okxCode string, okxMsg string) error {
	msg := fmt.Sprintf("http=%v&code=%v&msg=%v", httpCode, okxCode, okxMsg)
	if httpCode != http.StatusOK {
		return errors.New(msg)
	}
	if okxCode == "1" {
		return errors.New(msg)
	}
	return nil
}
