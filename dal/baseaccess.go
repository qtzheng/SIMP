package dal

import (
	"github.com/revel/revel"
	"gopkg.in/mgo.v2"
	"sync"
)

const (
	url           = ""
	DbName        = "db_simp"
	RoleColl      = "c_role"
	DepColl       = "c_department"
	UserColl      = "c_user"
	ModuleColl    = "c_module"
	FuncColl      = "c_func"
	RolePermiColl = "c_rolepermission"
	MachineColl   = "c_machine"
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
		revel.WARN.Panicln(err)
	}
	err = roleInit()
	if err != nil {
		revel.WARN.Panicln(err)
	}
	err = depInit()
	if err != nil {
		revel.WARN.Panicln(err)
	}
	err = userInit()
	if err != nil {
		revel.WARN.Panicln(err)
	}
	err = moduleInit()
	if err != nil {
		revel.WARN.Panicln(err)
	}
}

func CloneDB() *mgo.Database {
	return MgoSession.Clone().DB(DbName)
}
func NewDB() *mgo.Database {
	return MgoSession.Clone().DB(DbName)
}
func CopyDB() *mgo.Database {
	return MgoSession.Copy().DB(DbName)
}
func CloseDB() {
	if MgoSession != nil {
		MgoSession.Close()
	}

}
