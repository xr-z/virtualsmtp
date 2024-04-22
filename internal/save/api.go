package save

type msg_alias = struct {
	From string
	To   string
	Data []byte
}

type Saver interface {
	Save(from string, to string, data []byte) error
	GetList() []msg_alias
}
