package bll

import (
//"github.com/qtzheng/SIMP/dal"
"gopkg.in/mgo.v2"
"gopkg.in/mgo.v2/bson"
)

type RoleBussiness struct {
	BaseBussiness
}

func (r *RoleBussiness) coll() *mgo.Collection{
	return r.DB().("Role")
}
func (r *RoleBussiness) CreateRoleTree() {
	r.coll().Find(bson.M{"":""})
}
