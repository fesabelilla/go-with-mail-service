package mailservice

import (
	"fmt"
	"net/smtp"
)

func SendMail() {
	auth := smtp.PlainAuth(
		"fesabelilla",
		"fesabelilla.rokomari@gmail.com",
		"kdkmxkbajsrvfmlf",
		"smtp.gmail.com",
	)

	msgBody := "Subject:Test Mail\nHow are you ?"

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"fesabelilla.rokomari@gmail.com",
		[]string{"fesabelilla.rokomari@gmail.com"},
		[]byte(msgBody),
	)
	if err != nil {
		fmt.Println(err)
	}
}
