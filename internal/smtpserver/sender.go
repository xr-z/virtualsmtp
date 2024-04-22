package smtpserver

import (
	"net/smtp"
)

type Sender interface {
	Open() error
	Close()
	Send()
}

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

func (s *SimpleSender) Send(msgs []struct {
	from string
	to   string
	data []byte
}) error {
	for _, m := range msgs {
		// Set the sender and recipient first
		if err := s.Client.Mail(m.from); err != nil {
			return err
		}
		if err := s.Client.Rcpt(m.to); err != nil {
			return err
		}

		// Send the email body.
		if wc, err := s.Client.Data(); err != nil {
			return err
		} else {
			if _, err = wc.Write(m.data); err != nil {
				return err
			}

			if err = wc.Close(); err != nil {
				return err
			}
		}
	}

	return nil
}
