package bll

import (
	"github.com/qtzheng/SIMP/dal"
	"github.com/qtzheng/SIMP/modules"
	"gopkg.in/mgo.v2/bson"
	"strings"
)

func RoleInit() {

}

//创建权限树
func RoleCreateTree() (*[]modules.Role, error) {
	roles, err := dal.RoleTreeSelect()
	return roles, err
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

//根据id得到角色信息
func RoleInfo(id bson.ObjectId) (*modules.Role, error) {
	return dal.RoleInfo(id)
}

//====================================================分割线===========================================
func DepInfo(id bson.ObjectId) (*modules.Department, error) {
	return dal.DepInfo(id)
}
func DepTree(flag int) (*[]modules.Department, error) {
	var where bson.M
	if flag == 0 {
		where = bson.M{"isuse": false}
	} else if flag == 1 {
		where = bson.M{"isuse": true}
	}
	return dal.DepTree(where)
}

//插入一条部门信息
func DepInsert(dep *modules.Department) error {
	dep.ID = bson.NewObjectId()
	return dal.DepInsert(dep)
}

//修改部门信息
func DepEdit(dep *modules.Department) error {
	return dal.DepEdit(dep)
}

//删除部门信息
func DepDelete(id bson.ObjectId) error {
	return dal.DepDelete(id)
}

//==============================================分割线===========================================
func UserInfo(id bson.ObjectId) (*modules.User, error) {
	return dal.UserInfo(id)
}
func UserInsert(user *modules.User) error {
	return dal.UserInsert(user)
}
func UserUpdate(user *modules.User) error {
	return dal.UserUpdate(user)
}
func UserSelect(key, depId string, page, size int) (*[]modules.User, error) {
	where := bson.M{}
	if key = strings.TrimSpace(key); key != "" {
		where["$or"] = []bson.M{bson.M{"LoginName": key}, bson.M{"JobNumber": key}, bson.M{"UserName": key},
			bson.M{"EngName": key}, bson.M{"PinYin": key}, bson.M{"Abbreviation": key}}
	}
	if depId = strings.TrimSpace(depId); depId != "" {
		where["depid"] = depId
	}
	return dal.UserSelect(where, page, size)
}
