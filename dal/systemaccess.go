package dal

import (
	"fmt"
	"github.com/qtzheng/SIMP/modules"
	"gopkg.in/mgo.v2/bson"
)

//初始化角色集合，增加默认角色
func roleInit() error {
	count, err := CloneDB().C(RoleColl).Count()
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
		err = CloneDB().C(RoleColl).Insert(role)
		if err != nil {
			return err
		}
	}
	return nil
}

//角色总数./revel run github.com/qtDepartmentzheng/SIMP
func RoleCount() (int, error) {
	count, err := CloneDB().C(RoleColl).Count()
	if err != nil {
		return 0, err
	}
	return count, nil
}

//查询角色列表
func RoleSelect(page, size int) (*[]modules.Role, error) {
	roles := &[]modules.Role{}
	query := CloneDB().C(RoleColl).Find(bson.M{}).Skip(page * size).Limit(size)
	err := query.All(roles)
	return roles, err
}

//查询角色树
func RoleTreeSelect() (*[]modules.Role, error) {
	roles := &[]modules.Role{}
	query := CloneDB().C(RoleColl).Find(nil).Select(bson.M{"_id": 1, "rolename": 1, "parentid": 1})
	err := query.All(roles)
	return roles, err
}

//插入一条角色信息
func RoleInsert(role *modules.Role) error {
	coll := CloneDB().C(RoleColl)
	count, err := coll.Find(bson.M{"rolecode": role.RoleCode}).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("编码：%s已经存在！", role.RoleCode)
	}

	err = coll.Insert(role)
	return err
}

//修改权限信息
func RoleEdit(role *modules.Role) error {
	err := CloneDB().C(RoleColl).Update(bson.M{"_id": role.RoleID}, role)
	return err
}

//删除权限信息
func RoleDelete(id bson.ObjectId) error {
	err := CloneDB().C(RoleColl).Remove(bson.M{"_id": id})
	return err
}

//根据角色编码查询信息
func RoleInfoByCode(code string) (*modules.Role, error) {
	role := &modules.Role{}
	err := CloneDB().C(RoleColl).Find(bson.M{"RoleCode": code}).One(role)
	if err != nil {
		return nil, err
	} else {
		return role, nil
	}
}

//根据角色id查询信息
func RoleInfo(id bson.ObjectId) (*modules.Role, error) {
	role := &modules.Role{}
	err := CloneDB().C(RoleColl).Find(bson.M{"_id": id}).One(role)
	if err != nil {
		return nil, err
	} else {
		return role, nil
	}
}

/*
======================================================分割线========================================================
*/
func depInit() error {
	count, err := CloneDB().C(DepColl).Count()
	if err != nil {
		return err
	}
	if count == 0 {
		dep := &modules.Department{}
		dep.ID = bson.NewObjectId()
		dep.IsUse = true
		dep.Remark = "系统默认根节点，不可删除，不可编辑"
		dep.Code = "Defult"
		dep.Name = "组织架构"
		err = CloneDB().C(DepColl).Insert(dep)
		if err != nil {
			return err
		}
	}
	return nil
}

//根据部门id查询信息
func DepInfo(id bson.ObjectId) (*modules.Department, error) {
	dep := &modules.Department{}
	err := CloneDB().C(DepColl).Find(bson.M{"_id": id}).One(dep)
	if err != nil {
		return nil, err
	} else {
		return dep, nil
	}
}
func DepTree(where bson.M) (*[]modules.Department, error) {
	dep := &[]modules.Department{}
	query := CloneDB().C(DepColl).Find(where).Select(bson.M{"_id": 1, "name": 1, "parentid": 1})
	err := query.All(dep)
	return dep, err
}

//插入一条部门信息
func DepInsert(dep *modules.Department) error {
	coll := CloneDB().C(DepColl)
	count, err := coll.Find(bson.M{"code": dep.Code}).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("编码：%s已经存在！", dep.Code)
	}
	err = CloneDB().C(DepColl).Insert(dep)
	return err
}

//修改部门信息
func DepEdit(dep *modules.Department) error {
	err := CloneDB().C(DepColl).Update(bson.M{"_id": dep.ID}, dep)
	return err
}

//删除部门信息
func DepDelete(id bson.ObjectId) error {
	err := CloneDB().C(DepColl).Remove(bson.M{"_id": id})
	return err
}

//==============================================分割线===========================================
func userInit() error {
	count, err := CloneDB().C(UserColl).Count()
	if err != nil {
		return err
	}
	if count == 0 {
		user := &modules.User{}
		user.Abbreviation = "admin"
		user.DepID = ""
		user.Email = ""
		user.EngName = "admin"
		user.IsUse = true
		user.LoginName = "admin"
		user.UserID = bson.NewObjectId()
		user.Password = "admin"
		err = CloneDB().C(UserColl).Insert(user)
	}
	return err
}
func UserInfo(id bson.ObjectId) (*modules.User, error) {
	user := &modules.User{}
	err := CloneDB().C(UserColl).Find(bson.M{"_id": id}).One(user)
	if err != nil {
		return nil, err
	} else {
		return user, nil
	}
}
func UserInsert(user *modules.User) error {
	return CloneDB().C(UserColl).Insert(user)
}
func UserUpdate(user *modules.User) error {
	return CloneDB().C(UserColl).Update(bson.M{"_id": user.UserID}, user)
}
func UserSelect(where bson.M, page, size int) (*[]modules.User, error) {
	users := &[]modules.User{}
	query := CloneDB().C(UserColl).Find(where)
	if size > 0 {
		query = query.Skip(page * size).Limit(size)
	}
	err := query.All(users)
	return users, err
}
func UserRoles(ids []bson.ObjectId) (*[]modules.Role, error) {
	roles := &[]modules.Role{}
	err := CloneDB().C(RoleColl).Find(bson.M{"_id": bson.M{"$in": ids}}).All(roles)
	if err != nil {
		return nil, err
	} else {
		return roles, nil
	}
}

//=============================================================================================
func moduleInit() error {
	count, err = CloneDB().C(ModuleColl).Count()
	if err != nil {
		return err
	}
	if coun == 0 {
		module := &modules.Module{}
		module.ID = bson.NewObjectId()
		module.Name = "系统模块"
		err = CloneDB().C(ModuleColl).Insert(module)
	}
}
