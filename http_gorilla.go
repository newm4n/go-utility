package go_utility

import (
	"errors"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func GetHttpPathParam(r *http.Request, name string) (string, error) {
	vars := mux.Vars(r)
	val, ok := vars[name]
	if !ok {
		return "", errors.New("no path param key")
	}
	return val, nil
}

func GetHttpRequestParam(r *http.Request, name string) ([]string, error) {
	vals, ok := r.URL.Query()[name]
	if !ok || len(vals[0]) < 1 {
		return nil, errors.New("request key not found")
	}
	ret := make([]string, 0)
	for _, v := range vals {
		ret = append(ret, string(v))
	}
	return ret, nil
}

func GetHttpRequestBody(r *http.Request) (string, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", err
	} else {
		return string(body), nil
	}
}

func IsHttpContentJson(r *http.Request) bool {
	contentType := r.Header.Get("Content-Type")
	if len(contentType) == 0 {
		return false
	}
	if contentType != "application/json" {
		return false
	}
	return true
}
