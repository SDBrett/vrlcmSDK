package vrlcmsdk

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestSettingsAPIService_RefreshPatches(t *testing.T) {

	var ctx = context.Background()
	expectedURL := "127.0.0.1/lcm/gui/api/action/refreshPatches"

	t.Run("Test successful patch refresh", func(t *testing.T) {

		serverResponse := `{"requestId":"2c551b23da979e7558a0c66411072"}`
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

		request, err := c.SettingsService.RefreshPatches(ctx)
		if err != nil {
			t.Errorf("expected no error when refreshing patches")
		}
		if request.RequestID != "2c551b23da979e7558a0c66411072" {
			t.Errorf("expected requestID to equeal 2c551b23da979e7558a0c66411072, received %s", request.RequestID)
		}
	})

	t.Run("Test server error", func(t *testing.T) {

		serverResponse := `{}`
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

		_, err := c.SettingsService.RefreshPatches(ctx)
		if err == nil {
			t.Errorf("expected no error status code error, receive nil error")
		}
	})

	t.Run("Test server JSON decode error", func(t *testing.T) {

		serverResponse := "TEST"
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

		_, err := c.SettingsService.RefreshPatches(ctx)
		if err == nil {
			t.Errorf("expected error JSON decode error")
		}

	})

}
