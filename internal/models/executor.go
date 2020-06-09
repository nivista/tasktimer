package models

import (
	"errors"
	"fmt"

	v1 "github.com/nivista/tasktimer/api/v1"
)

type executable interface {
	assignToProto(*v1.Timer) error
	GetExec() (func(), error)
	IsValid() (bool, string)
}

type method v1.Method

const (
	get  method = method(v1.Method_GET)
	post        = method(v1.Method_POST)
)

// HTTPConfig represents the configuration for the HTTP executor
type HTTPConfig struct {
	url string
	method
	body    string
	headers []string
}

// IsValid returns true, "" if the configuration is valid, otherwise it returns false and an explanation
func (c *HTTPConfig) IsValid() (bool, string) {
	return true, ""
}

// GetExec returns the executor function for a given configuration, panics if IsValid returns false
func (c *HTTPConfig) GetExec() (func(), error) {
	if isValid, reason := c.IsValid(); !isValid {
		panic(reason)
	}
	switch c.method {
	case get:
		f := func() {
			fmt.Printf("GET %v", c.url)
		}
		return f, nil
	case post:
		f := func() {
			fmt.Printf("POST %v", c.url)
		}
		return f, nil
	}
	panic("Method not found")
}

func (c *HTTPConfig) assignToProto(p *v1.Timer) error {
	p.ExecutorConfig = &v1.Timer_HttpConfig{HttpConfig: &v1.HTTPConfig{
		Url:     c.url,
		Method:  v1.Method(c.method),
		Body:    c.body,
		Headers: c.headers,
	}}
	return nil
}

func toExecutable(p *v1.Timer) (executable, error) {
	switch config := p.ExecutorConfig.(type) {
	case *v1.Timer_HttpConfig:
		pHTTPConfig := config.HttpConfig
		httpConfig := HTTPConfig{
			url:     pHTTPConfig.Url,
			body:    pHTTPConfig.Body,
			headers: pHTTPConfig.Headers,
		}
		return &httpConfig, nil
	default:
		return nil, errors.New("unable to extract executable")
	}
}
