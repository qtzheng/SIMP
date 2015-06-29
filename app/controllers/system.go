package controllers

import (
	"github.com/revel/revel"
)

type System struct {
	Application
}

func (c System) Index() revel.Result {
	return c.Render()
}
