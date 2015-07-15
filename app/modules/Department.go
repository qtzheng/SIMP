package modules

import (
	"github.com/revel/revel"
	"gopkg.in/mgo.v2/bson"
)

type Department struct {
	ID       bson.ObjectId `bson:_id`
	Name     string
	Code     string
	Manager  string
	ParentID bson.ObjectId
	IsUse    bool
	Remark   string
}

func (d *Department) Validate(v *revel.Validation) {
	v.Check(d.IsUse, revel.Required{})
	v.Check(d.ParentID, revel.MaxSize{50})
	v.Check(d.Code, revel.Required{}, revel.MaxSize{50}, revel.MinSize{10})
	v.Check(d.Name, revel.Required{}, revel.MaxSize{10}, revel.MinSize{2})
}
