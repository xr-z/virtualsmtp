package store

import "time"

type msg_alias = struct {
	ID       uint
	From     string
	To       string
	Data     []byte
	Createat time.Time  `gorm:"createat"`
	Postat   *time.Time `gorm:"postat"`
}

type Saver interface {
	Save(from string, to string, data []byte) error // save message to storage
	GetList(where string) []msg_alias               // get list of messages from storage
}
