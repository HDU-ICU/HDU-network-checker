package utils

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"

	"github.com/ljcbaby/HDU-network-checker/log"
)

// Get performs an HTTP GET request to the specified URL and returns the response body as a string.
func Get(url string) (string, error) {
	log.Logger.Sugar().Debugf("GET %s", url)
	// Create a custom HTTP client with a transport that ignores certificate errors
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	// Send the GET request
	resp, err := client.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to perform GET request: %v", err)
	}
	defer resp.Body.Close()

	log.Logger.Sugar().Debugf("Response: %+v", resp)

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	log.Logger.Sugar().Debugf("Body: %s", body)

	return string(body), nil
}
