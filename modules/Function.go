package modules

import (
	//"github.com/revel/revel"
	"gopkg.in/mgo.v2/bson"
)

type Function struct {
	ID       bson.ObjectId `bson:"_id"`
	Name     string
	Code     string
	ModuleID string
	IsUse    bool
	Remark   string
}
