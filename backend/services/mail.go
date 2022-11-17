package services

import (
    "bytes"
    "crypto/tls"
    "html/template"
    "NagoBackend/config"

    "gopkg.in/gomail.v2"
)

type Mail struct {
    contentType        string
    from               string
    host               string
    port               int
    insecureSkipVerify bool
}

func NewMail() *Mail {
    contentType        := "text/plain"
    c                  := config.GetConfig()
    from               := c.GetString("mail.from")
    host               := c.GetString("mail.host")
    port               := c.GetInt("mail.port")
    insecureSkipVerify := c.GetBool("mail.insecureSkipVerify")

    mail := Mail{
        contentType,
        from,
        host,
        port,
        insecureSkipVerify,
    }
    return &mail
}

func (base *Mail) Send(to string, sub string, vars interface{}, files ...string) error {
    body, err := base.makeBody(vars, files...)
    if err != nil {
        return err
    }
    m := gomail.NewMessage()
    m.SetHeader("From", base.from)
    m.SetHeader("To", to)
    m.SetHeader("Subject", sub)
    m.SetBody(base.contentType, body)

    d := gomail.Dialer{Host: base.host, Port: base.port}
    // SSL/TLS証明書がサーバー上で有効でない場合はtrue
    d.TLSConfig = &tls.Config{InsecureSkipVerify: base.insecureSkipVerify}
    err = d.DialAndSend(m)
    return err
}

func (base *Mail) makeBody(vars interface{}, files ...string) (string, error) {
    t, err := template.ParseFiles(files...)
    if err != nil {
        return "", err
    }
    buff := &bytes.Buffer{}
    t.Execute(buff, vars)
    return buff.String(), nil
}
