package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	id     string
	apikey string
}

func (c *Client) baseUrl() string {
	if c.id != "" {
		return "https://head." + c.id + ".cyberfile.fr/api/v1/"
	}
	return "https://head.cyberfile.fr/api/v1/"
}

func deleteJsonData(service string, c *Client) (*ResponseError, error) {
	method := "DELETE"
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	url := c.baseUrl() + service
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, fmt.Errorf("Got error %s", err.Error())
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("user-agent", "golang application")
	req.Header.Add("Authorization", "Bearer "+c.apikey)

	response, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("Got error %s", err.Error())
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("Got error %s", err.Error())
	}

	var result ResponseError
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, nil
	}

	return &result, errors.New(result.Message)
}

func getJsonData[T Data](service string, c *Client) (*T, error) {
	method := "GET"
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	url := c.baseUrl() + service
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, fmt.Errorf("Got error %s", err.Error())
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("user-agent", "golang application")
	req.Header.Add("Authorization", "Bearer "+c.apikey)

	response, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("Got error %s", err.Error())
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("Got error %s", err.Error())
	}

	var result T
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, fmt.Errorf("Got error %s", err.Error())
	}

	return &result, nil
}

func postJsonData[T Data, R RequestData](service string, c *Client, requestData *R) (*T, error) {
	method := "POST"
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	url := c.baseUrl() + service

	reqBody, err := json.Marshal(requestData)
	if err != nil {
		return nil, fmt.Errorf("Got error %s", err.Error())
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("Got error %s", err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("user-agent", "golang application")
	req.Header.Add("Authorization", "Bearer "+c.apikey)

	response, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("Got error %s", err.Error())
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("Got error %s", err.Error())
	}

	var result T
	err = json.Unmarshal(body, &result)
	if err != nil {
		// try to unmarshall the ResponseError type
		// TODO
		//
		return nil, fmt.Errorf("Got error %s", err.Error())
	}

	return &result, nil
}
