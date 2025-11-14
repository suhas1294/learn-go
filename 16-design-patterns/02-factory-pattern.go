package main

import (
	"errors"
	"fmt"
)

// ============= INTERFACE =============

type Notifier interface {
	Send(message string) error
}

// ============= CONCRETE IMPLEMENTATIONS =============

type EmailNotifier struct {
	smtpHost string
	smtpPort int
	from     string
}

func (e *EmailNotifier) Send(message string) error {
	fmt.Printf("Sending email from %s: %s\n", e.from, message)
	return nil
}

type SMSNotifier struct {
	provider string
	apiKey   string
}

func (s *SMSNotifier) Send(message string) error {
	fmt.Printf("Sending SMS via %s: %s\n", s.provider, message)
	return nil
}

type PushNotifier struct {
	fcmServerKey string
}

func (p *PushNotifier) Send(message string) error {
	fmt.Printf("Sending push notification: %s\n", message)
	return nil
}

type SlackNotifier struct {
	webhookURL string
	channel    string
}

func (s *SlackNotifier) Send(message string) error {
	fmt.Printf("Sending to Slack #%s: %s\n", s.channel, message)
	return nil
}

// ============= FACTORY FUNCTION (Idiomatic Go) =============

type NotifierType string

const (
	Email NotifierType = "email"
	SMS   NotifierType = "sms"
	Push  NotifierType = "push"
	Slack NotifierType = "slack"
)

// Factory function with options
type NotifierConfig struct {
	Type         NotifierType
	SMTPHost     string
	SMTPPort     int
	EmailFrom    string
	SMSProvider  string
	APIKey       string
	FCMKey       string
	WebhookURL   string
	SlackChannel string
}

func NewNotifier(config NotifierConfig) (Notifier, error) {
	switch config.Type {
	case Email:
		return &EmailNotifier{
			smtpHost: config.SMTPHost,
			smtpPort: config.SMTPPort,
			from:     config.EmailFrom,
		}, nil
	case SMS:
		return &SMSNotifier{
			provider: config.SMSProvider,
			apiKey:   config.APIKey,
		}, nil
	case Push:
		return &PushNotifier{
			fcmServerKey: config.FCMKey,
		}, nil
	case Slack:
		return &SlackNotifier{
			webhookURL: config.WebhookURL,
			channel:    config.SlackChannel,
		}, nil
	default:
		return nil, errors.New("unknown notifier type")
	}
}

// ============= FACTORY WITH REGISTRY (Advanced Pattern) =============

type NotifierFactory func(config NotifierConfig) (Notifier, error)

var notifierRegistry = make(map[NotifierType]NotifierFactory)

func RegisterNotifier(notifierType NotifierType, factory NotifierFactory) {
	notifierRegistry[notifierType] = factory
}

func CreateNotifier(notifierType NotifierType, config NotifierConfig) (Notifier, error) {
	factory, exists := notifierRegistry[notifierType]
	if !exists {
		return nil, fmt.Errorf("notifier type %s not registered", notifierType)
	}
	return factory(config)
}

func init() {
	// Register notifiers at startup
	RegisterNotifier(Email, func(config NotifierConfig) (Notifier, error) {
		return &EmailNotifier{
			smtpHost: config.SMTPHost,
			smtpPort: config.SMTPPort,
			from:     config.EmailFrom,
		}, nil
	})

	RegisterNotifier(SMS, func(config NotifierConfig) (Notifier, error) {
		return &SMSNotifier{
			provider: config.SMSProvider,
			apiKey:   config.APIKey,
		}, nil
	})
}

func main() {
	// Simple factory
	notifier, err := NewNotifier(NotifierConfig{
		Type:      Email,
		SMTPHost:  "smtp.gmail.com",
		SMTPPort:  587,
		EmailFrom: "noreply@example.com",
	})
	if err != nil {
		panic(err)
	}
	notifier.Send("Hello from factory!")

	// Registry-based factory
	smsNotifier, _ := CreateNotifier(SMS, NotifierConfig{
		Type:        SMS,
		SMSProvider: "Twilio",
		APIKey:      "twilio_key",
	})
	smsNotifier.Send("Hello from registry!")
}
