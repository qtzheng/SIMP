package controllers

import (
	"github.com/revel/revel"
)

type BaseCollection struct {
	*revel.Controller
}

type opResult struct {
	Result  int
	Message interface{}
}

const (
	Success = 0 + iota
	Error
)

func returnMessage(c *revel.Controller, message interface{}, err error) revel.Result {
	result := &opResult{}
	if err != nil {
		result.Result = Error
		result.Message = err.Error()
	} else {
		result.Result = Success
		result.Message = message
	}

	return c.RenderJson(result)
}
