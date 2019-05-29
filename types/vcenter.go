package types

type Vcenter struct {
	Name                string `json:"vCenterName"`
	Username            string `json:"userName"`
	Password            string `json:"password"`
	DatacenterName      string `json:"datacenterName"`
	Type                int    `json:"type"`
	Datacenters         string `json:"vCDataCenteres"`
	TemplateCustomSpecs string `json:"templateCustomSpecs"`
}


