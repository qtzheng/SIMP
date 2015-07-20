package controllers

import (
	"fmt"
	"github.com/revel/revel"
	"gopkg.in/mgo.v2/bson"
	"reflect"
	"strings"
)

func init() {
	revel.KindBinders[reflect.Struct] = myBinder
	revel.TypeBinders[reflect.TypeOf(bson.NewObjectId())] = revel.Binder{
		Bind: func(params *revel.Params, name string, typ reflect.Type) reflect.Value {
			values, ok := params.Values[name]
			if !ok || len(values) == 0 {
				return reflect.Zero(typ)
			}
			var value bson.ObjectId = bson.ObjectId(values[0])
			return reflect.ValueOf(value)
		},
		Unbind: nil,
	}
}

type App struct {
	BaseCollection
}

func (c App) Index() revel.Result {
	return c.Render()
}
func (a App) Role() revel.Result {
	return a.Render()
}

type opResult struct {
	Result  int
	Message interface{}
}

const (
	Success = 0 + iota
	Fail
	Error
)

var myBinder = revel.Binder{
	bindStruct,
	unbindStruct,
}

func nextKey(key string) string {
	fieldLen := strings.IndexAny(key, ".[")
	if fieldLen == -1 {
		return key
	}
	return key[:fieldLen]
}
func bindStruct(params *revel.Params, name string, typ reflect.Type) reflect.Value {
	result := reflect.New(typ).Elem()
	fieldValues := make(map[string]reflect.Value)
	for key, _ := range params.Values {
		// Get the name of the struct property.
		// Strip off the prefix. e.g. foo.bar.baz => bar.baz
		suffix := strings.TrimSpace(key)
		fieldName := nextKey(suffix)
		//fieldLen := len(fieldName)
		fmt.Println(key)
		if _, ok := fieldValues[fieldName]; !ok {
			// Time to bind this field.  Get it and make sure we can set it.
			fieldValue := result.FieldByName(fieldName)
			if !fieldValue.IsValid() {
				revel.WARN.Println("W: bindStruct: Field not found:", fieldName)
				continue
			}
			if !fieldValue.CanSet() {
				revel.WARN.Println("W: bindStruct: Field not settable:", fieldName)
				continue
			}
			boundVal := revel.Bind(params, fieldName, fieldValue.Type())
			fieldValue.Set(boundVal)
			fieldValues[fieldName] = boundVal
		}
	}

	return result
}

func unbindStruct(output map[string]string, name string, iface interface{}) {
	val := reflect.ValueOf(iface)
	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		structField := typ.Field(i)
		fieldValue := val.Field(i)

		// PkgPath is specified to be empty exactly for exported fields.
		if structField.PkgPath == "" {
			revel.Unbind(output, fmt.Sprintf("%s.%s", name, structField.Name), fieldValue.Interface())
		}
	}
}
