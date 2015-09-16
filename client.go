package gomonkeylearn

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	baseAPI = "https://api.monkeylearn.com/v2/"
)

var (
	categorizers = map[string]string{
		"News Categorizer": "cl_hS9wMk9y",
	}
)

// ClassifyRequest is the JSON payload sent to MonkeyLearn for classification.
type ClassifyRequest struct {
	TextList []string `json:"text_list"`
}

// Client is responsible for all commications with the MonkeyLearn API (https://app.monkeylearn.com/.).
type Client struct {
	BaseURL      string
	Categorizers map[string]string
	transport    *http.Transport
	client       *http.Client
	apiToken     string
}

// NewClient generates a new HTTPS client that integrates with MonkeyLearn.
func NewClient(token string) *Client {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	return &Client{BaseURL: baseAPI, Categorizers: categorizers, transport: transport, client: &http.Client{Transport: transport}, apiToken: token}
}

// Classify takes one or more texts and passes them on to MonkeyLearn for classification.
func (c *Client) Classify(category string, texts []string) ([]byte, error) {
	// Ensure we support this category of classifier.
	class, ok := c.Categorizers[category]
	if !ok {
		return nil, fmt.Errorf("go-monkeylearn: unknown classifier category %s", category)
	}

	// "https://api.monkeylearn.com/v2/classifiers/cl_hS9wMk9y/classify/?"
	url := fmt.Sprintf("%s/classifiers/%s/classify/", c.BaseURL, class)
	token := fmt.Sprintf("token %s", c.apiToken)
	data, err := json.Marshal(&ClassifyRequest{texts})
	if err != nil {
		return nil, fmt.Errorf("go-monkeylearn: unable to marshal %#v to JSON: %s", texts, err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("go-monkeylearn: unable to create classification request: %s: %s", string(data), err)
	}
	req.Header.Add("Content-Type", `application/json`)
	req.Header.Add("Authorization", token)

	res, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("go-monkeylearn: unable to send classification request: %s", err)
	}
	defer res.Body.Close()
	mlData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("go-monkeylearn: unable to read response body: %s", err)
	}

	return mlData, nil
}
