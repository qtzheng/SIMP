// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tBaseCollection struct {}
var BaseCollection tBaseCollection



type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).Url
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).Url
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).Url
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).Url
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).Url
}


type tSystem struct {}
var System tSystem


func (_ tSystem) Role(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("System.Role", args).Url
}

func (_ tSystem) GetRoleInfo(
		id interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("System.GetRoleInfo", args).Url
}

func (_ tSystem) GetRoleTreeJson(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("System.GetRoleTreeJson", args).Url
}

func (_ tSystem) AddRole(
		role interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "role", role)
	return revel.MainRouter.Reverse("System.AddRole", args).Url
}

func (_ tSystem) EditRole(
		role interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "role", role)
	return revel.MainRouter.Reverse("System.EditRole", args).Url
}

func (_ tSystem) RoleDelete(
		id interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("System.RoleDelete", args).Url
}

func (_ tSystem) Dep(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("System.Dep", args).Url
}

func (_ tSystem) DepTree(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("System.DepTree", args).Url
}

func (_ tSystem) DepTreeUsed(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("System.DepTreeUsed", args).Url
}

func (_ tSystem) DepInfo(
		id interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("System.DepInfo", args).Url
}

func (_ tSystem) DepInsert(
		dep interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "dep", dep)
	return revel.MainRouter.Reverse("System.DepInsert", args).Url
}

func (_ tSystem) DepEdit(
		dep interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "dep", dep)
	return revel.MainRouter.Reverse("System.DepEdit", args).Url
}

func (_ tSystem) DepDelete(
		id interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("System.DepDelete", args).Url
}

func (_ tSystem) UserSelect(
		key string,
		depIds string,
		pageIndex int,
		pageSize int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "key", key)
	revel.Unbind(args, "depIds", depIds)
	revel.Unbind(args, "pageIndex", pageIndex)
	revel.Unbind(args, "pageSize", pageSize)
	return revel.MainRouter.Reverse("System.UserSelect", args).Url
}


type tApp struct {}
var App tApp


func (_ tApp) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Index", args).Url
}

func (_ tApp) Role(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Role", args).Url
}


