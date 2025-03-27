package main

import "fmt"

// Defines an interface for creating objects but allows subclasses (factories) to decide which class to instantiate.
// Step 1: Define the Notification interface
type Notification interface {
	SendNotification()
}

// Step 2: Concrete implementations
type EmailNotification struct{}

func (e *EmailNotification) SendNotification() {
	fmt.Println("Sending email Notification")
}

type SMSNotification struct{}

func (s *SMSNotification) SendNotification() {
	fmt.Println("Sending SMS Notifocation")
}

// Step 3: Define Factory interface
type NotificationFactory interface {
	CreateNotification() Notification
}

// Step 4: Implement concrete factories
type EmailFactory struct{}

func (e *EmailFactory) CreateNotification() Notification {
	return &EmailNotification{}
}

type SMSFactory struct{}

func (s *SMSFactory) CreateNotification() Notification {
	return &SMSNotification{}
}

// Step 5: Usage
func main() {
	emailfactory := &EmailFactory{}
	email := emailfactory.CreateNotification()
	email.SendNotification()

	smsFactory := &SMSFactory{}
	sms := smsFactory.CreateNotification()
	sms.SendNotification()
}
