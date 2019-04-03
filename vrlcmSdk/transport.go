package vrlcmSdk

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"
)

//Default Http transport configuration
func NewDefaultSdkTransport(skipCertVerify bool) http.RoundTripper {
	t := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 90 * time.Second,
			DualStack: true,
		}).DialContext,
		TLSClientConfig: &tls.Config{InsecureSkipVerify: skipCertVerify},
	}

	return t
}

func NewApiClient(options ...func(*SdkConnection)) (*SdkConnection, error) {
	c := http.Client{}

	client := SdkConnection{Client: c}

	for _, option := range options {
		option(&client)
	}

	return &client, nil
}

func (s *SdkConnection) newDefaultClient() {

	// Setup http transport using default transport
	t := func(c *SdkConnection) {
		c.Client.Transport = NewDefaultSdkTransport(s.IgnoreCertError)
	}

	// Create new http client
	c, _ := NewApiClient(t)

	s.Client = c.Client
}
