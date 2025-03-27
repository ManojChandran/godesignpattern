package main

import (
	"fmt"
)

// A factory of factories → Provides an interface for creating a family of related objects.
// Step 1: Define Notification Interface
type Notification interface {
	SendNotification()
}

// Step 2: Implement Concrete Classes
type EmailNotification struct{}

func (e *EmailNotification) SendNotification() {
	fmt.Println("Sending Email Notification")
}

type SMSNotification struct{}

func (s *SMSNotification) SendNotification() {
	fmt.Println("Sending SMS Notification")
}

// Step 3: Create an Abstract Factory Interface
type NotificationFactory interface {
	CreateEmailNotification() Notification
	CreateSMSNotification() Notification
}

// Step 4: Implement Concrete Factory
type ConcreteNotificationFactory struct{}

func (c *ConcreteNotificationFactory) CreateEmailNotification() Notification {
	return &EmailNotification{}
}

func (c *ConcreteNotificationFactory) CreateSMSNotification() Notification {
	return &SMSNotification{}
}

// Step 5: Usage
func main() {
	factory := &ConcreteNotificationFactory{}

	email := factory.CreateEmailNotification()
	email.SendNotification()

	sms := factory.CreateSMSNotification()
	sms.SendNotification()

}
