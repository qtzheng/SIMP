package modules

import (
	//"github.com/revel/revel"
	"gopkg.in/mgo.v2/bson"
)

type RolePermission struct {
	PermissionId bson.ObjectId `bson:"_id"`
	FunctionCode string
	FunctionID   string
	ModuleCode   string
	ModuleID     string
	RoleCode     string
	RoleID       string
	ParentPerID  string
	IsModule     bool
	//是否引用：true:子节点点选中，自己没选中；false:自己选中
	IsRef bool
}
