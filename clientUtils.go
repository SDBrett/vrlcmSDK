package vrlcmsdk

import (
	"net/http"
)

func (cli *ApiClient) setDefaultHeaders() {

	h := &http.Header{}
	h.Add("Accept", "application/json")
	h.Add("Content-Type", "application/json")

	cli.headers = h
}

func (cli *ApiClient) addAuthHeader() {

	cli.headers.Add("x-xenon-auth-token", cli.token)

}
