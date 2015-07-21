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
	return returnMessage(r.Controller, role, err)
}
func (s System) GetRoleTreeJson() revel.Result {
	roles, err := bll.RoleCreateTree()
	return returnMessage(s.Controller, roles, err)
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
