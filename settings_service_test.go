package vrlcmsdk

import (
	"bytes"
	"context"
	"fmt"
	"github.com/sdbrett/vrlcmsdk/types"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestSettingsAPIService_GetLcmConfigSettings(t *testing.T) {

	var ctx = context.Background()
	expectedURL := "127.0.0.1/lcm/api/settings/lcmconfig"

	t.Run("Test successful Get LCM settings", func(t *testing.T) {

		serverResponse := `{"sshEnabled":"True","telemetryEnabled":"False","hostName":"photon-machine"}`
		cli := newMockClient(func(req *http.Request) (*http.Response, error) {
			if !strings.HasPrefix(req.URL.Path, expectedURL) {
				return nil, fmt.Errorf("expected URL '%s', got '%s'", expectedURL, req.URL)
			}
			if req.Method != "GET" {
				return nil, fmt.Errorf("expected GET method, got %s", req.Method)
			}

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(serverResponse))),
			}, nil
		})

		c := NewApiClient("127.0.0.1", true, cli)

		currentSettings, err := c.SettingsService.GetLcmConfigSettings(ctx)
		if err != nil {
			t.Errorf("expected no error when getting LCM config settings")
		}
		if currentSettings.SshEnabled != "True" {
			t.Errorf("expected SSH enabled to be True, instead got %s", currentSettings.SshEnabled)
		}
		if currentSettings.TelemetryEnabled != "False" {
			t.Errorf("expected Telemetry enabled to be False, instead got %s", currentSettings.TelemetryEnabled)
		}
	})

	t.Run("Test server error response", func(t *testing.T) {

		serverResponse := `{"message":"forbidden","statusCode":403,"documentKind":"com:vmware:xenon:common:ServiceErrorResponse","errorCode":0}`
		cli := newMockClient(func(req *http.Request) (*http.Response, error) {
			if !strings.HasPrefix(req.URL.Path, expectedURL) {
				return nil, fmt.Errorf("expected URL '%s', got '%s'", expectedURL, req.URL)
			}
			if req.Method != "GET" {
				return nil, fmt.Errorf("expected GET method, got %s", req.Method)
			}

			return &http.Response{
				StatusCode: http.StatusForbidden,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(serverResponse))),
			}, nil
		})

		c := NewApiClient("127.0.0.1", true, cli)

		_, err := c.SettingsService.GetLcmConfigSettings(ctx)
		if err == nil {
			t.Errorf("expected error to be thrown when with server error response")
		}
	})

	t.Run("Test bad JSON decoding", func(t *testing.T) {

		serverResponse := "TEST"
		cli := newMockClient(func(req *http.Request) (*http.Response, error) {
			if !strings.HasPrefix(req.URL.Path, expectedURL) {
				return nil, fmt.Errorf("expected URL '%s', got '%s'", expectedURL, req.URL)
			}
			if req.Method != "GET" {
				return nil, fmt.Errorf("expected GET method, got %s", req.Method)
			}

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(serverResponse))),
			}, nil
		})

		c := NewApiClient("127.0.0.1", true, cli)

		_, err := c.SettingsService.GetLcmConfigSettings(ctx)
		if err == nil {
			t.Errorf("expected error to be thrown when decoding fails")
		}
	})

}

