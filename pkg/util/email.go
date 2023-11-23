package util

import (
	"crypto/tls"
	"github.com/jordan-wright/email"
	"net/smtp"
)

func SendCodeEmail(toUserEmail, code string) error {
	e := email.NewEmail()
	e.From = "Get <cqk54041@163.com>"
	e.To = []string{toUserEmail}
	e.Subject = "雨竹家庭农场 注册信息"
	e.HTML = []byte("您的验证码：<b>" + code + "</b>")
	return e.SendWithTLS("smtp.163.com:465",
		smtp.PlainAuth("", "cqk54041@163.com", "KXUHFYPJHOFVHVRE", "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
}
