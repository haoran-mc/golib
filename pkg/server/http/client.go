package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

func SetClientTimeout(timeout time.Duration) {
	myClient.Timeout = timeout
}

func ResetClientTimeout() {
	myClient.Timeout = 10 * time.Second
}

func doHttp(url string, method string, body io.Reader) (resp *http.Response, err error) {
	r, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	r.Header.Set("Content-Type", "application/json")
	return myClient.Do(r)
}

func GetJson(url string, target any) error {
	r, err := doHttp(url, http.MethodGet, nil)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}

func PostJson(url string, target any) error {
	var jsonValue []byte
	switch t := target.(type) {
	case string:
		jsonValue = []byte(t)
	default:
		jsonValue, _ = json.Marshal(target)
	}
	resp, err := doHttp(url, http.MethodPost, bytes.NewBuffer(jsonValue))
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}
	defer resp.Body.Close()
	return nil
}
