package modules

import (
	"encoding/json"
)

type User struct {
	BaseModule
	UserId string `json:`
}
