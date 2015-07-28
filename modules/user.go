package modules

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	UserID       bson.ObjectId `bson:"_id"`
	LoginName    string
	Password     string
	JobNumber    string
	UserName     string
	DepID        string
	RoleIDs      []string
	EngName      string
	PinYin       string
	Abbreviation string
	Telephone    string
	Mobilephone  string
	Email        string
	IsUse        bool
}
