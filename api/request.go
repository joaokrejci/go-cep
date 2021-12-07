package api

import (
	"io/ioutil"
	"net/http"
)

// Request makes an http request and returns a json object
func Request(url string) (json []byte, err error) {
	response, requestError := http.Get(url)
	if requestError != nil {
		err = requestError
		return
	}

	body, bodyError := ioutil.ReadAll(response.Body)
	if bodyError != nil {
		err = bodyError
		return
	}

	json = body

	defer response.Body.Close()

	return
}