package main

import (
	"fmt"
	"time"
)

// ============= TRADITIONAL BUILDER (Verbose) =============

type HTTPRequest struct {
	url             string
	method          string
	headers         map[string]string
	body            string
	timeout         time.Duration
	followRedirects bool
	maxRetries      int
}

type HTTPRequestBuilder struct {
	request *HTTPRequest
}

func NewHTTPRequestBuilder() *HTTPRequestBuilder {
	return &HTTPRequestBuilder{
		request: &HTTPRequest{
			method:          "GET",
			headers:         make(map[string]string),
			timeout:         30 * time.Second,
			followRedirects: true,
			maxRetries:      3,
		},
	}
}

func (b *HTTPRequestBuilder) URL(url string) *HTTPRequestBuilder {
	b.request.url = url
	return b
}

func (b *HTTPRequestBuilder) Method(method string) *HTTPRequestBuilder {
	b.request.method = method
	return b
}

func (b *HTTPRequestBuilder) Header(key, value string) *HTTPRequestBuilder {
	b.request.headers[key] = value
	return b
}

func (b *HTTPRequestBuilder) Body(body string) *HTTPRequestBuilder {
	b.request.body = body
	return b
}

func (b *HTTPRequestBuilder) Timeout(timeout time.Duration) *HTTPRequestBuilder {
	b.request.timeout = timeout
	return b
}

func (b *HTTPRequestBuilder) Build() (*HTTPRequest, error) {
	if b.request.url == "" {
		return nil, fmt.Errorf("URL is required")
	}
	return b.request, nil
}

// ============= FUNCTIONAL OPTIONS PATTERN (Idiomatic Go) =============
// This is the PREFERRED Go way - used by grpc-go, uber-zap, etc.

type RequestOption func(*HTTPRequest)

func NewHTTPRequest(url string, opts ...RequestOption) *HTTPRequest {
	req := &HTTPRequest{
		url:             url,
		method:          "GET",
		headers:         make(map[string]string),
		timeout:         30 * time.Second,
		followRedirects: true,
		maxRetries:      3,
	}

	// Apply all options
	for _, opt := range opts {
		opt(req)
	}

	return req
}

// Option functions
func WithMethod(method string) RequestOption {
	return func(r *HTTPRequest) {
		r.method = method
	}
}

func WithHeader(key, value string) RequestOption {
	return func(r *HTTPRequest) {
		r.headers[key] = value
	}
}

func WithBody(body string) RequestOption {
	return func(r *HTTPRequest) {
		r.body = body
	}
}

func WithTimeout(timeout time.Duration) RequestOption {
	return func(r *HTTPRequest) {
		r.timeout = timeout
	}
}

func WithoutRedirects() RequestOption {
	return func(r *HTTPRequest) {
		r.followRedirects = false
	}
}

func WithRetries(retries int) RequestOption {
	return func(r *HTTPRequest) {
		r.maxRetries = retries
	}
}

func (r *HTTPRequest) Execute() {
	fmt.Printf("Executing %s %s\n", r.method, r.url)
	fmt.Printf("Headers: %v\n", r.headers)
	fmt.Printf("Timeout: %v\n", r.timeout)
}

func main() {
	// Traditional builder (verbose)
	req1, _ := NewHTTPRequestBuilder().
		URL("https://api.example.com/users").
		Method("POST").
		Header("Content-Type", "application/json").
		Header("Authorization", "Bearer token123").
		Body(`{"name":"John"}`).
		Timeout(5 * time.Second).
		Build()
	req1.Execute()

	fmt.Println("---")

	// Functional options (idiomatic Go) - Clean and composable
	req2 := NewHTTPRequest(
		"https://api.example.com/users",
		WithMethod("POST"),
		WithHeader("Content-Type", "application/json"),
		WithHeader("Authorization", "Bearer token123"),
		WithBody(`{"name":"John"}`),
		WithTimeout(5*time.Second),
		WithoutRedirects(),
		WithRetries(5),
	)
	req2.Execute()

	// Can create reusable option sets
	authOptions := []RequestOption{
		WithHeader("Authorization", "Bearer token123"),
		WithHeader("User-Agent", "MyApp/1.0"),
		WithTimeout(10 * time.Second),
	}

	req3 := NewHTTPRequest(
		"https://api.example.com/profile",
		authOptions..., // Spread operator
	)
	req3.Execute()
}
