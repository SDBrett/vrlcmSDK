package main

import (
	"fmt"
	"vrlcmSDK/vrlcmSdk"
)

func main() {

	var l vrlcmSdk.Client
	l = &vrlcmSdk.SdkConnection{BaseUrl: "https://192.168.17.145/lcm/api/v1", IgnoreCertError: true}
	err := l.Login("admin@localhost", "vmware")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(l)

}
