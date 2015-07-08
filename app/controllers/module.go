package controllers

import (
	"github.com/revel/revel"
)

type Module struct {
	Application
}

func (c Module) Index() revel.Result {
	return c.Render()
}
func (m Module) GetFuncs() revel.Result {
	return m.RenderText("")
}
