package main

import (
	"github.com/China-Chris/gpt/client"
	"github.com/China-Chris/gpt/gpt35"
	"github.com/China-Chris/gpt/model/gpt3_5"
)

const apiKey = "your_api_key"

func main() {
	apiKey := "your_api_key"
	content := "hello"
	resp, err := gpt35.GetGPTResponse(apiKey, gpt3_5.TurboEngin, client.RoleUser, content, gpt35.OptionalParams{})
	if err != nil {
		panic(err)
	}
	println(resp)
}
