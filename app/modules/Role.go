package modules

import (
	"github.com/revel/revel"
	"gopkg.in/mgo.v2/bson"
)

type Role struct {
	RoleID   bson.ObjectId `"bson":_id`
	RoleName string
	RoleCode string
	IsUse    bool
	Remark   string
	ParentID bson.ObjectId
}

func (r *Role) Validate(v *revel.Validation) {
	v.Check(r.IsUse, revel.Required{})
	v.Check(r.ParentID, revel.MaxSize{50})
	v.Check(r.RoleCode, revel.Required{}, revel.MaxSize{50}, revel.MinSize{10})
	v.Check(r.RoleName, revel.Required{}, revel.MaxSize{10}, revel.MinSize{2})
}
