package client

import (
	"net/http"
)

type SdkConnection struct {
	BaseUrl         string
	Token           string
	IgnoreCertError bool
	Client          http.Client
	headers         http.Header
}
