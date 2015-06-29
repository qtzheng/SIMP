package controllers

import (
	"github.com/revel/revel"
)

type Department struct {
	Application
}

func (c Department) Index() revel.Result {
	return c.Render()
}
