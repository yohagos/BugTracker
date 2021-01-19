package mails

import (
	"fmt"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// SendVerificationMail func
func SendVerificationMail(user, adress, key string) {

}

// SendMail func
func SendMail() {
	from := mail.NewEmail("Me", "bugtracker2021@gmail.com")
	to := mail.NewEmail("Test", "bugtracker2021@gmail.com")
	subject := "Sending Test - Twilio SendGrid.."
	plainTextContent := "and easy to do anywhere, even with Go"
	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}

}
