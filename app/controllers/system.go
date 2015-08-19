package controllers

import (
	"github.com/qtzheng/SIMP/bll"
	"github.com/qtzheng/SIMP/modules"
	"github.com/revel/revel"
	"gopkg.in/mgo.v2/bson"
	"strings"
)

type System struct {
	BaseCollection
}

func (s System) Role() revel.Result {
	return s.Render()
}
func (r System) GetRoleInfo(id bson.ObjectId) revel.Result {
	role, err := bll.RoleInfo(id)
	return returnMessage(r.Controller, role, err)
}
func (s System) GetRoleTreeJson() revel.Result {
	roles, err := bll.RoleCreateTree()
	if err != nil {
		return s.RenderText("")
	} else {
		return s.RenderJson(roles)
	}
}
func (r System) AddRole(role *modules.Role) revel.Result {
	err := bll.RoleInsert(role)
	return returnMessage(r.Controller, role.RoleID, err)
}
func (s System) EditRole(role *modules.Role) revel.Result {
	err := bll.RoleEdit(role)
	return returnMessage(s.Controller, "", err)
}
func (s System) RoleDelete(id bson.ObjectId) revel.Result {
	err := bll.RoleDelete(id)
	return returnMessage(s.Controller, "", err)
}

/*==========================================分割线===========================*/
func (s System) Dep() revel.Result {
	return s.Render()
}
func (s System) DepTree() revel.Result {
	roles, err := bll.DepTree(2)
	if err != nil {
		return s.RenderText("")
	} else {
		return s.RenderJson(roles)
	}
}
func (s System) DepTreeUsed() revel.Result {
	roles, err := bll.DepTree(1)
	if err != nil {
		return s.RenderText("")
	} else {
		return s.RenderJson(roles)
	}
}
func (s System) DepInfo(id bson.ObjectId) revel.Result {
	dep, err := bll.DepInfo(id)
	return returnMessage(s.Controller, dep, err)
}
func (s System) DepInsert(dep *modules.Department) revel.Result {
	err := bll.DepInsert(dep)
	return returnMessage(s.Controller, dep.ID, err)
}
func (s System) DepEdit(dep *modules.Department) revel.Result {
	err := bll.DepEdit(dep)
	return returnMessage(s.Controller, "", err)
}
func (s System) DepDelete(id bson.ObjectId) revel.Result {
	err := bll.DepDelete(id)
	return returnMessage(s.Controller, "", err)
}

//===========================================================================
func (s System) UserSelect(key string, depIds string, pageIndex, pageSize int) revel.Result {
	users, count, err := bll.UserSelect(key, depIds, pageIndex, pageSize)
	if err != nil {
		return s.RenderText("")
	} else {
		return GetGridJson(s.Controller, count, users)
	}
}
func (s System) UserInsert(user *modules.User) revel.Result {
	user.RoleIDs = strings.Split(s.Params.Get("RoleIDs"), ",")
	err := bll.UserInsert(user)
	return returnMessage(s.Controller, user.UserID, err)
}
func (s System) UserUpdate(user *modules.User) revel.Result {
	user.RoleIDs = strings.Split(s.Params.Get("RoleIDs"), ",")
	err := bll.UserUpdate(user)
	return returnMessage(s.Controller, "", err)
}
func (s System) UserInfo(id bson.ObjectId) revel.Result {
	user, err := bll.UserInfo(id)
	return returnMessage(s.Controller, user, err)
}

//==============================================================================
func (s System) Module() revel.Result {
	return s.Render()
}
func (s System) ModuleInsert(module *modules.Module) revel.Result {
	err := bll.ModuleInsert(module)
	return returnMessage(s.Controller, module.ID, err)
}
func (s System) ModuleUpdate(module *modules.Module) revel.Result {
	err := bll.ModuleUpdate(module)
	return returnMessage(s.Controller, "", err)
}
func (s System) ModuleInfo(id bson.ObjectId) revel.Result {
	module, err := bll.ModuleInfo(id)
	return returnMessage(s.Controller, module, err)
}
func (s System) ModuleDelete(id bson.ObjectId) revel.Result {
	err := bll.ModuleDelete(id)
	return returnMessage(s.Controller, "", err)
}
func (s System) ModelTreeUsed() revel.Result {
	list, err := bll.ModuleTree(1)
	if err == nil {
		return s.RenderJson(list)
	} else {
		return s.RenderText("")
	}
}
func (s System) ModelTree() revel.Result {
	list, err := bll.ModuleTree(2)
	if err == nil {
		return s.RenderJson(list)
	} else {
		return s.RenderText("")
	}
}

//========================================================================================
func (s System) FuncInsert(function *modules.Function) revel.Result {
	err := bll.FuncInsert(function)
	return returnMessage(s.Controller, function.ID, err)
}
func (s System) FuncUpdate(function *modules.Function) revel.Result {
	err := bll.FuncUpdate(function)
	return returnMessage(s.Controller, "", err)
}
func (s System) FuncDelete(id bson.ObjectId) revel.Result {
	err := bll.FuncDelete(id)
	return returnMessage(s.Controller, "", err)
}
func (s System) FuncInfo(id bson.ObjectId) revel.Result {
	function, err := bll.FuncInfo(id)
	return returnMessage(s.Controller, function, err)
}
func (s System) FuncSelect(moduleid string) revel.Result {
	list, err := bll.FuncSelect(moduleid)
	if err == nil {
		return s.RenderJson(list)
	} else {
		revel.WARN.Fatalln(err)
		return s.RenderText("")
	}
}

//========================================================================================
func (s System) RolePerModuleAdd(RoleID, RoleCode, ModuleID string) revel.Result {
	data, err := bll.RolePerModuleAdd(RoleID, RoleCode, ModuleID)
	return returnMessage(s.Controller, data, err)
}
func (s System) RolePerModuleDelete(id bson.ObjectId) revel.Result {
	data, err := bll.RolePerModuleDelete(id)
	return returnMessage(s.Controller, data, err)
}
func (s System) RolePerFuncDelete(id bson.ObjectId) revel.Result {
	err := bll.RolePerFuncDelete(id)
	return returnMessage(s.Controller, "", err)
}
func (s System) RoleModule(roleID string) revel.Result {
	data, err := bll.RolePerModule(roleID)
	return returnMessage(s.Controller, data, err)
}
func (s System) RoleFunc(roleID, moduleID string) revel.Result {
	data, err := bll.RolePerFunc(roleID, moduleID)
	return returnMessage(s.Controller, data, err)
}
