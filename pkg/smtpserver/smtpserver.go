// 这个文件实现了 SMTP 服务器的功能。
// 它实现了 Backend 结构，该结构包含一个 Save 方法，该方法用于保存邮件。
// 它还实现了 NewBackend 方法，该方法用于创建 Backend 实例。
// 它还实现了 NewSession 方法，该方法用于创建 Session 实例。
// Session 结构包含一个 Backend 实例，From、To、RawData 字段，以及 Save 方法。
// 该方法用于保存邮件。
// 其他方法实现了 SMTP 服务器的功能。
package smtpserver

import (
	"io"
	"log"
	"time"

	"github.com/emersion/go-smtp"
)

// The Backend implements SMTP server methods.
type Backend struct {
	Save func(string, []string, []byte) //数据存储函数
}

func (bkd *Backend) NewBackend(fun func(string, []string, []byte)) *Backend {
	bkd.Save = fun
	return bkd
}

// NewSession is called after client greeting (EHLO, HELO).
func (bkd *Backend) NewSession(c *smtp.Conn) (smtp.Session, error) {
	return &Session{
		To:   []string{},
		Save: bkd.Save,
	}, nil
}

func (bkd *Backend) ListenAndServe(option ...func(*smtp.Server)) {
	s := smtp.NewServer(bkd)

	s.Addr = "localhost:1025"
	s.Domain = "localhost"
	s.WriteTimeout = 10 * time.Second
	s.ReadTimeout = 10 * time.Second
	s.MaxMessageBytes = 1024 * 1024
	s.MaxRecipients = 50
	s.AllowInsecureAuth = true

	for _, opt := range option {
		opt(s)
	}

	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

// A Session is returned after successful login.
type Session struct {
	From    string
	To      []string
	RawData []byte

	Save func(string, []string, []byte)
}

func (s *Session) Mail(from string, opts *smtp.MailOptions) error {
	s.From = from
	return nil
}

func (s *Session) Rcpt(to string, opts *smtp.RcptOptions) error {
	s.To = append(s.To, to)
	return nil
}

func (s *Session) Data(r io.Reader) error {
	if b, err := io.ReadAll(r); err != nil {
		return err
	} else {
		s.RawData = b
		if s.Save != nil {
			s.Save(s.From, s.To, s.RawData)
		}
	}
	return nil
}

func (s *Session) Reset() {}

func (s *Session) Logout() error {
	return nil
}
