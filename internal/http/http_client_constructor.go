package http

import (
	"net/http"
)

func NewHttpClient() *http.Client { // Wire's Provider
	return &http.Client{}
}
