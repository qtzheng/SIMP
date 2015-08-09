package controllers

import (
	"github.com/qtzheng/SIMP/bll"
	"github.com/qtzheng/SIMP/modules"
	"github.com/revel/revel"
	//"gopkg.in/mgo.v2/bson"
	//"strings"
)

type Docker struct {
	BaseCollection
}

func (d Docker) Index() revel.Result {
	return d.Render()
}
func (d Docker) MachineAdd(machine *modules.Machine) revel.Result {
	err := bll.MachineAdd(machine)
	return returnMessage(d.Controller, machine.ID, err)
}
