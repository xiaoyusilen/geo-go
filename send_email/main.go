// author by @xiaoyusilen

package main

import (
	"crypto/tls"

	"github.com/xiaoyusilen/geo-go/send_email/config"

	gomail "gopkg.in/gomail.v2"
)

func SendMail() {

	d := gomail.NewDialer(config.DefaultEmailHost, 25, config.DefaultEmailUser, config.DefaultEmailPass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	m := gomail.NewMessage()
	m.SetHeader("From", "name@example.com")
	m.SetHeader("To", "debug@xiaoyu.fail")
	m.SetHeader("Subject", "Hello")

	m.SetBody("text/html", "Hello <b>xiaoyusilen~</b>")

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

func main() {

	SendMail()
}
