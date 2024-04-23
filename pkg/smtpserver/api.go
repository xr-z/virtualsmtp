package smtpserver

type msg = struct {
	ID   uint
	From string
	To   string
	Data []byte
}

type Sender interface {
	Open() error
	Close()
	Send([]msg) ([]any, error)
}
