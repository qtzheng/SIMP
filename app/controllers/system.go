package controllers

import (
	"github.com/qtzheng/SIMP/bll"
	"github.com/qtzheng/SIMP/modules"
	"github.com/revel/revel"
	"gopkg.in/mgo.v2/bson"
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
func (s System) UserSelect(key string, depIds *[]bson.ObjectId) revel.Result {
	users, err := bll.UserSelect(depIds, key)
	if err != nil {
		return s.RenderText("")
	} else {
		return s.RenderJson(users)
	}
}
func (s System) UserInsert() {

}
