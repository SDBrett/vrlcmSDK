package vrlcmSdk

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// This struct is used for marshalling the response body returned from the login API call
type LoginResponse struct {
	//Token holds the returned token code
	Token string `json:"token"`
}

// Creates the request body used for the login API call
func CreateLoginRequestBody(u, p string) []byte {
	bodyString := `{"username":"` + u + `", "password":"` + p + `"}`
	return []byte(bodyString)
}

//
func getAuthToken(r *http.Response) (string, error) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", err
	}

	loginResponse := LoginResponse{}

	err = json.Unmarshal(body, &loginResponse)

	if err != nil {
		return "", err
	}

	return loginResponse.Token, nil

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
	s.Token, err = getAuthToken(response)
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
