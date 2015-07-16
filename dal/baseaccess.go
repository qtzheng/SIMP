package dal

import (
	"github.com/revel/revel"
	"gopkg.in/mgo.v2"
	"sync"
)

const (
	url      = ""
	DbName   = "DB_SIMP"
	RoleColl = "C_Role"
	DepColl  = "C_Department"
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
	}
}

func CloneDB() *mgo.Database {
	return MgoSession.Clone().DB(DbName)
}
func NewDB() *mgo.Database {
	return MgoSession.New().DB(DbName)
}
func CopyDB() *mgo.Database {
	return MgoSession.Copy().DB(DbName)
}
func CloseDB() {
	if MgoSession != nil {
		MgoSession.Close()
	}

}
