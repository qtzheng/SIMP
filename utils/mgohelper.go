package utils

import (
	"gopkg.in/mgo.v2"
)

func init() {
	session, _ := mgo.Dial("url")
	session.DB("name").C("name").Find("query")
}
