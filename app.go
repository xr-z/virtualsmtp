package main

import (
	"fmt"
	"strings"

	"github.com/glebarez/sqlite"
	"github.com/zxr-cn/virtualsmtp/pkg/smtpserver"
	"github.com/zxr-cn/virtualsmtp/pkg/store"
	"gorm.io/gorm"
)

func main1() {
	var saver store.Saver
	backend := new(smtpserver.Backend).NewBackend(func(from string, to []string, data []byte) {
		saver.Save(from, strings.Join(to, ";"), data)
	})

	go backend.ListenAndServe()

}

// func forward(db *gorm.DB) {
// 	for {
// 		stor := store.New(db)
// 		list := stor.GetList("")
// 		if len(list) == 0 {
// 			continue
// 		}
// 		sender := smtpserver.SimpleSender{}

// 		for _, v := range list {
// 			if e := sender.Send(v); e == nil {

// 			}
// 			sender.Open()

// 		}
// 	}
// }

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	s := store.New(db)
	s.Save("a@a.a", "b@b.b", []byte("hello"))
	list := s.GetList("")
	for _, v := range list {
		fmt.Printf("v: %v\n", v)
	}

}
