// GENERATED CODE - DO NOT EDIT
package main

import (
	"flag"
	"reflect"
	"github.com/revel/revel"
	_ "github.com/qtzheng/SIMP/app"
	controllers "github.com/qtzheng/SIMP/app/controllers"
	modules "github.com/qtzheng/SIMP/modules"
	tests "github.com/qtzheng/SIMP/tests"
	controllers0 "github.com/revel/modules/static/app/controllers"
	_ "github.com/revel/modules/testrunner/app"
	controllers1 "github.com/revel/modules/testrunner/app/controllers"
	bson "gopkg.in/mgo.v2/bson"
	"github.com/revel/revel/testing"
)

var (
	runMode    *string = flag.String("runMode", "", "Run mode.")
	port       *int    = flag.Int("port", 0, "By default, read from app.conf")
	importPath *string = flag.String("importPath", "", "Go Import Path for the app.")
	srcPath    *string = flag.String("srcPath", "", "Path to the source root.")

	// So compiler won't complain if the generated code doesn't reference reflect package...
	_ = reflect.Invalid
)

func main() {
	flag.Parse()
	revel.Init(*runMode, *importPath, *srcPath)
	revel.INFO.Println("Running revel server")
	
	revel.RegisterController((*controllers.BaseCollection)(nil),
		[]*revel.MethodType{
			
		})
	
	revel.RegisterController((*controllers0.Static)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Serve",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ServeModule",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "moduleName", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers1.TestRunner)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					70: []string{ 
						"testSuites",
					},
				},
			},
			&revel.MethodType{
				Name: "Run",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "suite", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "test", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					107: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "List",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers.System)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Role",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					15: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "GetRoleInfo",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*bson.ObjectId)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "GetRoleTreeJson",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "AddRole",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "role", Type: reflect.TypeOf((**modules.Role)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "EditRole",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "role", Type: reflect.TypeOf((**modules.Role)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "RoleDelete",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*bson.ObjectId)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Dep",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					44: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "DepTree",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "DepTreeUsed",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "DepInfo",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*bson.ObjectId)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "DepInsert",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "dep", Type: reflect.TypeOf((**modules.Department)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "DepEdit",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "dep", Type: reflect.TypeOf((**modules.Department)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "DepDelete",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*bson.ObjectId)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "UserSelect",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "key", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "depIds", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "pageIndex", Type: reflect.TypeOf((*int)(nil)) },
					&revel.MethodArg{Name: "pageSize", Type: reflect.TypeOf((*int)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "UserInsert",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "user", Type: reflect.TypeOf((**modules.User)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "UserUpdate",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "user", Type: reflect.TypeOf((**modules.User)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "UserInfo",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*bson.ObjectId)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ModuleInsert",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "module", Type: reflect.TypeOf((**modules.Module)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ModuleUpdate",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "module", Type: reflect.TypeOf((**modules.Module)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ModuleInfo",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*bson.ObjectId)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ModuleDelete",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*bson.ObjectId)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ModelTreeUsed",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ModelTree",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "FuncInsert",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "function", Type: reflect.TypeOf((**modules.Function)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "FuncUpdate",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "function", Type: reflect.TypeOf((**modules.Function)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "FuncDelete",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*bson.ObjectId)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "FuncInfo",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "id", Type: reflect.TypeOf((*bson.ObjectId)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "FuncSelect",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "moduleID", Type: reflect.TypeOf((*bson.ObjectId)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers.App)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					31: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "Role",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					34: []string{ 
					},
				},
			},
			
		})
	
	revel.DefaultValidationKeys = map[string]map[int]string{ 
	}
	testing.TestSuites = []interface{}{ 
		(*tests.AppTest)(nil),
	}

	revel.Run(*port)
}
