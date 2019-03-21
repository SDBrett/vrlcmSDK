package vrlcmSdk

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

/*
// TODO Add result parsing and return statement
func (s *SdkConnection) GetDataCenters() error {

	t := func(c *SdkConnection) {
		c.Client.Transport = NewDefaultSdkTransport(s.IgnoreCertError)
	}

	c, _ := NewApiClient(t)

	url := s.BaseUrl + "/view/datacenter"

	req, err := http.NewRequest("GET", url, nil)
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

}

*/
