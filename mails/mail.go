package mails

import (
	"crypto/tls"
	"log"
	"os"

	"gopkg.in/gomail.v2"
)

var (
	// Message var
	Message *gomail.Message

	bugtrackerEmail string
	bugtrackerPw    string
)

// EmailInit func
func EmailInit() {
	bugtrackerEmail = os.Getenv("BUGTRACKER_EMAIL")
	bugtrackerPw = os.Getenv("BUGTRACKER_PW")
	if bugtrackerEmail == "" || bugtrackerPw == "" {
		log.Printf("%T, %v\n", bugtrackerEmail, bugtrackerEmail)
		log.Fatalln("BugTracker Email credentials empty - Please check your Environment Variables")

	}
}

// SendVerificationMail func
func SendVerificationMail(username, email, key string) {
	bodyText :=
		"Hey " + username + ", you registration needs just one more step.\n\n" +
			"Verification Key : " + key +
			" For activating your Account please click on the following link and paste your Verification Key : " +
			"http://localhost:8888/verification"

	m := gomail.NewMessage()
	m.SetHeader("From", bugtrackerEmail)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "BugTracker Verification")
	m.SetBody("text/html", bodyText)

	d := gomail.NewDialer("smtp.gmail.com", 587, bugtrackerEmail, bugtrackerPw)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

}
