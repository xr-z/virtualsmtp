package store

import (
	"time"

	"gorm.io/gorm"
)

func New(orm *gorm.DB) Saver {
	orm.AutoMigrate(&msg{})
	return &simpleSaver{orm}
}

type simpleSaver struct {
	*gorm.DB
}

// "msgs" is the table name in the database.
// storage the email
type msg struct {
	ID       uint       `gorm:"primarykey"`
	From     string     `gorm:"from"`
	To       string     `gorm:"to"`
	Data     []byte     `gorm:"data"`
	Createat time.Time  `gorm:"createat"`
	Postat   *time.Time `gorm:"postat"`
}

func (s *simpleSaver) Save(from string, to string, data []byte) error {
	err := s.DB.Create(&msg{0, from, to, data, time.Now(), nil}).Error
	return err
}

func (s *simpleSaver) GetList(where string) []msg_alias {
	var msgs []msg
	s.DB.Where(where).Find(&msgs)

	var rs = make([]msg_alias, len(msgs))
	for i, m := range msgs {
		rs[i] = msg_alias(m)
	}
	return rs
}
