package main

import (
	"fmt"
)

// A single function (Factory Method) returns different object types based on input parameters.
// Step 1: Define an interface
type Notification interface {
	SendNotification()
}

// Step 2: Implement concrete classes
type EmailNotification struct{}

func (*EmailNotification) SendNotification() {
	fmt.Println("Send Email notification")
}

type SMSNotification struct{}

func (*SMSNotification) SendNotification() {
	fmt.Println("Send SMS notification")
}

// Step 3: Create a Simple Factory Function
func NotificationFactory(notificationType string) Notification {
	switch notificationType {
	case "email":
		return &EmailNotification{}
	case "sms":
		return &SMSNotification{}
	default:
		return nil
	}
}

// Step 4: Usage
func main() {
	email := NotificationFactory("email")
	email.SendNotification()

	sms := NotificationFactory("sms")
	sms.SendNotification()
}
