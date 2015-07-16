package bll

import (
	"github.com/revel/revel"
)

var (
	DevMode bool
)

func init() {
	DevMode = revel.DevMode
}
func ErrorLog(err error) {
	revel.WARN.Fatalln(err)
}
