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

func (c *ApiClient) setDefaultHeaders() {

	h := &http.Header{}
	h.Add("Accept", "application/json")
	h.Add("Content-Type", "application/json")

	c.headers = h
}

func (c *ApiClient) addAuthHeader() {

	c.headers.Add("x-xenon-auth-token", c.token)

}
