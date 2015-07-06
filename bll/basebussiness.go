package bll

import (
	"github.com/qtzheng/SIMP/utils"
	"github.com/revel/revel"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	url = "172.0.0.1"
)

var (
	MgoConn mgo.Session
)

func init() {
	MgoConn, err := mgo.Dial(url)
	if err != nil {
		if revel.DevMode {

		}

	}
}

type BaseBussiness struct {
	revel.DevMode
}
