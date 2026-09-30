package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type Client struct {
	httpClient *http.Client
	baseURL    string
}

type Response struct {
	StatusCode int
	Headers    http.Header
	Body       []byte
}

type APIError struct {
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

func NewClient(baseURL string) *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		baseURL: strings.TrimRight(baseURL, "/"),
	}
}

func decodeResponse[T any](resp *http.Response) (T, *APIError, error) {
	var result T

	defer resp.Body.Close()

	switch {
	case resp.StatusCode >= 200 && resp.StatusCode < 300:
		err := json.NewDecoder(resp.Body).Decode(&result)
		if err != nil {
			return result, nil, err
		}

		return result, nil, nil

	default:
		var apiErr APIError

		err := json.NewDecoder(resp.Body).Decode(&apiErr)
		if err != nil {
			return result, nil, err
		}

		return result, &apiErr, nil
	}
}

func (c *Client) Do(method, path string, body any) (*http.Response, error) {
	var requestBody []byte

	if body != nil {
		var err error

		requestBody, err = json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal request body: %w", err)
		}
	}

	req, err := http.NewRequest(
		method,
		c.baseURL+path,
		bytes.NewReader(requestBody),
	)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("send request: %w", err)
	}


	return resp, nil
}