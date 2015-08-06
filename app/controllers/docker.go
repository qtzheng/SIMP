package controllers

import (
	"github.com/qtzheng/SIMP/bll"
	"github.com/qtzheng/SIMP/modules"
	"github.com/revel/revel"
	"gopkg.in/mgo.v2/bson"
	"strings"
)

type Docker struct {
	BaseCollection
}

func (d Docker) Index() {

}
