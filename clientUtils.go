package vrlcmsdk

import (
	"net/http"
)

func (cli *ApiClient) SetDefaultHeaders() {

	h := &http.Header{}
	h.Add("Accept", "application/json")
	h.Add("Content-Type", "application/json")

	cli.headers = h
}

func (cli *ApiClient) AddAuthHeader() {

	cli.headers.Add("x-xenon-auth-Token", cli.Token)

}
