package okxhelper

import "github.com/google/go-querystring/query"

func ToQueryParameter(s interface{}) string {
	v, _ := query.Values(s)
	return v.Encode()
}
