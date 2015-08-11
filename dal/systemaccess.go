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
	query := CloneDB().C(RoleColl).Find(nil).Select(bson.M{"_id": 1, "rolename": 1, "parentid": 1, "rolecode": 1})
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
	info := bson.M{}
	info["rolename"] = role.RoleName
	info["isuse"] = role.IsUse
	info["remark"] = role.Remark
	info["sort"] = role.Sort
	err := CloneDB().C(RoleColl).Update(bson.M{"_id": role.RoleID}, bson.M{"$set": info})
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
	info := bson.M{}
	info["name"] = dep.Name
	info["manager"] = dep.Manager
	info["isuse"] = dep.IsUse
	info["remark"] = dep.Remark
	err := CloneDB().C(DepColl).Update(bson.M{"_id": dep.ID}, bson.M{"$set": info})
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
	info := bson.M{}
	info["loginname"] = user.LoginName
	info["username"] = user.UserName
	info["depid"] = user.DepID
	info["roleids"] = user.RoleIDs
	info["engname"] = user.EngName
	info["pinyin"] = user.PinYin
	info["abbreviation"] = user.Abbreviation
	info["telephone"] = user.Telephone
	info["mobilephone"] = user.Mobilephone
	info["email"] = user.Email
	info["isuse"] = user.IsUse
	return CloneDB().C(UserColl).Update(bson.M{"_id": user.UserID}, bson.M{"$set": info})
}
func UserSelect(where bson.M, page, size int) (*[]modules.User, int, error) {
	users := &[]modules.User{}
	query := CloneDB().C(UserColl).Find(where)
	count, err := query.Count()
	if err != nil {
		return users, 0, err
	}
	if size > 0 {
		query = query.Skip(page * size).Limit(size)
	}
	err = query.All(users)
	return users, count, err
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
	count, err := CloneDB().C(ModuleColl).Count()
	if err != nil {
		return err
	}
	if count == 0 {
		module := &modules.Module{}
		module.ID = bson.NewObjectId()
		module.Name = "系统模块"
		module.Code = "system"
		module.IsUse = true
		module.DisplayType = 1
		module.Remark = "系统默认添加"
		err = CloneDB().C(ModuleColl).Insert(module)
	}
	return err
}
func ModuleInsert(module *modules.Module) error {
	coll := CloneDB().C(ModuleColl)
	if count, err := coll.Find(bson.M{"code": module.Code}).Count(); err == nil {
		if count == 0 {
			return coll.Insert(module)
		} else {
			return fmt.Errorf("%s已经存在该编码！", module.Code)
		}
	} else {
		return err
	}
	return CloneDB().C(ModuleColl).Insert(module)
}
func ModuleUpdate(module *modules.Module) error {
	info := bson.M{}
	info["name"] = module.Name
	info["url"] = module.URL
	info["isuse"] = module.IsUse
	info["sort"] = module.Sort
	info["remark"] = module.Remark
	info["displaytype"] = module.DisplayType
	return CloneDB().C(ModuleColl).Update(bson.M{"_id": module.ID}, bson.M{"$set": info})
}
func ModuleDelete(id bson.ObjectId) error {
	if err := RolePerDeleteByModule(id.String()); err != nil {
		return err
	}
	if err := CloneDB().C(FuncColl).Remove(bson.M{"moduleid": id.String()}); err != nil {
		return err
	}
	return CloneDB().C(ModuleColl).Remove(bson.M{"_id": id})

}
func ModuleInfo(id bson.ObjectId) (*modules.Module, error) {
	module := &modules.Module{}
	err := CloneDB().C(ModuleColl).Find(bson.M{"_id": id}).One(module)
	return module, err
}
func ModuleTree(where bson.M) (*[]modules.Module, error) {
	list := &[]modules.Module{}
	err := CloneDB().C(ModuleColl).Find(where).Select(bson.M{"_id": 1, "parentid": 1, "name": 1, "code": 1}).All(list)
	return list, err
}

