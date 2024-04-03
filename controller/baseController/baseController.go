package baseController

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"log"
)

type BaseController struct {
	beego.Controller
}

func (B *BaseController) GetInput() map[string]interface{} {
	res := map[string]interface{}{}
	input := B.Ctx.Input.RequestBody
	err := json.Unmarshal(input, &res)
	if err != nil {
		log.Printf("Failed to unmarshal Input RequestBody, and err is: %s\n", err.Error())
		return res
	}
	return res
}
