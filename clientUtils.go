package vrlcmsdk

import (
	"errors"
	"io/ioutil"
	"net/http"
)

func ValidateHttpResponse(r http.Response) error {

	if r.StatusCode >= 300 {
		b, _ := ioutil.ReadAll(r.Body)
		err := errors.New(string(b))
		return err
	}
	return nil
}

func (cli *ApiClient) setDefaultHeaders() {

	h := &http.Header{}
	h.Add("Accept", "application/json")
	h.Add("Content-Type", "application/json")

	cli.headers = h
}

func (cli *ApiClient) addAuthHeader() {

	cli.headers.Add("x-xenon-auth-token", cli.token)

}
