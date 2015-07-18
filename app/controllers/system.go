package controllers

import "github.com/revel/revel"

type System struct {
	*revel.Controller
}

func (s System) Role() revel.Result {
	return s.Render()
}
func (s System) GetRoleTreeJson() revel.Result {
	return s.RenderText(`[
	{id: "base", text: "系统管理员"},
	{id: "1", text: "服务经理"},
	{id: "2", text: "服务产品部管理员"},
	{id: "3", text: "办事处主任"},
	{id: "ajax", text: "一级管理员", pid: "base"},
	{id: "json", text: "二级管理员", pid: "base"}
]`)
}
