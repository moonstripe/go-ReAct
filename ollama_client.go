package main

import (
	"fmt"
	"net/url"
)

type OllamaClient struct {
	BaseUrl url.URL
}

type OllamaQueryResponse string

func NewOllamaClient(baseUrl string) (*OllamaClient, error) {
	bU, err := url.Parse(baseUrl)

	if err != nil {
		return nil, fmt.Errorf("could not parse baseUrl: %v", err)
	}

	return &OllamaClient{
		BaseUrl: *bU,
	}, nil
}

func (oC *OllamaClient) Query() (OllamaQueryResponse, error) {
	chatUrlPath := oC.BaseUrl.Path + "/api/chat"

	chatUrl := url.URL{
		Path: chatUrlPath,
	}
}
