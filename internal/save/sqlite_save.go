package save

import (
	"gorm.io/gorm"
)

func New(orm *gorm.DB) Saver {
	orm.AutoMigrate(&msg{})
	return &simpleSaver{orm}
}

type simpleSaver struct {
	*gorm.DB
}

type msg struct {
	From string `gorm:"from"`
	To   string `gorm:"to"`
	Data []byte `gorm:"data"`
}

func (s *simpleSaver) Save(from string, to string, data []byte) error {
	err := s.DB.Create(msg{from, to, data}).Error
	return err
}

func (s *simpleSaver) GetList() []msg_alias {
	var msgs []msg
	s.DB.Find(&msgs)

	var rs = make([]msg_alias, len(msgs))
	for i, m := range msgs {
		rs[i] = msg_alias(m)
	}
	return rs
}
