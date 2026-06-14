package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"go-matrix-api/internal/domain"
)

type httpClient struct {
	client *http.Client
	url    string
}

// NewNodeClient creates a new instance of the Node.js API client
func NewNodeClient() domain.NodeClient {
	url := os.Getenv("NODE_API_URL")
	if url == "" {
		url = "http://localhost:4000/api/analytics"
	}

	return &httpClient{
		client: &http.Client{Timeout: 10 * time.Second},
		url:    url,
	}
}

type analyticsRequest struct {
	Q [][]float64 `json:"Q"`
	R [][]float64 `json:"R"`
}

type analyticsResponse struct {
	Success bool                   `json:"success"`
	Data    *domain.NodeStatistics `json:"data"`
	Error   string                 `json:"error"`
}

// GetStatistics sends Q and R matrices to Node.js API and retrieves stats
func (c *httpClient) GetStatistics(q, r [][]float64, token string) (*domain.NodeStatistics, error) {
	reqBody := analyticsRequest{
		Q: q,
		R: r,
	}

	jsonBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, c.url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	if token != "" {
		req.Header.Set("Authorization", token)
	}

	log.Printf("[NodeClient] Sending POST request to %s", c.url)
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Node API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("node API returned status code %d", resp.StatusCode)
	}

	var resData analyticsResponse
	if err := json.NewDecoder(resp.Body).Decode(&resData); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	if !resData.Success {
		return nil, errors.New(resData.Error)
	}

	log.Println("[NodeClient] Analytics request succeeded")
	return resData.Data, nil
}
