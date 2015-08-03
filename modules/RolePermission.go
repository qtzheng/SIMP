package modules

import (
	//"github.com/revel/revel"
	"gopkg.in/mgo.v2/bson"
)

type RolePermission struct {
	PermissionId bson.ObjectId `bson:"_id"`
	FunctionID   string
	ModuleID     string
	RoleID       string
	IsModule     bool
}
