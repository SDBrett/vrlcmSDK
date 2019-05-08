package vrlcmsdk

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/sdbrett/vrlcmsdk/datacenter"
	"github.com/sdbrett/vrlcmsdk/types"
	"net/http"
)

type DatacenterAPIService service

func (dc *DatacenterAPIService) GetAllDatacenters(ctx context.Context) (types.Datacenters, error) {

	url := dc.client.basePath + "/view/datacenter"
	d := types.Datacenters{}

	resp, err := dc.client.get(ctx, url, *dc.client.headers)

	if err != nil {
		return d, err
	}

	err = json.NewDecoder(resp.body).Decode(&d.Datacenter)
	ensureReaderClosed(resp)
	return d, nil

}

func (dc *DatacenterAPIService) GetDatacenter(ctx context.Context, id string) (types.Datacenter, error) {

	url := dc.client.basePath + "/view/datacenter?datacenterId=" + id
	d := types.Datacenter{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return d, err
	}

	req.Header = *dc.client.headers
	r, err := dc.client.httpClient.Do(req)
	if err != nil {
		return d, err
	}
	defer r.Body.Close()

	d, err = datacenter.GetDatacenterResponse(r)
	if err != nil {
		return d, err
	}

	return d, nil

}

func (dc *DatacenterAPIService) Create(ctx context.Context, d *types.Datacenter) error {

	url := dc.client.basePath + "/action/create/datacenter"
	var tempDC types.Datacenter
	if d.Name == "" {
		err := errors.New("Datacenter name cannot be empty")
		return err
	}

	resp, err := dc.client.post(ctx, url, d, *dc.client.headers)

	if err != nil {
		return err
	}

	err = json.NewDecoder(resp.body).Decode(&tempDC)
	if err != nil {
		return err
	}

	d.ID = tempDC.ID
	ensureReaderClosed(resp)

	return nil

}
