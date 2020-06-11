package models

import (
	"errors"

	v1 "github.com/nivista/tasktimer/api/v1"
)

// Task represents a recurring task
type Task interface {
	Visit(TaskVisitor)
}

// TaskVisitor should be implemented by anyone who wants to write a method that can handle any task
type TaskVisitor interface {
	VisitHTTP(HTTP) // Do you want your implementation to return something? Use a closure instead.
}

// Method is the HTTP request method
type Method int

const (
	// GET is code for a get request
	GET Method = Method(v1.Method_GET)

	// POST is code for a post request
	POST = Method(v1.Method_POST)
)

// HTTP represents the configuration for executing an HTTP request
type HTTP struct {
	URL string
	Method
	Body    string
	Headers []string
}

// Visit calls VisitHTTP on the TaskVisitor
func (h HTTP) Visit(t TaskVisitor) {
	t.VisitHTTP(h)
}

type protoTaskGenerator struct{ *v1.Timer }

func (p protoTaskGenerator) VisitHTTP(h HTTP) {
	p.TaskConfig = &v1.Timer_HttpConfig{HttpConfig: &v1.HTTPConfig{
		Url:     h.URL,
		Method:  v1.Method(h.Method),
		Body:    h.Body,
		Headers: h.Headers,
	}}
}

func toTaskConfig(p *v1.Timer) (Task, error) {
	switch config := p.TaskConfig.(type) {
	case *v1.Timer_HttpConfig:
		pHTTPConfig := config.HttpConfig
		http := HTTP{
			URL:     pHTTPConfig.Url,
			Method:  Method(pHTTPConfig.Method),
			Body:    pHTTPConfig.Body,
			Headers: pHTTPConfig.Headers,
		}
		return &http, nil
	default:
		return nil, errors.New("unable to extract task config")
	}
}
