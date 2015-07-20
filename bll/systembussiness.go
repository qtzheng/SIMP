package bll

import (
	"github.com/qtzheng/SIMP/app/modules"
	"github.com/qtzheng/SIMP/dal"
	"gopkg.in/mgo.v2/bson"
)

func RoleInit() {

}

//创建权限树
func RoleCreateTree() *[]modules.Role {
	roles, err := dal.RoleTreeSelect()
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

//编辑权限信息
func RoleEdit(role *modules.Role) error {
	return dal.RoleEdit(role)
}

//删除权限信息
func RoleDelete(id bson.ObjectId) error {
	return dal.RoleDelete(id)
}
func RoleInfo(id bson.ObjectId) (*modules.Role, error) {
	return dal.RoleInfo(id)
}
