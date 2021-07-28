package client

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Get request client
func GetRequest(baseUri string, parameters string) (*http.Response, error) {
	if parameters != "" {
		parameters = "/" + parameters
	}
	response, err := http.Get(baseUri + parameters)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	return response, nil
}

// Post request client
func PostRequest(baseUri string, parameters string, body io.Reader) (*http.Response, error) {
	if parameters != "" {
		parameters = "/" + parameters
	}
	response, err := http.Post(baseUri+parameters, "application/json", body)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	return response, nil
}

// Put request client
func PutRequest(baseUri string, parameters string, body io.Reader) (*http.Response, error) {
	if parameters != "" {
		parameters = "/" + parameters
	}
	req, err := http.NewRequest(http.MethodPut, baseUri+parameters, body)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// Delete request client
func DeleteRequest(baseUri string, parameters string) (*http.Response, error) {
	if parameters != "" {
		parameters = "/" + parameters
	}
	req, err := http.NewRequest(http.MethodDelete, baseUri+parameters, nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return response, nil
}
