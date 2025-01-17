package sender

import (
	"encoding/json"
	"errors"

	"gopkg.in/gomail.v2"
)

type Email struct {
	User     string `json:"user"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
}

func init() {
	register("email", "邮箱", NewEmail)
}

// NewEmail 初始化邮件发送器
func NewEmail(opts ...ISenderInitOptionFunc) (ISender, error) {
	options := &ISenderInitFuncOption{}
	for _, opt := range opts {
		opt(options)
	}

	if options.extra == "" {
		return nil, errors.New("extra info in empty")
	}
	email := &Email{User: options.ak, Password: options.sk}
	if err := json.Unmarshal([]byte(options.extra), email); err != nil {
		return nil, err
	}
	return email, nil
}

// Send 发送邮件
func (s *Email) Send(email, title, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(s.User, s.Name))
	m.SetHeader("To", m.FormatAddress(email, email))
	m.SetHeader("Subject", title)
	m.SetBody("text/html; charset=UTF-8", body)
	d := gomail.NewDialer(
		s.Host,
		s.Port,
		s.User,
		s.Password,
	)
	return d.DialAndSend(m)
}
