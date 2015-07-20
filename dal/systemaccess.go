package dal

import (
	"github.com/qtzheng/SIMP/app/modules"
	"gopkg.in/mgo.v2/bson"
)

//初始化角色集合，增加默认角色
func roleInit() error {
	count, err := NewDB().C(RoleColl).Count()
	if err != nil {
		return err
	}
	if count == 0 {
		role := &modules.Role{}
		role.RoleID = bson.NewObjectId()
		role.IsUse = true
		role.Remark = "系统默认根节点，不可删除，不可编辑"
		role.RoleCode = "Defult"
		role.RoleName = "系统角色"
		err = NewDB().C(RoleColl).Insert(role)
		if err != nil {
			return err
		}
	}
	return nil
}

//角色总数./revel run github.com/qtzheng/SIMP
func RoleCount() (int, error) {
	count, err := NewDB().C(RoleColl).Count()
	if err != nil {
		return 0, err
	}
	return count, nil
}

//查询角色列表
func RoleSelect(page, size int) (*[]modules.Role, error) {
	roles := &[]modules.Role{}
	query := NewDB().C(RoleColl).Find(bson.M{}).Skip(page * size).Limit(size)
	err := query.All(roles)
	return roles, err
}

//查询角色树
func RoleTreeSelect() (*[]modules.Role, error) {
	roles := &[]modules.Role{}
	query := NewDB().C(RoleColl).Find(nil).Select(bson.M{"_id": 1, "rolename": 1, "parentid": 1})
	err := query.All(roles)
	return roles, err
}

//插入一条角色信息
func RoleInsert(role *modules.Role) error {
	err := NewDB().C(RoleColl).Insert(role)
	return err
}

//修改权限信息
func RoleEdit(role *modules.Role) error {
	err := NewDB().C(RoleColl).Update(bson.M{"_id": role.RoleID}, role)
	return err
}

//删除权限信息
func RoleDelete(id bson.ObjectId) error {
	err := NewDB().C(RoleColl).Remove(bson.M{"_id": id})
	return err
}

//根据角色编码查询信息
func RoleInfoByCode(code string) (*modules.Role, error) {
	role := &modules.Role{}
	err := NewDB().C(RoleColl).Find(bson.M{"RoleCode": code}).One(role)
	if err != nil {
		return nil, err
	} else {
		return role, nil
	}
}

//根据角色id查询信息
func RoleInfo(id bson.ObjectId) (*modules.Role, error) {
	role := &modules.Role{}
	err := NewDB().C(RoleColl).Find(bson.M{"_id": id}).One(role)
	if err != nil {
		return nil, err
	} else {
		return role, nil
	}
}

/*
======================================================分割线========================================================
*/
