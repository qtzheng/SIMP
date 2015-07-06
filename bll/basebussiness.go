package bll

import (
	"fmt"
	"github.com/qtzheng/SIMP/utils"
	"github.com/revel/revel"
	"gopkg.in/mgo.v2"
)

const (
	url = ""
)

var (
//MgoConn *mgo.Session
)

func init() {
	MgoConn, err := mgo.Dial(url)
	if err != nil {
		if revel.DevMode {
			utils.Display(err)
		} else {
			panic(err)
		}
	} else {
		fmt.Print("数据库链接成功！")
		dbName, _ := MgoConn.DatabaseNames()
		fmt.Print(dbName)
	}
}

type BaseBussiness struct {
}