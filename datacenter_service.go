package vrlcmsdk

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/goharbor/harbor/src/common/utils/log"
	"github.com/pkg/errors"
	"github.com/sdbrett/vrlcmsdk/types"
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

	for x, i := range d.Datacenter{
		id := i.ID
		d.Datacenter[x], err = dc.GetDatacenter(ctx,id)
		if err != nil{
			log.Errorf("received error getting datacenter for datacenter id %s", id)
			return d, err
		}
		d.Datacenter[x].ID = id
		fmt.Printf(id)
	}
	return d, nil

}

func (dc *DatacenterAPIService) GetDatacenter(ctx context.Context, id string) (types.Datacenter, error) {

	url := dc.client.basePath + "/view/datacenter?datacenterId=" + id
	d := types.Datacenter{}

	resp, err := dc.client.get(ctx, url, *dc.client.headers)
	if err != nil {
		return d, err
	}

	err = json.NewDecoder(resp.body).Decode(&d)
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
