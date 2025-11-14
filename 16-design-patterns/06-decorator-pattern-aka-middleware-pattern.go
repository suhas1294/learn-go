package designpatternsingo
package main

import (
	"fmt"
	"time"
)

// ============= TRADITIONAL DECORATOR =============

type DataSource interface {
	WriteData(data string) error
	ReadData() (string, error)
}

type FileDataSource struct {
	filename string
}

func NewFileDataSource(filename string) *FileDataSource {
	return &FileDataSource{filename: filename}
}

func (f *FileDataSource) WriteData(data string) error {
	fmt.Printf("Writing to file %s: %s\n", f.filename, data)
	return nil
}

func (f *FileDataSource) ReadData() (string, error) {
	return "file data", nil
}

// Encryption decorator
type EncryptionDecorator struct {
	wrappee DataSource
}

func NewEncryptionDecorator(source DataSource) *EncryptionDecorator {
	return &EncryptionDecorator{wrappee: source}
}

func (e *EncryptionDecorator) WriteData(data string) error {
	encrypted := e.encrypt(data)
	return e.wrappee.WriteData(encrypted)
}

func (e *EncryptionDecorator) ReadData() (string, error) {
	data, err := e.wrappee.ReadData()
	if err != nil {
		return "", err
	}
	return e.decrypt(data), nil
}

func (e *EncryptionDecorator) encrypt(data string) string {
	return fmt.Sprintf("encrypted(%s)", data)
}

func (e *EncryptionDecorator) decrypt(data string) string {
	return data // Simplified
}

// Compression decorator
type CompressionDecorator struct {
	wrappee DataSource
}

func NewCompressionDecorator(source DataSource) *CompressionDecorator {
	return &CompressionDecorator{wrappee: source}
}

func (c *CompressionDecorator) WriteData(data string) error {
	compressed := c.compress(data)
	return c.wrappee.WriteData(compressed)
}

func (c *CompressionDecorator) ReadData() (string, error) {
	data, err := c.wrappee.ReadData()
	if err != nil {
		return "", err
	}
	return c.decompress(data), nil
}

func (c *CompressionDecorator) compress(data string) string {
	return fmt.Sprintf("compressed(%s)", data)
}

func (c *CompressionDecorator) decompress(data string) string {
	return data // Simplified
}

// ============= FUNCTIONAL DECORATOR (Middleware - Idiomatic Go) =============

type DataHandler func(data string) (string, error)

// Middleware functions
func WithEncryption(next DataHandler) DataHandler {
	return func(data string) (string, error) {
		fmt.Println("Encrypting data...")
		encrypted := fmt.Sprintf("encrypted(%s)", data)
		result, err := next(encrypted)
		if err != nil {
			return "", err
		}
		fmt.Println("Decrypting data...")
		return result, nil
	}
}

func WithCompression(next DataHandler) DataHandler {
	return func(data string) (string, error) {
		fmt.Println("Compressing data...")
		compressed := fmt.Sprintf("compressed(%s)", data)
		result, err := next(compressed)
		if err != nil {
			return "", err
		}
		fmt.Println("Decompressing data...")
		return result, nil
	}
}

func WithLogging(next DataHandler) DataHandler {
	return func(data string) (string, error) {
		start := time.Now()
		fmt.Printf("Processing started at %s\n", start.Format("15:04:05"))
		
		result, err := next(data)
		
		fmt.Printf("Processing completed in %v\n", time.Since(start))
		return result, err
	}
}

func WithRetry(maxRetries int) func(DataHandler) DataHandler {
	return func(next DataHandler) DataHandler {
		return func(data string) (string, error) {
			var err error
			var result string
			
			for i := 0; i < maxRetries; i++ {
				result, err = next(data)
				if err == nil {
					return result, nil
				}
				fmt.Printf("Retry %d/%d\n", i+1, maxRetries)
			}
			
			return "", fmt.Errorf("failed after %d retries: %w", maxRetries, err)
		}
	}
}

// Base handler
func BaseDataHandler(data string) (string, error) {
	fmt.Printf("Base handler processing: %s\n", data)
	return data, nil
}

// Chain decorators
func ChainMiddleware(handler DataHandler, middlewares ...func(DataHandler) DataHandler) DataHandler {
	// Apply middlewares in reverse order (like HTTP middleware)
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}

func main() {
	fmt.Println("=== Traditional Decorator ===")
	source := NewFileDataSource("data.txt")
	source = NewEncryptionDecorator(source)
	source = NewCompressionDecorator(source)
	
	source.WriteData("sensitive data")

	fmt.Println("\n=== Functional Decorator (Middleware) ===")
	
	// Compose middlewares
	handler := ChainMiddleware(
		BaseDataHandler,
		WithLogging,
		WithEncryption,
		WithCompression,
		WithRetry(3),
	)
	
	result, err := handler("sensitive data")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Result: %s\n", result)
	}

	fmt.Println("\n=== HTTP Middleware Example ===")
	// This is how HTTP middleware works in Go
	type HTTPHandler func(request string) string

	authMiddleware := func(next HTTPHandler) HTTPHandler {
		return func(request string) string {
			fmt.Println("Checking authentication...")
			if request == "" {
				return "Unauthorized"
			}
			return next(request)
		}
	}

	loggingMiddleware := func(next HTTPHandler) HTTPHandler {
		return func(request string) string {
			fmt.Printf("Request: %s\n", request)
			response := next(request)
			fmt.Printf("Response: %s\n", response)
			return response
		}
	}

	baseHandler := func(request string) string {
		return fmt.Sprintf("Processed: %s", request)
	}

	// Chain HTTP middlewares
	handler2 := authMiddleware(loggingMiddleware(baseHandler))
	response := handler2("GET/api/users")
	fmt.Printf("Final response: %s\n", response)
}