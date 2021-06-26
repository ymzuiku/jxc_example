package kit

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func Get(url string) map[string]interface{} {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	res, _ := ioutil.ReadAll(resp.Body)
	var out map[string]interface{}
	_ = json.Unmarshal(res, &out)
	return out
}

func Post(url string, body map[string]interface{}) map[string]interface{} {
	_body, err := json.Marshal(body)
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := http.Post(url, "application/json", strings.NewReader(("heel=" + string(_body))))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	res, _ := ioutil.ReadAll(resp.Body)
	var out map[string]interface{}
	_ = json.Unmarshal(res, &out)
	return out
}
