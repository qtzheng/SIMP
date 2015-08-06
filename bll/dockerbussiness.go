package bll

import (
	"fmt"
	"github.com/qtzheng/SIMP/dal"
	"github.com/qtzheng/SIMP/modules"
	"gopkg.in/mgo.v2/bson"
	//"strings"
)

func MachineAdd(machine *modules.Machine) error {
	machine.ID = bson.NewObjectId()

}
