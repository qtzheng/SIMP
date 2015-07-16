package bll

import (
	"github.com/qtzheng/SIMP/app/modules"
	"github.com/qtzheng/SIMP/dal"
)

//创建权限树
func RoleCreateTree() *[]modules.Role {
	roles, err := dal.SelectRoleTree()
	if err != nil {
		ErrorLog(err)
	}
	return roles
}

//插入角色信息
func RoleInsert(role *modules.Role) error {
	role.RoleID = NewObjectID()
	return dal.RoleInsert(role)
}