func TestSettingsAPIService_GetNetworkStatus(t *testing.T) {

	var ctx = context.Background()
	expectedURL := "127.0.0.1/lcm/gui/api/lcmVaNetworkStatus"

	t.Run("Test successful Get network settings", func(t *testing.T) {

		serverResponse := `{"hostName" : "photon-machine","diskSize" : "99G","diskUsedPercentage" : "1%","netmask" : "255.255.255.0","diskUsed" : "704M","diskAvail" : "93G","ipv4Address" : "192.168.17.128","preferredDns" : "192.168.17.2","type" : "DHCPV4+NONEV6","gateway" : "192.168.17.2"}`
		cli := newMockClient(func(req *http.Request) (*http.Response, error) {
			if !strings.HasPrefix(req.URL.Path, expectedURL) {
				return nil, fmt.Errorf("expected URL '%s', got '%s'", expectedURL, req.URL)
			}
			if req.Method != "GET" {
				return nil, fmt.Errorf("expected GET method, got %s", req.Method)
			}

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(serverResponse))),
			}, nil
		})

		c := NewApiClient("127.0.0.1", true, cli)

		networkStatus, err := c.SettingsService.GetNetworkStatus(ctx)
		if err != nil {
			t.Errorf("expected no error when getting LCM config settings")
		}
		if networkStatus.Hostname != "photon-machine" {
			t.Errorf("expected hostname of photon-machine, instead got %s", networkStatus.Hostname)
		}
		if networkStatus.Netmask != "255.255.255.0" {
			t.Errorf("expected netask of 255.255.255.0, instead got %s", networkStatus.Netmask)
		}
	})

	t.Run("Test server error response", func(t *testing.T) {

		serverResponse := `{"message":"forbidden","statusCode":403,"documentKind":"com:vmware:xenon:common:ServiceErrorResponse","errorCode":0}`
		cli := newMockClient(func(req *http.Request) (*http.Response, error) {
			if !strings.HasPrefix(req.URL.Path, expectedURL) {
				return nil, fmt.Errorf("expected URL '%s', got '%s'", expectedURL, req.URL)
			}
			if req.Method != "GET" {
				return nil, fmt.Errorf("expected GET method, got %s", req.Method)
			}

			return &http.Response{
				StatusCode: http.StatusForbidden,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(serverResponse))),
			}, nil
		})

		c := NewApiClient("127.0.0.1", true, cli)

		_, err := c.SettingsService.GetNetworkStatus(ctx)
		if err == nil {
			t.Errorf("expected error to be thrown when with server error response")
		}
	})

	t.Run("Test bad JSON decoding ", func(t *testing.T) {

		serverResponse := "TEST"
		cli := newMockClient(func(req *http.Request) (*http.Response, error) {
			if !strings.HasPrefix(req.URL.Path, expectedURL) {
				return nil, fmt.Errorf("expected URL '%s', got '%s'", expectedURL, req.URL)
			}
			if req.Method != "GET" {
				return nil, fmt.Errorf("expected GET method, got %s", req.Method)
			}

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(serverResponse))),
			}, nil
		})

		c := NewApiClient("127.0.0.1", true, cli)

		_, err := c.SettingsService.GetNetworkStatus(ctx)
		if err == nil {
			t.Errorf("expected error to be thrown when decoding fails")
		}
	})

}

func TestSettingsAPIService_SetRootPassword(t *testing.T) {

	var ctx = context.Background()
	expectedURL := "127.0.0.1/lcm/api/v1/settings"
	newPassword := "vmware1!"

	t.Run("Test successful Set SSH password", func(t *testing.T) {

		serverResponse := `{"id":null,"type":null,"state":null,"status":"SUCCESS","isRetriable":null,"retryParameters":null}`
		cli := newMockClient(func(req *http.Request) (*http.Response, error) {
			if !strings.HasPrefix(req.URL.Path, expectedURL) {
				return nil, fmt.Errorf("expected URL '%s', got '%s'", expectedURL, req.URL)
			}
			if req.Method != "POST" {
				return nil, fmt.Errorf("expected POST method, got %s", req.Method)
			}

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(serverResponse))),
			}, nil
		})

		c := NewApiClient("127.0.0.1", true, cli)

		err := c.SettingsService.SetRootPassword(ctx, newPassword)
		if err != nil {
			t.Errorf("expected no error when getting LCM config settings")
		}
	})

	t.Run("Test server error response", func(t *testing.T) {

		serverResponse := `{"message":"forbidden","statusCode":403,"documentKind":"com:vmware:xenon:common:ServiceErrorResponse","errorCode":0}`
		cli := newMockClient(func(req *http.Request) (*http.Response, error) {
			if !strings.HasPrefix(req.URL.Path, expectedURL) {
				return nil, fmt.Errorf("expected URL '%s', got '%s'", expectedURL, req.URL)
			}
			if req.Method != "POST" {
				return nil, fmt.Errorf("expected POST method, got %s", req.Method)
			}

			return &http.Response{
				StatusCode: http.StatusForbidden,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(serverResponse))),
			}, nil
		})

		c := NewApiClient("127.0.0.1", true, cli)

		err := c.SettingsService.SetRootPassword(ctx, newPassword)
		if err == nil {
			t.Errorf("expected error to be thrown when with server error response")
		}
	})

}