//=====================================================================
func FuncInsert(function *modules.Function) error {
	coll := CloneDB().C(FuncColl)
	if count, err := coll.Find(bson.M{"modulecode": function.ModuleCode, "code": function.Code}).Count(); err == nil {
		if count > 0 {
			return fmt.Errorf("%s已经存在该编码！", function.Code)
		} else {
			err = coll.Insert(function)
			return err
		}
	} else {
		return err
	}
}
func FuncUpdate(function *modules.Function) error {
	info := bson.M{}
	info["name"] = function.Name
	info["isuse"] = function.IsUse
	info["remark"] = function.Remark
	return CloneDB().C(FuncColl).Update(bson.M{"_id": function.ID}, bson.M{"$set": info})
}
func FuncDelete(id bson.ObjectId) error {
	if err := RolePerDeleteByFunc(id.String()); err != nil {
		return err
	}
	return CloneDB().C(FuncColl).Remove(bson.M{"_id": id})
}
func FuncInfo(id bson.ObjectId) (*modules.Function, error) {
	function := &modules.Function{}
	err := CloneDB().C(FuncColl).Find(bson.M{"_id": id}).One(function)
	return function, err
}
func FuncSelect(moduleID string) (*[]modules.Function, error) {
	list := &[]modules.Function{}
	err := CloneDB().C(FuncColl).Find(bson.M{"moduleid": moduleID}).All(list)
	return list, err
}

//=========================================================================
func RolePerInfo(id bson.ObjectId) (*modules.RolePermission, error) {
	per := &modules.RolePermission{}
	err := CloneDB().C(RolePermiColl).Find(bson.M{"_id": id}).One(per)
	return per, err
}
func RolePersByParentID(parentPerID string, isModule bool) (*[]modules.RolePermission, error) {
	list := &[]modules.RolePermission{}
	err := CloneDB().C(RolePermiColl).Find(bson.M{"parentperid": parentPerID, "ismodule": isModule}).Select(bson.M{"_id": 1, ""}).All(list)
	return list, err
}
func RolePerCheck(moduleID, roleID string, isModule bool) (*[]modules.RolePermission, error) {
	pers := &[]modules.RolePermission{}
	err := CloneDB().C(RolePermiColl).Find(bson.M{"moduleid": moduleID, "roleid": roleID, "ismodule": isModule}).All(pers)
	return pers, err
}
func RolePerInsert(rp *modules.RolePermission) error {
	var where bson.M
	if rp.IsModule {
		where = bson.M{"moduleid": rp.FunctionID, "roleid": rp.RoleID, "ismodule": true}
	} else {
		where = bson.M{"functionid": rp.FunctionID, "roleid": rp.RoleID, "ismodule": false}
	}
	count, err := CloneDB().C(RolePermiColl).Find(where).Count()
	if err != nil {
		return err
	}
	if count == 0 {
		err = CloneDB().C(RolePermiColl).Insert(rp)
	} else {
		err = CloneDB().C(RolePermiColl).Update(where, bson.M{"#set": bson.M{"isref": false}})
	}
	return err
}
func RolePersInsert(rps *[]modules.RolePermission) error {

}
func RolePerModuleAdd(*[]modules.RolePermission) error {

}
func RolePerDelete(id bson.ObjectId) error {
	err := CloneDB().C(RolePermiColl).Remove(bson.M{"_id": id})
	return err
}
func RolePerDeleteByModule(moduleID string) error {
	err := CloneDB().C(RolePermiColl).Remove(bson.M{"moduleid": moduleID})
	return err
}
func RolePerDeleteByFunc(functionID string) error {
	err := CloneDB().C(RolePermiColl).Remove(bson.M{"functionid": functionID})
	return err
}
func RoleModuleSelect(roleID string) (*[]modules.RolePermission, error) {
	list := &[]modules.RolePermission{}
	err := CloneDB().C(RolePermiColl).Find(bson.M{"roleid": roleID, "ismodule": true}).All(list)
	return list, err
}
func RoleFuncSelect(roleID, moduleID string) (*[]modules.RolePermission, error) {
	list := &[]modules.RolePermission{}
	err := CloneDB().C(RolePermiColl).Find(bson.M{"roleid": roleID, "moduleID": moduleID, "ismodule": false}).All(list)
	return list, err
}
