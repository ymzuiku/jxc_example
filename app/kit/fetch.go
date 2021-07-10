package kit

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/google/go-querystring/query"
)

func BaseLocalUrl() string {
	return "http://127.0.0.1:3100"
}

var BaseURL = ""

func Get(url string, params interface{}) map[string]interface{} {
	v, err := query.Values(params)
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := http.Get(BaseURL + url + "?" + v.Encode())
	if err != nil {
		return map[string]interface{}{"error": err.Error()}
	}
	defer resp.Body.Close()

	out := map[string]interface{}{}

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		out["error"] = string(res)
		return out
	}

	if err := json.Unmarshal(res, &out); err != nil {
		out["error"] = string(res)
	}

	return out
}

func Post(url string, body interface{}) map[string]interface{} {
	_body, err := json.Marshal(body)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post(BaseURL+url, "application/json", strings.NewReader((string(_body))))
	if err != nil {
		return map[string]interface{}{"error": err.Error()}
	}

	defer resp.Body.Close()
	var out map[string]interface{}

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		out["error"] = string(res)
	}

	if err := json.Unmarshal(res, &out); err != nil {
		out["error"] = string(res)
	}

	return out
}
