package utils

import (
	"github.com/revel/revel"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func init() {
	session, _ := mgo.Dial("url")
	session.DB("name").C("name").Find("query")
}
