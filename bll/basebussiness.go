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

type BaseBussiness struct {
}
