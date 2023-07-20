package main

import (
	"log"
	"net/smtp"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env: %s\n", err)
	}

	var (
		authUserName = os.Getenv("SES_AUTH_USERNAME") 
		authPassword = os.Getenv("SES_AUTH_PASSWORD")
		smtpServerAddr = os.Getenv("SES_SMTP_SERVER_ADDR")
		smtpServerPort = os.Getenv("SES_SMTP_SERVER_PORT")
	)
	
//	fmt.Printf("authUserName -> %s\nauthPassword -> %s\nsmtpServerAddr -> %s\nsmtpServerPort -> %s\n\n", authUserName, authPassword,smtpServerAddr, smtpServerPort)	

	// Note: If your SES account is in a sandbox you need to verify both sender and recipient emails
	destinationEmails := []string{"test@email.com"}
	sourceEmail := "test@email.com"
	
	// the message
	msg := []byte("This is a test message from SES and net/smtp.")
	
	// authorization
	auth := smtp.PlainAuth("", authUserName, authPassword, smtpServerAddr)
	
	// send email
	err_send := smtp.SendMail(smtpServerAddr+":"+smtpServerPort, auth, sourceEmail, destinationEmails, msg)

	if err_send != nil {
		log.Printf("There was an error sending the email: %s\n", err_send)
		return
	}
	
	log.Printf("Successfully sent the email.")	
}
