package vrlcmSdk

import (
	"errors"
	"io/ioutil"
	"net/http"
)

func ValidateHttpResponse(r http.Response) error {

	if r.StatusCode != 200 {
		b, _ := ioutil.ReadAll(r.Body)
		err := errors.New(string(b))
		return err
	}
	return nil
}
