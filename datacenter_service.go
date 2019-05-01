package vrlcmsdk

import (
	"github.com/sdbrett/vrlcmsdk/datacenter"
	"net/http"
)

type DatacenterAPIService service

func (dc *DatacenterAPIService) GetAllDatacenters() (*datacenter.Datacenters, error) {

	url := dc.client.basePath + "/view/datacenter"
	d := datacenter.Datacenters{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header = *dc.client.headers
	r, err := dc.client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	d.Datacenter, err = datacenter.GetDatacentersResponse(r)
	if err != nil {
		return nil, err
	}

	for k := range d.Datacenter {

		id := d.Datacenter[k].ID
		d.Datacenter[k], err = dc.GetDatacenter(id)
		if err != nil {
			return nil, err
		}

		d.Datacenter[k].ID = id

	}

	return &d, nil

}

func (dc *DatacenterAPIService) GetDatacenter(id string) (datacenter.Datacenter, error) {

	url := dc.client.basePath + "/view/datacenter?datacenterId=" + id
	d := datacenter.Datacenter{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return d, err
	}

	req.Header = *dc.client.headers

	r, err := dc.client.httpClient.Do(req)
	if err != nil {
		return d, err
	}

	//var d *datacenter.Datacenter
	d, err = datacenter.GetDatacenterResponse(r)
	if err != nil {
		return d, err
	}

	return d, nil

}