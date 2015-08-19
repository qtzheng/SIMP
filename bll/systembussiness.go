package bll

import (
	"fmt"
	"github.com/qtzheng/SIMP/dal"
	"github.com/qtzheng/SIMP/modules"
	"gopkg.in/mgo.v2/bson"
	"strings"
)

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
	user.UserID = bson.NewObjectId()
	return dal.UserInsert(user)
}
func UserUpdate(user *modules.User) error {
	return dal.UserUpdate(user)
}
func UserSelect(key string, depIds string, page, size int) (*[]modules.User, int, error) {
	where := bson.M{}
	if key = strings.TrimSpace(key); key != "" {
		where["$or"] = []bson.M{bson.M{"LoginName": key}, bson.M{"JobNumber": key}, bson.M{"UserName": key},
			bson.M{"EngName": key}, bson.M{"PinYin": key}, bson.M{"Abbreviation": key}}
	}
	ids := strings.Split(depIds, ",")
	if len(ids) > 0 && strings.TrimSpace(ids[0]) != "" {
		where["depid"] = bson.M{"$in": ids}
	}
	return dal.UserSelect(where, page, size)
}

//===============================================================================================
func ModuleInsert(module *modules.Module) error {
	module.ID = bson.NewObjectId()
	return dal.ModuleInsert(module)
}
func ModuleUpdate(module *modules.Module) error {
	return dal.ModuleUpdate(module)
}
func ModuleDelete(id bson.ObjectId) error {
	return dal.ModuleDelete(id)
}
func ModuleInfo(id bson.ObjectId) (*modules.Module, error) {
	return dal.ModuleInfo(id)
}
func ModuleTree(flag int) (*[]modules.Module, error) {
	var where bson.M
	if flag == 0 {
		where = bson.M{"isuse": false}
	} else if flag == 1 {
		where = bson.M{"isuse": true}
	}
	return dal.ModuleTree(where)
}

//=====================================================================
func FuncInsert(function *modules.Function) error {
	function.ID = bson.NewObjectId()
	return dal.FuncInsert(function)
}
func FuncUpdate(function *modules.Function) error {
	return dal.FuncUpdate(function)
}
func FuncDelete(id bson.ObjectId) error {
	return dal.FuncDelete(id)
}
func FuncInfo(id bson.ObjectId) (*modules.Function, error) {
	return dal.FuncInfo(id)
}
func FuncSelect(moduleCode string) (*[]modules.Function, error) {
	return dal.FuncSelect(moduleCode)
}

//=====================================================================
func RolePerModuleAdd(roleID, roleCode, moduleID string) ([]modules.RolePermission, error) {
	list := []modules.RolePermission{}
	per, err := dal.RolePerCheck(moduleID, roleID, true)
	if err != nil {
		return nil, err
	}
	if per != nil && per.IsRef {
		err = dal.RolePerIsRefUpdate(per.PermissionId, false)
		if err != nil {
			return nil, err
		}
	} else if per == nil {
		mod, modErr := dal.ModuleInfo(bson.ObjectIdHex(moduleID))
		if modErr != nil {
			return nil, modErr
		}
		per = &modules.RolePermission{}
		per.PermissionId = bson.NewObjectId()
		per.RoleID = roleID
		per.RoleCode = roleCode
		per.ModuleID = mod.ID.Hex()
		per.ModuleCode = mod.Code
		per.IsModule = true
		per.IsRef = false
		parentMod, parErr := dal.ModuleInfo(bson.ObjectIdHex(mod.ParentID))
		if parErr != nil {
			return nil, parErr
		} else if parentMod.Code != "system" {
			parentPer, perErr := dal.RolePerCheck(parentMod.ID.Hex(), roleID, true)
			if perErr != nil {
				return nil, perErr
			}
			if parentPer == nil {
				parentPer.PermissionId = bson.NewObjectId()
				parentPer.RoleID = roleID
				parentPer.RoleCode = roleCode
				parentPer.ModuleID = parentMod.ID.Hex()
				parentPer.ModuleCode = parentMod.Code
				parentPer.IsModule = true
				parentPer.IsRef = true
				per.ParentPerID = parentPer.PermissionId.Hex()
			} else {
				per.ParentPerID = parentPer.PermissionId.Hex()
			}
			err = rolePerAddCheck(roleID, roleCode, parentPer, list)
		}
	}
	list = append(list, *per)
	if err == nil {
		err = dal.RolePersInsert(list)
	}
	return list, err
}
func rolePerAddCheck(roleID, roleCode string, per *modules.RolePermission, parentPers []modules.RolePermission) error {

	parentMod, parErr := dal.ModuleInfo(bson.ObjectIdHex(per.ModuleID))
	if parErr != nil {
		return parErr
	} else if parentMod.Code == "system" {
		return nil
	}
	parentPer, perErr := dal.RolePerCheck(parentMod.ID.Hex(), roleID, true)
	if perErr != nil {
		return perErr
	}
	if parentPer == nil {
		parentPer = &modules.RolePermission{}
		parentPer.PermissionId = bson.NewObjectId()
		parentPer.RoleID = roleID
		parentPer.RoleCode = roleCode
		parentPer.ModuleID = parentMod.ID.Hex()
		parentPer.ModuleCode = parentMod.Code
		parentPer.IsModule = true
		parentPer.IsRef = true
		per.ParentPerID = parentPer.PermissionId.Hex()
	} else {
		per.ParentPerID = parentPer.PermissionId.Hex()
	}
	parentPers = append(parentPers, *per)
	return rolePerAddCheck(roleID, roleCode, parentPer, parentPers)
}
func RolePerModuleDelete(id bson.ObjectId) ([]string, error) {
	perIDs := make([]string, 20)
	per, err := dal.RolePerInfo(id)
	if err != nil {
		return nil, err
	}
	count, err := dal.RolePersCountByParentID(id.Hex(), true)
	if err != nil {
		return nil, err
	}
	if count > 0 {
		err = dal.RolePerIsRefUpdate(id, true)
	} else {
		err = dal.RolePerDelete(id)
	}
	if err != nil {
		return nil, err
	}
	perIDs = append(perIDs, id.Hex())
	err = rolePerDelCheck(per.ParentPerID, perIDs)
	return perIDs, err
}
func rolePerDelCheck(permissionID string, perIDs []string) error {
	per, err := dal.RolePerInfo(bson.ObjectIdHex(permissionID))
	if err != nil {
		return err
	}
	if !per.IsRef {
		return nil
	}
	count, err := dal.RolePersCountByParentID(permissionID, true)
	if err != nil {
		return err
	}
	if count == 0 {
		err = dal.RolePerDelete(bson.ObjectIdHex(permissionID))
		if err != nil {
			return err
		}
		perIDs = append(perIDs, permissionID)
		return rolePerDelCheck(per.ParentPerID, perIDs)
	}
	return nil

}
func RolePerFuncDelete(id bson.ObjectId) error {
	return dal.RolePerDelete(id)
}
func RolePerModule(roleID string) (*[]modules.RolePermission, error) {
	roleID = strings.TrimSpace(roleID)
	if roleID == "" {
		return nil, fmt.Errorf("角色ID不能为空！")
	}
	per_m, err := dal.RoleModuleSelect(roleID)
	return per_m, err

}
func RolePerFunc(roleID, moduleID string) (*[]modules.RolePermission, error) {
	roleID = strings.TrimSpace(roleID)
	moduleID = strings.TrimSpace(moduleID)
	if roleID == "" {
		return nil, fmt.Errorf("角色ID不能为空！")
	}
	list, err := dal.RoleFuncSelect(roleID, moduleID)
	return list, err
}
