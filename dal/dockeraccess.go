package dal

import (
	"fmt"
	"github.com/qtzheng/SIMP/modules"
	"gopkg.in/mgo.v2/bson"
)

func MachineAdd(machine *modules.Machine) error {
	coll := CloneDB().C(MachineColl)
	count, err := coll.Find(bson.M{"code": machine.Code}).Count()
	if err != nil {
		return err
	}
	if count == 0 {
		err = coll.Insert(machine)
		return err
	} else {
		return fmt.Errorf("该编码已经存在：%s", machine.Code)
	}
}
func MachineUpdate(machine *modules.Machine) error {
	return CloneDB().C(MachineColl).Update(bson.M{"_id": machine.ID}, machine)
}
func MachineSelect() (*[]modules.Machine, error) {
	list := &[]modules.Machine{}
	err := CloneDB().C(MachineColl).Find(nil).Select(bson.M{"_id": 1, "name": 1}).All(list)
	return list, err
}
