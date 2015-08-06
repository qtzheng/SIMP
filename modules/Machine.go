package modules

import (
	//"github.com/revel/revel"
	"gopkg.in/mgo.v2/bson"
)

type Machine struct {
	ID       bson.ObjectId `bson:"_id"`
	Name     string
	Code     string
	IP       string
	ParentID string
}
