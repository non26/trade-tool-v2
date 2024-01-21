package helper

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"
	"time"
)

func GenerateKeyHeader(k string) (key string) {
	return key
}

/*
Sign the prehash string with the SecretKey using the HMAC SHA256.
*/
func GenerateSignHeader(tStp time.Time, method string, requestPath string, body string, sk string) (signBase64 string) {
	queryString := fmt.Sprintf("%v?%v", requestPath, body)
	t := GenerateTimeHeader(tStp)
	s := fmt.Sprintf("%v%v%v%v", t, strings.ToUpper(method), queryString, sk)
	signHmacSHA256 := hmac.New(sha256.New, []byte(s))
	signBase64 = base64.RawURLEncoding.EncodeToString(signHmacSHA256.Sum(nil))
	return signBase64
}

/*
time show as UTC
*/
func GenerateTimeHeader(tStp time.Time) (timeStamp string) {
	tUtc := tStp.UTC()
	return tUtc.String()
}

/*
The passphrase you specified when creating the APIKey.
*/
func GeneratePassPhaseHeader(pp string) (passPhase string) {
	return passPhase
}

/*
`x-simulated-trading: 1` needs to be added to the header of the Demo Trading request.
*/
func GenerateDemoHeader() (v string) {
	return fmt.Sprintf("%v", 1)
}
