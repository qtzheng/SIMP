package controllers

import (
	"github.com/qtzheng/SIMP/app/modules"
	"github.com/qtzheng/SIMP/bll"
	"github.com/revel/revel"
)

type System struct {
	*revel.Controller
}

func (s System) Role() revel.Result {
	return s.Render()
}
func (r System) GetRoleInfo() revel.Result {
	return r.RenderText(`{"RoleID":"1","RoleName":"测试角色","RoleCode":"Test","Sort":0,"IsUse":2,"Remark":"车上"}`)
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
