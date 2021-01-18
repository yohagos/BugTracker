package mails

import (
	"log"
	"net/smtp"

	"../utils"
)

func test() {
	log.Println(utils.CreateTimeStamp())
}

// SendVerificationMail func
func SendVerificationMail(user, adress, key string) {
	from := "********"
	pass := "********"
	to := adress

	msg :=
		"\nFrom: WMD Postman\n" +
			"To: " + adress + "\n" +
			"Subject: Verifying your Account\n" +
			"Hey " + user + ", you registration needs just one more step.\n\n" +
			"Verification Key : " + key +
			"\n\nFor activating your Account please click on the following link and paste your Verification Key : \n\n" +
			"http://localhost:8888/verification"

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))
	if err != nil {
		log.Println(err)
	}
}
