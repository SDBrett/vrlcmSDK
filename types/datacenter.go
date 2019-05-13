package types

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
	Vcenters  []Vcenter `json:"vCenters, omitempty"`
}
