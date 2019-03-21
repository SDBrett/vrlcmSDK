package vrlcmSdk

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type LoginResponse struct {
	Token string `json:"token"`
}

func CreateLoginRequestBody(u, p string) []byte {
	bodyString := `{"username":"` + u + `", "password":"` + p + `"}`
	return []byte(bodyString)
}

func getLoginResponse(r *http.Response, s *SdkConnection) error {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	loginResponse := LoginResponse{}

	err = json.Unmarshal(body, &loginResponse)

	if err != nil {
		return err
	}

	s.Token = loginResponse.Token

	return nil

}

func (s *SdkConnection) Login(u, p string) error {

	t := func(c *SdkConnection) {
		c.Client.Transport = NewDefaultSdkTransport(s.IgnoreCertError)
	}

	c, _ := NewApiClient(t)

	url := s.BaseUrl + "/login"

	body := CreateLoginRequestBody(u, p)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	response, err := c.Client.Do(req)
	if err != nil {
		return err
	}

	err = ValidateHttpResponse(*response)
	if err != nil {
		return err
	}
	err = getLoginResponse(response, s)
	if err != nil {
		return err
	}

	s.newSdkHeaders()

	return nil
}

func (s *SdkConnection) Logout() error {

	t := func(c *SdkConnection) {
		c.Client.Transport = NewDefaultSdkTransport(s.IgnoreCertError)
	}

	c, _ := NewApiClient(t)

	url := s.BaseUrl + "/logout"

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}
	req.Header = s.headers

	response, err := c.Client.Do(req)
	if err != nil {
		return err
	}

	err = ValidateHttpResponse(*response)
	if err != nil {
		return err
	}

	return nil
}
