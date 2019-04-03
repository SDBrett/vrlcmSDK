package main

import (
	"fmt"
	"github.com/sdbrett/vrlcmsdk/client"
	"log"
)

func main() {

	var c client.SdkConnection
	c = client.SdkConnection{BaseUrl: "https://192.168.17.145/lcm/api/v1", IgnoreCertError: true}

	err := c.Login("admin@localhost", "vmware")
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(c)
	var d client.Datacenters
	err = d.GetDatacenters(c)
	if err != nil {
		log.Fatalf("received error: %s", err)
	}
	fmt.Println(d.Datacenter[0].ID)

}
