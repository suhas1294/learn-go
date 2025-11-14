package main

import (
	"context"
	"fmt"
	"time"
)

// ============= REQUEST/RESPONSE =============

type Request struct {
	UserID    string
	Token     string
	Body      map[string]interface{}
	Timestamp time.Time
	Context   context.Context
}

type Response struct {
	StatusCode int
	Body       map[string]interface{}
	Error      error
}

// ============= TRADITIONAL CHAIN OF RESPONSIBILITY =============

type Handler interface {
	SetNext(handler Handler)
	Handle(request *Request) *Response
}

type BaseHandler struct {
	next Handler
}

func (b *BaseHandler) SetNext(handler Handler) {
	b.next = handler
}

func (b *BaseHandler) HandleNext(request *Request) *Response {
	if b.next != nil {
		return b.next.Handle(request)
	}
	return &Response{StatusCode: 200, Body: map[string]interface{}{"status": "ok"}}
}

// ============= CONCRETE HANDLERS =============

type AuthenticationHandler struct {
	BaseHandler
}

func (a *AuthenticationHandler) Handle(request *Request) *Response {
	fmt.Println("🔐 Authentication Handler")

	if request.Token == "" {
		return &Response{
			StatusCode: 401,
			Body:       map[string]interface{}{"error": "Unauthorized"},
		}
	}

	// Validate token
	if request.Token != "valid_token" {
		return &Response{
			StatusCode: 401,
			Body:       map[string]interface{}{"error": "Invalid token"},
		}
	}

	fmt.Println("✓ Authentication passed")
	return a.HandleNext(request)
}

type RateLimitHandler struct {
	BaseHandler
	requestCounts map[string]int
	limit         int
}

func NewRateLimitHandler(limit int) *RateLimitHandler {
	return &RateLimitHandler{
		requestCounts: make(map[string]int),
		limit:         limit,
	}
}

func (r *RateLimitHandler) Handle(request *Request) *Response {
	fmt.Println("⏱️  Rate Limit Handler")

	count := r.requestCounts[request.UserID]
	if count >= r.limit {
		return &Response{
			StatusCode: 429,
			Body:       map[string]interface{}{"error": "Rate limit exceeded"},
		}
	}

	r.requestCounts[request.UserID]++
	fmt.Printf("✓ Rate limit check passed (%d/%d)\n", r.requestCounts[request.UserID], r.limit)
	return r.HandleNext(request)
}

type ValidationHandler struct {
	BaseHandler
}

func (v *ValidationHandler) Handle(request *Request) *Response {
	fmt.Println("✅ Validation Handler")

	if request.Body == nil {
		return &Response{
			StatusCode: 400,
			Body:       map[string]interface{}{"error": "Request body is required"},
		}
	}

	fmt.Println("✓ Validation passed")
	return v.HandleNext(request)
}

type LoggingHandler struct {
	BaseHandler
}

func (l *LoggingHandler) Handle(request *Request) *Response {
	start := time.Now()
	fmt.Printf("📝 Logging: Request from user %s at %s\n", request.UserID, start.Format("15:04:05"))

	response := l.HandleNext(request)

	duration := time.Since(start)
	fmt.Printf("📝 Logging: Response %d in %v\n", response.StatusCode, duration)

	return response
}

// ============= FUNCTIONAL MIDDLEWARE (Idiomatic Go) =============

type Middleware func(HandlerFunc) HandlerFunc
type HandlerFunc func(*Request) *Response

// Middleware functions
func AuthMiddleware(next HandlerFunc) HandlerFunc {
	return func(req *Request) *Response {
		fmt.Println("🔐 [Middleware] Authentication")

		if req.Token == "" || req.Token != "valid_token" {
			return &Response{
				StatusCode: 401,
				Body:       map[string]interface{}{"error": "Unauthorized"},
			}
		}

		fmt.Println("✓ Auth passed")
		return next(req)
	}
}

func RateLimitMiddleware(limit int) Middleware {
	counts := make(map[string]int)

	return func(next HandlerFunc) HandlerFunc {
		return func(req *Request) *Response {
			fmt.Println("⏱️  [Middleware] Rate Limiting")

			if counts[req.UserID] >= limit {
				return &Response{
					StatusCode: 429,
					Body:       map[string]interface{}{"error": "Too many requests"},
				}
			}

			counts[req.UserID]++
			fmt.Printf("✓ Rate limit OK (%d/%d)\n", counts[req.UserID], limit)
			return next(req)
		}
	}
}

func ValidationMiddleware(next HandlerFunc) HandlerFunc {
	return func(req *Request) *Response {
		fmt.Println("✅ [Middleware] Validation")

		if req.Body == nil {
			return &Response{
				StatusCode: 400,
				Body:       map[string]interface{}{"error": "Invalid request"},
			}
		}

		fmt.Println("✓ Validation passed")
		return next(req)
	}
}

func LoggingMiddleware(next HandlerFunc) HandlerFunc {
	return func(req *Request) *Response {
		start := time.Now()
		fmt.Printf("📝 [Middleware] Request from %s\n", req.UserID)

		resp := next(req)

		fmt.Printf("📝 [Middleware] Response %d in %v\n", resp.StatusCode, time.Since(start))
		return resp
	}
}

func RecoveryMiddleware(next HandlerFunc) HandlerFunc {
	return func(req *Request) *Response {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("🚨 Panic recovered: %v\n", r)
			}
		}()

		return next(req)
	}
}

// Chain multiple middlewares
func Chain(handler HandlerFunc, middlewares ...Middleware) HandlerFunc {
	// Apply middlewares in reverse order
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}

// Base handler
func FinalHandler(req *Request) *Response {
	fmt.Println("🎯 Final Handler - Processing request")
	return &Response{
		StatusCode: 200,
		Body: map[string]interface{}{
			"message": "Request processed successfully",
			"user":    req.UserID,
		},
	}
}

func main() {
	fmt.Println("=== Traditional Chain of Responsibility ===")

	// Build chain
	auth := &AuthenticationHandler{}
	rateLimit := NewRateLimitHandler(5)
	validation := &ValidationHandler{}
	logging := &LoggingHandler{}

	logging.SetNext(auth)
	auth.SetNext(rateLimit)
	rateLimit.SetNext(validation)

	// Test requests
	request1 := &Request{
		UserID: "user123",
		Token:  "valid_token",
		Body:   map[string]interface{}{"data": "test"},
	}

	response := logging.Handle(request1)
	fmt.Printf("Response: %+v\n", response)

	fmt.Println("\n=== Functional Middleware (Idiomatic Go) ===")

	// Compose middlewares
	handler := Chain(
		FinalHandler,
		RecoveryMiddleware,
		LoggingMiddleware,
		AuthMiddleware,
		RateLimitMiddleware(5),
		ValidationMiddleware,
	)

	// Test request
	request2 := &Request{
		UserID: "user456",
		Token:  "valid_token",
		Body:   map[string]interface{}{"action": "create"},
	}

	response2 := handler(request2)
	fmt.Printf("\nFinal Response: %+v\n", response2)

	// Test with invalid token
	fmt.Println("\n=== Testing with invalid token ===")
	request3 := &Request{
		UserID: "user789",
		Token:  "invalid",
		Body:   map[string]interface{}{"action": "create"},
	}

	response3 := handler(request3)
	fmt.Printf("Final Response: %+v\n", response3)
}
