package main

import (
	"bytes"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"gofr.dev/pkg/gofr/request"
)

func Test_Integration(t *testing.T) {
	go main()

	time.Sleep(5 * time.Second)

	carCreateBody := []byte(`{"id": 1, "name": "test", "color": "test-color"}`)
	carUpdateBody := []byte(`{"id": 1, "name": "test-name", "color": "test-color"}`)

	successResp := `{"data":{"ID":1,"name":"test","color":"test-color"}}`
	successUpdateResp := `{"data":{"ID":1,"name":"test-name","color":"test-color"}}`

	testCases := []struct {
		desc          string
		method        string
		endpoint      string
		body          []byte
		expStatusCode int
		expResp       string
	}{
		{"Create car", http.MethodPost, "/car", carCreateBody, http.StatusCreated,
			successResp},
		{"Get car", http.MethodGet, "/car/1", nil, http.StatusOK, successResp},
		{"Update car", http.MethodPut, "/car/1", carUpdateBody, http.StatusOK,
			successUpdateResp},
		{"Delete car", http.MethodDelete, "/car/1", nil, http.StatusNoContent, ``},
	}

	for i, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			req, _ := request.NewMock(tc.method, "http://localhost:9000"+tc.endpoint, bytes.NewBuffer(tc.body))
			client := http.Client{}

			resp, err := client.Do(req)
			if err != nil {
				t.Fatalf("Error occurred in calling api: %v", err)
			}

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("Error while reading response: %v", err)
			}

			respBody := strings.TrimSpace(string(body))

			assert.Equal(t, tc.expStatusCode, resp.StatusCode, "Test [%d] failed", i+1)
			assert.Equal(t, tc.expResp, respBody, "Test [%d] failed", i+1)

			resp.Body.Close()
		})
	}
}
