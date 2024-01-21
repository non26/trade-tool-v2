package helper

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func ToQueryParameter(s interface{}) string {
	// structType := reflect.TypeOf(s)
	// valueOfStruct := reflect.ValueOf(s)
	// queryParam := ""
	// numberOfField := structType.NumField()

	// for idx := 0; idx < numberOfField; idx++ {
	// 	tag := structType.Field(idx).Tag.Get("json")
	// 	structFieldValue := valueOfStruct.Field(idx)
	// 	if idx == numberOfField-1 {
	// 		queryParam += fmt.Sprintf("%v=%v", tag, structFieldValue)
	// 		break
	// 	}
	// 	queryParam += fmt.Sprintf("%v=%v&", tag, structFieldValue)

	// }
	// return queryParam
	b, _ := json.Marshal(s)
	return string(b)
}

func StructToJson(s interface{}) (*bytes.Reader, error) {
	j, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(j), nil
}

func AddInstIdUSDTSWAPPostfix(instId string) string {
	return fmt.Sprintf("%v-USDT-SWAP", instId)
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
