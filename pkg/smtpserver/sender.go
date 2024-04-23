package smtpserver

import (
	"net/smtp"
)

// SimpleSender is a simple implementation of the Sender interface.
type SimpleSender struct {
	DSN string

	*smtp.Client
}

func (s *SimpleSender) Open() error {
	c, err := smtp.Dial(s.DSN)
	if err != nil {
		return err
	}
	s.Client = c
	return nil
}

func (s *SimpleSender) Close() {
	if s.Client != nil {
		s.Client.Quit()
		s.Client.Close()
	}
}

func (s *SimpleSender) Send(m msg) error {
	if err := s.Client.Mail(m.From); err != nil {
		return err
	}
	if err := s.Client.Rcpt(m.To); err != nil {
		return err
	}

	// Send the email body.
	if wc, err := s.Client.Data(); err != nil {
		return err
	} else {
		if _, err = wc.Write(m.Data); err != nil {
			return err
		}

		if err = wc.Close(); err != nil {
			return err
		}
	}
	return nil
}
