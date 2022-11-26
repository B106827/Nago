package handlers

import (
    "bytes"
    "crypto/tls"
    "html/template"
    "NagoBackend/config"

    "gopkg.in/gomail.v2"
)

type Mail struct {}

func (*Mail) Send(to string, sub string, vars interface{}, files ...string) error {
    body, err := makeBody(vars, files...)
    if err != nil {
        return err
    }
    conf := config.GetConfig()

    mail := gomail.NewMessage()
    mail.SetHeader("From", conf.GetString("mail.from"))
    mail.SetHeader("To", to)
    mail.SetHeader("Subject", sub)
    mail.SetBody(conf.GetString("mail.contentType"), body)

    d := gomail.Dialer{Host: conf.GetString("mail.host"), Port: conf.GetInt("mail.port")}
    // SSL/TLS証明書がサーバー上で有効でない場合はtrue
    d.TLSConfig = &tls.Config{InsecureSkipVerify: conf.GetBool("mail.insecureSkipVerify")}
    err = d.DialAndSend(mail)
    return err
}

func makeBody(vars interface{}, files ...string) (string, error) {
    t, err := template.ParseFiles(files...)
    if err != nil {
        return "", err
    }
    buff := &bytes.Buffer{}
    t.Execute(buff, vars)
    return buff.String(), nil
}
