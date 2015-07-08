package modules

import (
	"encoding/json"
	"github.com/qtzheng/SIMP/utils"
	"github.com/revel/revel"
)

type BaseModule struct {
	revel.DevMode
}

func (b *BaseModule) ConvertToJson() string {
	js, err := json.Marshal(b)
	if err != nil {
		if b.DevMode {
			panic(err)
		}
		return ""
	} else {
		return js
	}
}
