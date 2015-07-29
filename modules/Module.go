package modules

import (
	//"github.com/revel/revel"
	"gopkg.in/mgo.v2/bson"
)

type Module struct {
	ID          bson.ObjectId `bson:"_id"`
	Name        string
	Code        string
	ParentID    string
	URL         string
	IsUse       bool
	Sort        int
	Remark      string
	DisplayType int
}
