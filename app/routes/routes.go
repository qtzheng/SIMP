// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tBaseCollection struct {}
var BaseCollection tBaseCollection



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

func (_ tSystem) UserInsert(
		user interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	return revel.MainRouter.Reverse("System.UserInsert", args).Url
}

func (_ tSystem) UserUpdate(
		user interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "user", user)
	return revel.MainRouter.Reverse("System.UserUpdate", args).Url
}

func (_ tSystem) UserInfo(
		id interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("System.UserInfo", args).Url
}

func (_ tSystem) Module(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("System.Module", args).Url
}

func (_ tSystem) ModuleInsert(
		module interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "module", module)
	return revel.MainRouter.Reverse("System.ModuleInsert", args).Url
}

func (_ tSystem) ModuleUpdate(
		module interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "module", module)
	return revel.MainRouter.Reverse("System.ModuleUpdate", args).Url
}

func (_ tSystem) ModuleInfo(
		id interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("System.ModuleInfo", args).Url
}

func (_ tSystem) ModuleDelete(
		id interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("System.ModuleDelete", args).Url
}

func (_ tSystem) ModelTreeUsed(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("System.ModelTreeUsed", args).Url
}

func (_ tSystem) ModelTree(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("System.ModelTree", args).Url
}

func (_ tSystem) FuncInsert(
		function interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "function", function)
	return revel.MainRouter.Reverse("System.FuncInsert", args).Url
}

func (_ tSystem) FuncUpdate(
		function interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "function", function)
	return revel.MainRouter.Reverse("System.FuncUpdate", args).Url
}

func (_ tSystem) FuncDelete(
		id interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("System.FuncDelete", args).Url
}

func (_ tSystem) FuncInfo(
		id interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("System.FuncInfo", args).Url
}

func (_ tSystem) FuncSelect(
		moduleID interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleID", moduleID)
	return revel.MainRouter.Reverse("System.FuncSelect", args).Url
}

func (_ tSystem) RolePerAdd(
		per interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "per", per)
	return revel.MainRouter.Reverse("System.RolePerAdd", args).Url
}

func (_ tSystem) RolePerDelete(
		id interface{},
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("System.RolePerDelete", args).Url
}

func (_ tSystem) RoleModule(
		roleID string,
		moduleID string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "roleID", roleID)
	revel.Unbind(args, "moduleID", moduleID)
	return revel.MainRouter.Reverse("System.RoleModule", args).Url
}

func (_ tSystem) RoleFunc(
		roleID string,
		moduleID string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "roleID", roleID)
	revel.Unbind(args, "moduleID", moduleID)
	return revel.MainRouter.Reverse("System.RoleFunc", args).Url
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


