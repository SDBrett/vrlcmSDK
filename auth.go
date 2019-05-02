package vrlcmsdk

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
// Adds auth token string to the ApiClient
func (c *ApiClient) Login(u, p string) error {

	url := c.basePath + "/login"
	body := CreateLoginRequestBody(u, p)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	response, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	err = ValidateHttpResponse(*response)
	if err != nil {
		return err
	}

	c.token, err = getAuthToken(response)
	if err != nil {
		return err
	}

	c.addAuthHeader()

	return nil
}

// Performs logout action against vRLCM server
func (c *ApiClient) Logout() error {

	url := c.basePath + "/logout"
	req, err := http.NewRequest("POST", url, nil)

	req.Header = *c.headers
	if err != nil {
		return err
	}

	response, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	err = ValidateHttpResponse(*response)
	if err != nil {
		return err
	}

	return nil
}
