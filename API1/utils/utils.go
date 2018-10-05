package utils

import (
	"io/ioutil"
	"net/http"
)

func GetJsonBody(request *http.Request) []byte {
	bytes := request.Body
	data, _ := ioutil.ReadAll(bytes)
	return data
}
