package datacenter

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Datacenters struct {
	Datacenter []Datacenter
}
type Datacenter struct {
	ID        string `json:"id"`
	Name      string `json:"datacenterName, omitempty"`
	City      string `json:"city, omitempty"`
	Country   string `json:"country, omitempty"`
	Latitude  string `json:"latitude, omitempty"`
	Longitude string `json:"longitude, omitempty"`
	State     string `json:"state, omitempty"`
	Vcenters  string `json:"vcenters, omitempty"`
}

func GetDatacentersResponse(r *http.Response) ([]Datacenter, error) {

	d := &[]Datacenter{}
	// Parse response body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	// Marshall response body into loginResponse struct

	err = json.Unmarshal(body, &d)
	if err != nil {
		return nil, err
	}

	return *d, nil
}

func GetDatacenterResponse(r *http.Response) (Datacenter, error) {

	d := &Datacenter{}
	// Parse response body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return *d, err
	}

	// Marshall response body into loginResponse struct

	err = json.Unmarshal(body, &d)
	if err != nil {
		return *d, err
	}

	return *d, nil

}
