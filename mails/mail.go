package mails

import (
	"net/mail"
	"strings"
)

func encodeRFC2047(String string) string {
	// use mail's rfc2047 to encode any string
	addr := mail.Address{String, ""}
	return strings.Trim(addr.String(), " <>")
}

// SMTPServer func
func SMTPServer(username, email, key string) {
	/* admin := "bugtracker2021@gmail.com"
	pass := "+++++++++"

	smtpServer := "smtp.gmail.com"
	auth := smtp.PlainAuth(
		"",
		admin,
		pass,
		smtpServer,
	)

	from := mail.Address{"Admin", admin}
	to := mail.Address{username, email}
	title := "Verification Key"

	body := "Hey " + username + ", you registration needs just one more step.\n\n" +
		"Verification Key : " + key +
		"\n\nFor activating your Account please click on the following link and paste your Verification Key : \n\n" +
		"http://localhost:8888/verification"

	header := make(map[string]string)
	header["From"] = from.String()
	header["To"] = to.String()
	header["Subject"] = encodeRFC2047(title)
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(body))

	err := smtp.SendMail(
		smtpServer+":587",
		auth,
		from.Address,
		[]string{to.Address},
		[]byte(message),
	)

	if err != nil {
		log.Fatalln(err)
	} */
}

// SendVerificationMail func
func SendVerificationMail(user, adress, key string) {
	/*	from := "bugtracker2021@gmail.com"
		pass := ""
		to := adress

		msg := "\nFrom: WMD Postman\n" +
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
		}*/
}

/* // SendMail func
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

} */
