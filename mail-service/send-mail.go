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

	msgBody := "Hi, Tapu vai. How are you ?"

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"fesabelilla.rokomari@gmail.com",
		[]string{"tapu@rokomari.com"},
		[]byte(msgBody),
	)
	if err != nil {
		fmt.Println(err)
	}
}
