package bll

import (
	"fmt"
	//"github.com/qtzheng/SIMP/utils"
	"github.com/revel/revel"
	"gopkg.in/mgo.v2"
)

const (
	url      = ""
	DbName   = "DB_SIMP"
	RoleColl = "C_Role"
)

var (
	MgoConn *mgo.Session
	DevMode bool
)

func init() {
	DevMode = revel.DevMode
	var err error
	MgoConn, err = mgo.Dial(url)
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

func (b *BaseBussiness) CloneDB() *mgo.Database {
	return MgoConn.Clone().DB(DbName)
}
func (b *BaseBussiness) NewDB() *mgo.Database {
	return MgoConn.New().DB(DbName)
}
