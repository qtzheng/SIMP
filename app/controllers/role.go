package controllers

import (
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
