﻿package dal

import (
	"github.com/revel/revel"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"sync"
)

const (
	url      = ""
	DbName   = "DB_SIMP"
	RoleColl = "C_Role"
)

var (
	MgoSession *mgo.Session
	DevMode    bool
	mux        sync.Mutex
)

func init() {
	mux.Lock()
	defer mux.Unlock()
	if MgoSession != nil {
		return
	}
	DevMode = revel.DevMode
	var err error
	MgoSession, err = mgo.Dial(url)
	if err != nil {
		revel.WARN.Fatalln(err)
	} else {
		fmt.Print("数据库链接成功！")
		dbName, _ := MgoSession.DatabaseNames()
		fmt.Print(dbName)
	}
}

type BaseAccess struct {
}

func (b *BaseAccess) CloneDB() *mgo.Database {
	return MgoSession.Clone().DB(DbName)
}
func (b *BaseAccess) NewDB() *mgo.Database {
	return MgoSession.New().DB(DbName)
}
func (b *BaseAccess) CopyDB() *mgo.Database {
	return MgoSession.Copy().DB(DbName)
}
func (b *BaseAccess) CloseDB() {
	return MgoSession.Close()
}
