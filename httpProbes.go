package main

import (
	"io/ioutil"
	"net/http"
)

func httpGetRequest(url string) (string, error) {
	// Simple HTTP Request which returns the Body as String or an Error
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
