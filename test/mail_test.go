package test

import (
	"crypto/tls"
	"github.com/jordan-wright/email"
	"net/smtp"
	"testing"
)

func TestSendMail(t *testing.T) {
	e := email.NewEmail()
	e.From = "cjy <2812485760@qq.com>"
	e.To = []string{"cjy@cjy.ink"}
	e.Subject = "验证码测试"
	e.HTML = []byte("验证码是：<h1>123456</h1>")
	err := e.SendWithTLS("smtp.qq.com:465", smtp.PlainAuth("", "2812485760@qq.com", "ugencsnkcxyadcif", "smtp.qq.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.qq.com"})
	if err != nil {
		t.Fatal(err)
	}
}
