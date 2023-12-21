package apiclients

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"
)

type clientBase struct {
	baseURL    string
	httpClient *http.Client
}

func newClientBase(baseURL string) *clientBase {
	return &clientBase{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 60 * time.Second,
		},
	}
}

func (c *clientBase) request(method, endpoint string, query map[string]string, header map[string]string, data []byte) (*http.Response, error) {
	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	queries := req.URL.Query()
	for key, value := range query {
		queries.Add(key, value)
	}

	req.URL.RawQuery = queries.Encode()

	for key, value := range header {
		req.Header.Add(key, value)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *clientBase) createEndpoint(path string) (string, error) {
	base, err := url.Parse(c.baseURL)
	if err != nil {
		return "", err
	}

	pathURL, err := url.Parse(path)
	if err != nil {
		return "", err
	}

	endpoint := base.ResolveReference(pathURL).String()

	return endpoint, nil
}

func (c *clientBase) unmarshalResponse(resp *http.Response, out any) error {
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, out)
}
