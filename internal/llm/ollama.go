package llm

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Ollama struct {
	Model    string
	Endpoint string // e.g., http://localhost:11434
	Client   *http.Client
}

type ollamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
}

type ollamaResponse struct {
	Response string `json:"response"`
}

func (o *Ollama) httpClient() *http.Client {
	if o.Client != nil {
		return o.Client
	}
	return &http.Client{Timeout: 30 * time.Second}
}

func (o *Ollama) Generate(ctx context.Context, prompt string) (string, error) {
	reqBody, _ := json.Marshal(ollamaRequest{Model: o.Model, Prompt: prompt})
	url := fmt.Sprintf("%s/api/generate", o.Endpoint)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(reqBody))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := o.httpClient().Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	var out ollamaResponse
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return "", err
	}
	return out.Response, nil
}
