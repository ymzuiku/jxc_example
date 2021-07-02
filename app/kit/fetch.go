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

	res, _ := ioutil.ReadAll(resp.Body)

	out := map[string]interface{}{}

	err = json.Unmarshal(res, &out)

	if err != nil {
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
	res, _ := ioutil.ReadAll(resp.Body)
	var out map[string]interface{}
	err = json.Unmarshal(res, &out)
	if err != nil {
		out["error"] = string(res)
	}
	return out
}
