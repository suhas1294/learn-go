package main

import "fmt"

// ============= INTERFACES (Implicit Implementation) =============

// PaymentStrategy - no "implements" keyword needed
type PaymentStrategy interface {
	Process(amount float64, details PaymentDetails) PaymentResult
}

type PaymentDetails struct {
	CustomerID string
	CardToken  string
	UpiID      string
}

type PaymentResult struct {
	Success       bool
	TransactionID string
}

// ============= CONCRETE STRATEGIES =============

// StripePayment - automatically implements PaymentStrategy
type StripePayment struct {
	apiKey string
}

func NewStripePayment(apiKey string) *StripePayment {
	return &StripePayment{apiKey: apiKey}
}

func (s *StripePayment) Process(amount float64, details PaymentDetails) PaymentResult {
	// Stripe API integration
	fmt.Printf("Processing $%.2f via Stripe\n", amount)
	return PaymentResult{
		Success:       true,
		TransactionID: "stripe_" + generateID(),
	}
}

// PayPalPayment
type PayPalPayment struct {
	clientID     string
	clientSecret string
}

func NewPayPalPayment(clientID, clientSecret string) *PayPalPayment {
	return &PayPalPayment{
		clientID:     clientID,
		clientSecret: clientSecret,
	}
}

func (p *PayPalPayment) Process(amount float64, details PaymentDetails) PaymentResult {
	fmt.Printf("Processing $%.2f via PayPal\n", amount)
	return PaymentResult{
		Success:       true,
		TransactionID: "paypal_" + generateID(),
	}
}

// RazorpayPayment
type RazorpayPayment struct {
	keyID     string
	keySecret string
}

func NewRazorpayPayment(keyID, keySecret string) *RazorpayPayment {
	return &RazorpayPayment{
		keyID:     keyID,
		keySecret: keySecret,
	}
}

func (r *RazorpayPayment) Process(amount float64, details PaymentDetails) PaymentResult {
	fmt.Printf("Processing ₹%.2f via Razorpay\n", amount)
	return PaymentResult{
		Success:       true,
		TransactionID: "razorpay_" + generateID(),
	}
}

// ============= CONTEXT (Uses Strategy) =============

type PaymentProcessor struct {
	strategy PaymentStrategy
}

func NewPaymentProcessor(strategy PaymentStrategy) *PaymentProcessor {
	return &PaymentProcessor{strategy: strategy}
}

func (p *PaymentProcessor) SetStrategy(strategy PaymentStrategy) {
	p.strategy = strategy
}

func (p *PaymentProcessor) ExecutePayment(amount float64, details PaymentDetails) PaymentResult {
	return p.strategy.Process(amount, details)
}

// ============= FUNCTIONAL APPROACH (Alternative) =============

// Go supports first-class functions - can use functional approach too
type PaymentFunc func(amount float64, details PaymentDetails) PaymentResult

type FunctionalPaymentProcessor struct {
	processFn PaymentFunc
}

func NewFunctionalProcessor(fn PaymentFunc) *FunctionalPaymentProcessor {
	return &FunctionalPaymentProcessor{processFn: fn}
}

func (f *FunctionalPaymentProcessor) ExecutePayment(amount float64, details PaymentDetails) PaymentResult {
	return f.processFn(amount, details)
}

// Usage with closures
func stripeProcessorFunc(apiKey string) PaymentFunc {
	return func(amount float64, details PaymentDetails) PaymentResult {
		fmt.Printf("Processing $%.2f via Stripe (functional)\n", amount)
		return PaymentResult{Success: true, TransactionID: "stripe_" + generateID()}
	}
}

func generateID() string {
	return "123456"
}

func main() {
	// ========== OOP APPROACH ==========
	processor := NewPaymentProcessor(NewStripePayment("sk_test_123"))
	result := processor.ExecutePayment(100.50, PaymentDetails{CustomerID: "cust_123"})
	fmt.Printf("Result: %+v\n", result)

	// Change strategy at runtime
	processor.SetStrategy(NewRazorpayPayment("rzp_key", "rzp_secret"))
	result = processor.ExecutePayment(5000, PaymentDetails{CustomerID: "cust_456"})
	fmt.Printf("Result: %+v\n", result)

	// ========== FUNCTIONAL APPROACH ==========
	funcProcessor := NewFunctionalProcessor(stripeProcessorFunc("sk_test_123"))
	result = funcProcessor.ExecutePayment(200.00, PaymentDetails{})
	fmt.Printf("Result: %+v\n", result)
}
