package controllers

import (
	"github.com/revel/revel"
)

type Index struct {
	Application
}

func (i Index) Index() revel.Result {
	return i.Render()
}
