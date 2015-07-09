package dal

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type RoleAccess struct {
	BaseAccess
}

const (
	dbname = "Test"
)

//查询角色列表
func (r *RoleAccess) SelectRoles() {
	r.NewDB().C(RoleColl).Find(bson.M{})
}
func (r *RoleAccess) SelectRoleTree() {
	r.NewDB().C(RoleColl).Find(bson.M{}).Select(bson.s)
}
