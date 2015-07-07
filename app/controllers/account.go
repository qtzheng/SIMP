package controllers

import (
	"github.com/revel/revel"
)

type Account struct {
	Application
}

func (c Account) Login() revel.Result {
	return c.Render()
}
