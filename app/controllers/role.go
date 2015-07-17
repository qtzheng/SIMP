package controllers

import (
	//"encoding/json"
	"github.com/qtzheng/SIMP/app/modules"
	//"github.com/qtzheng/SIMP/bll"
	"github.com/revel/revel"
)

type Role struct {
	Application
}

func (c Role) Index() revel.Result {
	return c.Render()
}
func (r Role) GetRoleTreeJson() revel.Result {

	return r.RenderText(`[
	{id: "base", text: "系统管理员"},
	{id: "1", text: "服务经理"},
	{id: "2", text: "服务产品部管理员"},
	{id: "3", text: "办事处主任"},
	{id: "ajax", text: "一级管理员", pid: "base"},
	{id: "json", text: "二级管理员", pid: "base"}
]`)
}
func (r Role) name() revel.Result {
	return r.RenderText(`[
	{id: "base", text: "系统管理模块"},
	{id: "1", text: "服务经理"},
	{id: "2", text: "服务产品部管理员"},
	{id: "3", text: "办事处主任"},
	{id: "ajax", text: "一级管理员", pid: "base"},
	{id: "json", text: "二级管理员", pid: "base"}
]`)
}
func (r Role) GetRoleInfo() revel.Result {
	return r.RenderText(`{"RoleID":"1","RoleName":"测试角色","RoleCode":"Test","Sort":0,"IsUse":2,"Remark":"车上"}`)
}
func (r Role) AddRole(role *modules.Role) revel.Result {
	r.Params.Bind(role, "role")
	//err := bll.RoleInsert(role)
	//if err != nil {
	//	revel.WARN.Fatal(err)
	//	return r.RenderText("保存失败")
	//}
	return r.RenderJson(role)
}