func TestSettingsAPIService_SetAdminPassword(t *testing.T) {

	var ctx = context.Background()
	expectedURL := "127.0.0.1/lcm/api/v1/settings"
	newPassword := "vmware1!"

	t.Run("Test successful set admin password", func(t *testing.T) {

		serverResponse := `{"id":null,"type":null,"state":null,"status":"SUCCESS","isRetriable":null,"retryParameters":null}`
		cli := newMockClient(func(req *http.Request) (*http.Response, error) {
			if !strings.HasPrefix(req.URL.Path, expectedURL) {
				return nil, fmt.Errorf("expected URL '%s', got '%s'", expectedURL, req.URL)
			}
			if req.Method != "POST" {
				return nil, fmt.Errorf("expected POST method, got %s", req.Method)
			}

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(serverResponse))),
			}, nil
		})

		c := NewApiClient("127.0.0.1", true, cli)

		err := c.SettingsService.SetAdminPassword(ctx, newPassword)
		if err != nil {
			t.Errorf("expected no error when getting LCM config settings")
		}
	})

	t.Run("Test server error response", func(t *testing.T) {

		serverResponse := `{"message":"forbidden","statusCode":403,"documentKind":"com:vmware:xenon:common:ServiceErrorResponse","errorCode":0}`
		cli := newMockClient(func(req *http.Request) (*http.Response, error) {
			if !strings.HasPrefix(req.URL.Path, expectedURL) {
				return nil, fmt.Errorf("expected URL '%s', got '%s'", expectedURL, req.URL)
			}
			if req.Method != "POST" {
				return nil, fmt.Errorf("expected POST method, got %s", req.Method)
			}

			return &http.Response{
				StatusCode: http.StatusForbidden,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(serverResponse))),
			}, nil
		})

		c := NewApiClient("127.0.0.1", true, cli)

		err := c.SettingsService.SetAdminPassword(ctx, newPassword)
		if err == nil {
			t.Errorf("expected error to be thrown when with server error response")
		}
	})

}

func TestSettingsAPIService_SetRestartSchedule(t *testing.T) {

	var ctx = context.Background()
	expectedURL := "127.0.0.1/lcm/api/maintenance/xserver-restart-config"

	t.Run("Test successful setting restart schedule", func(t *testing.T) {

		serverResponse := `{"day":"1","hour":"16","weeklyServerRestartEnable":true}`
		schedule := types.RestartSchedule{WeeklyServerRestartEnable: true, Day: "1", Hour: "14"}
		cli := newMockClient(func(req *http.Request) (*http.Response, error) {
			if !strings.HasPrefix(req.URL.Path, expectedURL) {
				return nil, fmt.Errorf("expected URL '%s', got '%s'", expectedURL, req.URL)
			}
			if req.Method != "POST" {
				return nil, fmt.Errorf("expected POST method, got %s", req.Method)
			}

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(serverResponse))),
			}, nil
		})

		c := NewApiClient("127.0.0.1", true, cli)

		err := c.SettingsService.SetRestartSchedule(ctx, schedule)
		if err != nil {
			t.Errorf("expected no error when setting restart schedule")
		}
	})

	t.Run("Test successful disable restart schedule", func(t *testing.T) {

		serverResponse := `{"day":"","hour":"","weeklyServerRestartEnable":false}`
		schedule := types.RestartSchedule{WeeklyServerRestartEnable: false, Day: "", Hour: ""}
		cli := newMockClient(func(req *http.Request) (*http.Response, error) {
			if !strings.HasPrefix(req.URL.Path, expectedURL) {
				return nil, fmt.Errorf("expected URL '%s', got '%s'", expectedURL, req.URL)
			}
			if req.Method != "POST" {
				return nil, fmt.Errorf("expected POST method, got %s", req.Method)
			}

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(serverResponse))),
			}, nil
		})

		c := NewApiClient("127.0.0.1", true, cli)

		err := c.SettingsService.SetRestartSchedule(ctx, schedule)
		if err != nil {
			t.Errorf("expected no error when disabling restart schedule")
		}
	})

	t.Run("Test server error response", func(t *testing.T) {

		serverResponse := `{"message":"com.google.gson.stream.MalformedJsonException: Unterminated object at line 3 column 6 path $.null","statusCode":500,"documentKind":"com:vmware:xenon:common:ServiceErrorResponse","errorCode":0}`
		schedule := types.RestartSchedule{WeeklyServerRestartEnable: false, Day: "", Hour: ""}
		cli := newMockClient(func(req *http.Request) (*http.Response, error) {
			if !strings.HasPrefix(req.URL.Path, expectedURL) {
				return nil, fmt.Errorf("expected URL '%s', got '%s'", expectedURL, req.URL)
			}
			if req.Method != "POST" {
				return nil, fmt.Errorf("expected POST method, got %s", req.Method)
			}

			return &http.Response{
				StatusCode: http.StatusInternalServerError,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(serverResponse))),
			}, nil
		})

		c := NewApiClient("127.0.0.1", true, cli)

		err := c.SettingsService.SetRestartSchedule(ctx, schedule)
		if err == nil {
			t.Errorf("expected error to be thrown when with server error response")
		}
	})

}
