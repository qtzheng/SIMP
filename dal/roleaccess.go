package dal

import (
	"gopkg.in/mgo.v2"
)

type RoleAccess struct {
}

const (
	dbname = "Test"
)

//查询角色列表
func (r *RoleAccess) SelectRoles(session mgo.Session) {

}
func (r *RoleAccess) SelectRoleTree() {

}
