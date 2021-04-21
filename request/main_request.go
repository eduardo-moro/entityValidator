package request

import (
	"bytes"
	"crypto/tls"
	"net/http"
	"strings"
)

func GetCfnByCode(codigo string) (string, string) {
	return makeRequest("https://cnn.cfn.org.br/application/front-resource/get",
		"{\"comando\":\"get-nutricionista\",\"options\":{\"geral\":true, \"registro\": \""+codigo+"\"}}")
}


func makeRequest(url string, input string) (string, string) {

	tr := &http.Transport{
	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
}
	client := &http.Client{Transport: tr}

	body := strings.NewReader(input)

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		// handle err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()


	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	respBody := buf.String()


	return resp.Status, respBody
}
