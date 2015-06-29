package controllers

import (
	"github.com/revel/revel"
)

type Role struct {
	Application
}

func (c Role) Index() revel.Result {
	return c.Render()
}
