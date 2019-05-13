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

func TestApiClient_Login(t *testing.T) {

	var ctx = context.Background()
	expectedURL := "127.0.0.1/lcm/api/v1/login"

	t.Run("Test successful vCenter login", func(t *testing.T) {

		serverResponse := `{"token":"eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJ4biIsInN1YiI6Ii9jb3JlL2F1dGh6L3VzZXJzL3ZMQ01BZG1pbiIsImV4cCI6NDcxMTc4ODMxN30.LBFSBgdwr9T2xh0qD0ElcfqpWozP_C6SNmXvWG9vz44"}`
		expectedToken := "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJ4biIsInN1YiI6Ii9jb3JlL2F1dGh6L3VzZXJzL3ZMQ01BZG1pbiIsImV4cCI6NDcxMTc4ODMxN30.LBFSBgdwr9T2xh0qD0ElcfqpWozP_C6SNmXvWG9vz44"

		cli := newMockClient(func(req *http.Request) (*http.Response, error) {
			if !strings.HasPrefix(req.URL.Path, expectedURL) {
				return nil, fmt.Errorf("expected URL '%s', got '%s'", expectedURL, req.URL)
			}
			if req.Method != "POST" {
				return nil, fmt.Errorf("expected GET method, got %s", req.Method)
			}

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(serverResponse))),
			}, nil
		})

		c := NewApiClient("127.0.0.1", true, cli)

		err := c.Login(ctx, "admin@localhost", "vmware")
		if err != nil {
			t.Errorf("expected no error during login")
		}
		if c.token != expectedToken {
			t.Errorf("expected token code did not match result")
		}
	})

	t.Run("Test bad JSON decoding", func(t *testing.T) {

		serverResponse := `{"message":"Unparseable JSON body: com.google.gson.stream.MalformedJsonException: Expected name at line 4 column 2 path $.","statusCode":400,"documentKind":"com:vmware:xenon:common:ServiceErrorResponse","errorCode":0}`
		cli := newMockClient(func(req *http.Request) (*http.Response, error) {
			if !strings.HasPrefix(req.URL.Path, expectedURL) {
				return nil, fmt.Errorf("expected URL '%s', got '%s'", expectedURL, req.URL)
			}
			if req.Method != "POST" {
				return nil, fmt.Errorf("expected GET method, got %s", req.Method)
			}

			return &http.Response{
				StatusCode: http.StatusBadRequest,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(serverResponse))),
			}, nil
		})

		c := NewApiClient("127.0.0.1", true, cli)

		err := c.Login(ctx, "admin@localhost", "vmware")
		if err == nil {
			t.Errorf("expected error to be thrown when with server error response")
		}
	})

	t.Run("Test server response error for vCenter creation", func(t *testing.T) {

		serverResponse := "TEST"
		cli := newMockClient(func(req *http.Request) (*http.Response, error) {
			if !strings.HasPrefix(req.URL.Path, expectedURL) {
				return nil, fmt.Errorf("expected URL '%s', got '%s'", expectedURL, req.URL)
			}
			if req.Method != "POST" {
				return nil, fmt.Errorf("expected GET method, got %s", req.Method)
			}

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(serverResponse))),
			}, nil
		})

		c := NewApiClient("127.0.0.1", true, cli)

		err := c.Login(ctx, "admin@localhost", "vmware")
		if err == nil {
			t.Errorf("expected error to be thrown when decoding fails")
		}
	})
}

func TestApiClient_Logout(t *testing.T) {

	var ctx = context.Background()
	expectedURL := "127.0.0.1/lcm/api/v1/logout"

	t.Run("Test successful vCenter logout", func(t *testing.T) {

		serverResponse := ""
		cli := newMockClient(func(req *http.Request) (*http.Response, error) {
			if !strings.HasPrefix(req.URL.Path, expectedURL) {
				return nil, fmt.Errorf("expected URL '%s', got '%s'", expectedURL, req.URL)
			}
			if req.Method != "POST" {
				return nil, fmt.Errorf("expected GET method, got %s", req.Method)
			}

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(serverResponse))),
			}, nil
		})

		c := NewApiClient("127.0.0.1", true, cli)

		err := c.Logout(ctx)
		if err != nil {
			t.Errorf("expected no error during login")
		}
	})

	t.Run("Test server error response", func(t *testing.T) {

		serverResponse := `{"message":"Unparseable JSON body: com.google.gson.stream.MalformedJsonException: Expected name at line 4 column 2 path $.","statusCode":400,"documentKind":"com:vmware:xenon:common:ServiceErrorResponse","errorCode":0}`
		cli := newMockClient(func(req *http.Request) (*http.Response, error) {
			if !strings.HasPrefix(req.URL.Path, expectedURL) {
				return nil, fmt.Errorf("expected URL '%s', got '%s'", expectedURL, req.URL)
			}
			if req.Method != "POST" {
				return nil, fmt.Errorf("expected GET method, got %s", req.Method)
			}

			return &http.Response{
				StatusCode: http.StatusBadRequest,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(serverResponse))),
			}, nil
		})

		c := NewApiClient("127.0.0.1", true, cli)

		err := c.Logout(ctx)
		if err == nil {
			t.Errorf("expected error to be thrown when with server error response")
		}
	})
}
