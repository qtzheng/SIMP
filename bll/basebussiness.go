package bll

import (
	"fmt"
	"github.com/revel/revel"
	"gopkg.in/mgo.v2/bson"
)

var (
	DevMode bool
)

func init() {
	DevMode = revel.DevMode
}
func ErrorLog(err error) {
	revel.WARN.Fatalln(err)
}
func NewObjectID() bson.ObjectId {
	return bson.NewObjectId()
}
func Log(e interface{}) {
	fmt.Println(e)
}
