package controllers

import (
	"github.com/qtzheng/SIMP/app/modules"
	"github.com/qtzheng/SIMP/bll"
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
	if err != nil {
		revel.WARN.Fatal(err)
		return r.RenderJson(&opResult{Error, err})
	}
	result := &opResult{Success, role}
	return r.RenderJson(result)
}
func (s System) GetRoleTreeJson() revel.Result {
	roles := bll.RoleCreateTree()
	return s.RenderJson(roles)
}
func (r System) AddRole(role *modules.Role) revel.Result {
	err := bll.RoleInsert(role)
	if err != nil {
		revel.WARN.Fatal(err)
		return r.RenderJson(&opResult{Error, err})
	}
	result := &opResult{Success, role.RoleID}
	return r.RenderJson(result)
}
func (s System) EditRole(role *modules.Role) revel.Result {
	err := bll.RoleEdit(role)
	if err != nil {
		revel.WARN.Fatal(err)
		return s.RenderJson(&opResult{Error, err})
	}
	result := &opResult{Success, ""}
	return s.RenderJson(result)
}
func (s System) RoleDelete(id bson.ObjectId) revel.Result {
	err := bll.RoleDelete(id)
	if err != nil {
		revel.WARN.Fatal(err)
		return s.RenderJson(&opResult{Error, err})
	}

	result := &opResult{Success, ""}
	return s.RenderJson(result)
}
