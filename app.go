package main

import (
	"strings"

	"github.com/glebarez/sqlite"
	"github.com/zxr-cn/virtualsmtp/internal/save"
	"github.com/zxr-cn/virtualsmtp/internal/smtpserver"
	"gorm.io/gorm"
)

func main1() {
	var saver save.Saver
	backend := new(smtpserver.Backend).NewBackend(func(from string, to []string, data []byte) {
		saver.Save(from, strings.Join(to, ";"), data)
	})

	go backend.ListenAndServe()

}

func forward() {
	for {
		sender := smtpserver.SimpleSender{}
		sender.Open()

	}
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	s := save.New(db)
	// s.Save("a@a.a", "b@b.b", []byte("hello"))
	list := s.GetList()
	for _, v := range list {
		println(v.From, v.To, string(v.Data))
	}

}
