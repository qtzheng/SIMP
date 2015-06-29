package controllers

import (
	"github.com/revel/revel"
)

type Account struct {
	Application
}

func (c Account) Index() revel.Result {
	return c.Render()
}
