package util

import (
	"net/http"
	"time"
)

var client *http.Client

func GetClient() *http.Client {
	if client == nil {
		tr := &http.Transport{
			MaxIdleConns:       3,
			IdleConnTimeout:    30 * time.Second,
			DisableCompression: true,
		}

		client = &http.Client{Transport: tr}
	}
	return client
}
