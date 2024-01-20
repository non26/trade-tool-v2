package okxhelper

import (
	"net/http"
	"strings"
	"time"
	"tradetoolv2/config"
)

func GenerateHeaders(
	req *http.Request,
	method string,
	requestPath string,
	body string,
	env string,
	secrets *config.Secrets,
) *http.Request {
	t := time.Now()
	req.Header.Add("OK-ACCESS-KEY", GenerateKeyHeader(secrets.OkxApiKey))
	req.Header.Add("OK-ACCESS-SIGN", GenerateSignHeader(t, method, requestPath, body, secrets.OkxSecretKey))
	req.Header.Add("OK-ACCESS-TIMESTAMP", GenerateTimeHeader(t))
	req.Header.Add("OK-ACCESS-PASSPHRASE", GeneratePassPhaseHeader(secrets.OkxSecretPassPhase))
	switch method {
	case http.MethodPost:
		req.Header.Add("Content-Type", "application/json")
	default:
		break
	}
	if strings.ToUpper(env) == "LOCAL" {
		req.Header.Add("x-simulated-trading", "1")
	}
	return req
}
