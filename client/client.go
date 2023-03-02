package client

import (
	"bytes"
	"encoding/json"
	"github.com/China-Chris/gpt/model/gpt3_5"
	"net/http"
)

const (
	RoleUser      gpt3_5.RoleType = "user"
	RoleAssistant gpt3_5.RoleType = "assistant"
	RoleSystem    gpt3_5.RoleType = "system"
)

const DefaultUrl = "https://api.openai.com/v1/chat/completions"

type Client struct {
	transport *http.Client
	apiKey    string
	url       string
}

func NewClient35(apiKey string) *Client {
	return &Client{
		transport: http.DefaultClient,
		apiKey:    apiKey,
		url:       DefaultUrl,
	}
}

func (c *Client) GetChat(r *gpt3_5.Request) (*gpt3_5.Response, error) {
	jsonData, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	client := &http.Client{}
	httpResp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer func() {
		_ = httpResp.Body.Close()
	}()
	var resp gpt3_5.Response
	err = json.NewDecoder(httpResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
