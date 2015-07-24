package modules

import (
//"encoding/json"
)

type User struct {
	UserID    string `bson:"_id"`
	LoginName string
	Password  string
	UserName  string
}
