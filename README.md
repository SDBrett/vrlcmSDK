# vRealize Lifecycle Manager

[![Build Status](https://travis-ci.org/SDBrett/vrlcmSDK.svg?branch=master)](https://travis-ci.org/SDBrett/vrlcmSDK)
[![Coverage Status](https://coveralls.io/repos/github/SDBrett/vrlcmSDK/badge.svg?branch=master)](https://coveralls.io/github/SDBrett/vrlcmSDK?branch=master)

This project provides an GO SDK for vRealize Lifecycle Manager. 

This is a personal learning project and may or may not eventuate into something usable.

## Versions tested against
- GO 1.11.1
- vRLCM 2.1.0

#### Installation

To obtain the package run the following GO command.

`go get github.com/sdbrett/vrlcmsdk`

## Examples

#### Connect to a vRLCM server

You need to authenticate against the vRLCM server before you're able to perform any tasks.

```go
var ctx = context.Background() // Setup context
cli := NewApiClient("127.0.0.1", true, cli) // Create new  API client
cli.login(Username: "Username", Password: "Password") // Authenticate against vRLCM instance
```

#### Get Datacenters

You can interact with services exposed through the vRLCM API after successful authentication.

```go
dataCenters, err := cli.DatacenterService.GetAllDatacenters(ctx)
if err != nil {
    t.Error(err)
}
```

## API Spec

The vRLCM API spec is available [here](https://code.vmware.com/apis/228/vrealize-suite-lifecycle-manager#/)