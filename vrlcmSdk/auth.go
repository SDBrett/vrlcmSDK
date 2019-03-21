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

// Retrieves auth token from login API call response
func getAuthToken(r *http.Response) (string, error) {

	// Parse response body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", err
	}

	// Marshall response body into loginResponse struct
	loginResponse := LoginResponse{}
	err = json.Unmarshal(body, &loginResponse)
	if err != nil {
		return "", err
	}

	// Return token code
	return loginResponse.Token, nil

}

// Performs authentication function with vRLCM server
// Adds auth token string to the SdkConnection
func (s *SdkConnection) Login(u, p string) error {

	// Setup http transport using default transport
	t := func(c *SdkConnection) {
		c.Client.Transport = NewDefaultSdkTransport(s.IgnoreCertError)
	}

	// Create new http client
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

	// Create new headers for SDK connection which will contain the auth token
	s.newSdkHeaders()

	return nil
}

// Performs logout action against vRLCM server
func (s *SdkConnection) Logout() error {

	// Setup http transport using default transport
	t := func(c *SdkConnection) {
		c.Client.Transport = NewDefaultSdkTransport(s.IgnoreCertError)
	}

	// Create new http client
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
