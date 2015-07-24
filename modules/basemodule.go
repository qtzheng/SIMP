package modules

import (
	"encoding/json"
	//"github.com/qtzheng/SIMP/utils"
	"github.com/revel/revel"
)

type BaseModule struct {
}

func (b *BaseModule) ConvertToJson() string {
	js, err := json.Marshal(b)
	if err != nil {
		if revel.DevMode {
			panic(err)
		}
		return ""
	} else {
		return string(js)
	}
}
