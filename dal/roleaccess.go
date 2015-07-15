package dal

import (
	"github.com/qtzheng/SIMP/app/modules"
	"gopkg.in/mgo.v2/bson"
)

type RoleAccess struct {
	BaseAccess
}

//查询角色列表
func (r *RoleAccess) SelectRoles() {
	r.NewDB().C(RoleColl).Find(bson.M{})
}
func (r *RoleAccess) SelectRoleTree() *[]modules.Role {
	roles := &[]modules.Role{}
	query := r.NewDB().C(RoleColl).Find(nil).Select(bson.M{"RoleID": 1, "RoleName": 1, "ParentID": 1})
	query.All(roles)
	return roles
}
func (r *RoleAccess) Insert(role *modules.Role) error {
	err := r.NewDB().C(RoleColl).Insert(role)
	return err
}
