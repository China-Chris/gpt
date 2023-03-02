package gpt35

import (
	"github.com/China-Chris/gpt/client"
	"github.com/China-Chris/gpt/model/gpt3_5"
	"reflect"
)

type OptionalParams struct {
	Temperature      *float64
	TopP             *float64
	N                *int
	Stream           *bool
	Stop             interface{}
	MaxTokens        *int
	PresencePenalty  *float64
	FrequencyPenalty *float64
	LogitBias        interface{}
	User             *string
}

func setOptionalParam(paramName string, paramValue reflect.Value, req *gpt3_5.Request) {
	if paramValue.IsValid() {
		valueOfReq := reflect.ValueOf(req).Elem()
		valueOfReq.FieldByName(paramName).Set(paramValue.Elem())
	}
}

func GetGPTResponse(apiKey string, model gpt3_5.ModelType, role gpt3_5.RoleType, content string, optionalParams OptionalParams) (string, error) {
	c := client.NewClient35(apiKey)
	req := &gpt3_5.Request{
		Model: model,
		Messages: []*gpt3_5.Message{
			{
				Role:    role,
				Content: content,
			},
		},
	}
	valueOfParams := reflect.ValueOf(optionalParams)
	typeOfParams := valueOfParams.Type()
	for i := 0; i < valueOfParams.NumField(); i++ {
		paramName := typeOfParams.Field(i).Name
		paramValue := valueOfParams.Field(i)
		setOptionalParam(paramName, paramValue, req)
	}

	resp, err := c.GetChat(req)
	if err != nil {
		return "", err
	}
	return resp.Choices[0].Message.Content, nil
}
