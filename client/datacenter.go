package client

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Datacenters struct {
	Datacenter []Datacenter
}
type Datacenter struct {
	ID       string   `json:"id"`
	Name     string   `json:"name",json:"datacenterName"`
	City     string   `json:"city, omitempty"`
	State    string   `json:"state, omitempty"`
	Country  string   `json:"country, omitempty"`
	Vcenters []string `json:"vcenters, omitempty"`
}

func getDataCentersResponse(r *http.Response) ([]Datacenter, error) {

	// Parse response body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	// Marshall response body into loginResponse struct
	d := &[]Datacenter{}
	err = json.Unmarshal(body, &d)
	if err != nil {
		return nil, err
	}

	return *d, nil

}

func (d *Datacenters) GetDatacenters(s SdkConnection) error {

	url := s.BaseUrl + "/view/datacenter"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header = s.headers

	response, err := s.Client.Do(req)
	if err != nil {
		return err
	}

	err = ValidateHttpResponse(*response)
	if err != nil {
		return err
	}
	d.Datacenter, err = getDataCentersResponse(response)
	if err != nil {
		return err
	}

	return nil
}
