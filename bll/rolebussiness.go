package bll

import (
	"github.com/qtzheng/SIMP/app/modules"
	"github.com/qtzheng/SIMP/dal"
)

type RoleBussiness struct {
	BaseBussiness
}

func (r *RoleBussiness) CreateRoleTree() *[]modules.Role {
	roles := *dal.RoleAccess.SelectRoleTree()
	return roles
}
func (r *RoleBussiness) Insert(role *modules.Role) error {
	return dal.RoleAccess.Insert(role)
}
