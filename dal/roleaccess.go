package dal

import (
	"github.com/qtzheng/SIMP/app/modules"
	"gopkg.in/mgo.v2/bson"
)

//查询角色列表
func SelectRoles(page, size int) (*[]modules.Role, error) {
	roles := &[]modules.Role{}
	query := NewDB().C(RoleColl).Find(bson.M{}).Skip(page * size).Limit(size)
	err := query.All(roles)
	return roles, err
}
func SelectRoleTree() (*[]modules.Role, error) {
	roles := &[]modules.Role{}
	query := NewDB().C(RoleColl).Find(nil).Select(bson.M{"RoleID": 1, "RoleName": 1, "ParentID": 1})
	err := query.All(roles)
	return roles, err
}
func RoleInsert(role *modules.Role) error {
	err := NewDB().C(RoleColl).Insert(role)
	return err
}
