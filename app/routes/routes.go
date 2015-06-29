// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tApplication struct {}
var Application tApplication



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


type tRole struct {}
var Role tRole


func (_ tRole) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Role.Index", args).Url
}


type tSystem struct {}
var System tSystem


func (_ tSystem) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("System.Index", args).Url
}


type tAccount struct {}
var Account tAccount


func (_ tAccount) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Account.Index", args).Url
}


type tApp struct {}
var App tApp


func (_ tApp) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Index", args).Url
}


type tDepartment struct {}
var Department tDepartment


func (_ tDepartment) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Department.Index", args).Url
}


type tModule struct {}
var Module tModule


func (_ tModule) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Module.Index", args).Url
}


