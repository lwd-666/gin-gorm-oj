package test

import (
	"crypto/tls"
	"github.com/jordan-wright/email"
	"net/smtp"
	"testing"
)

func TestSend(t *testing.T)  {
	e := email.NewEmail()
	e.From = "Get <getcharzhaopan@163.com>"
	//e.From = "Get <1124099628@qq.com>"
	e.To = []string{"17322701649@163.com"}
	e.Subject = "验证码发送测试"
	e.HTML = []byte("您的验证码是<b>123456</b>")
	//err :=e.Send("smtp.163.com:465", smtp.PlainAuth("", "getcharzhaopan@163.com", "XYQHBVISUAHXXRHQ", "smtp.163.com:465"))
	//返回EOF时，关闭SSL重试
	err := e.SendWithTLS("smtp.163.com:465", smtp.PlainAuth("", "getcharzhaopan@163.com", "XYQHBVISUAHXXRHQ", "smtp.163.com"),&tls.Config{ServerName: "smtp.163.com",InsecureSkipVerify: true})
	if err != nil {
		t.Fatal(err)
	}
}
