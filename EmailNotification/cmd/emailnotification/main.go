package main

import (
	"fmt"
	"log"
	"net/smtp"
)

func main() {
	// SMTP server configuration.
	smtpHost := "smtp.example.com"   // Replace with your SMTP server
	smtpPort := "587"                // Common SMTP ports: 587, 465, 25
	senderEmail := "you@example.com" // Sender email address
	senderPassword := "yourpassword" // Sender email password or app-specific password
	recipientEmail := "recipient@example.com"

	// Email content.
	subject := "Subject: Notification Alert\r\n"
	body := "Hi,\n\nThis is to notify you that a specific event has been triggered.\n\nRegards,\nNotifier"
	msg := []byte(subject + "\r\n" + body)

	// Authentication.
	auth := smtp.PlainAuth("", senderEmail, senderPassword, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, senderEmail, []string{recipientEmail}, msg)
	if err != nil {
		log.Fatal("Failed to send email:", err)
	}

	fmt.Println("Notification email sent successfully.")
}
