package routers

import (
	"beego/lifegame"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &lifegame.LifegameController{}, "get:GetIndex")
	beego.Router("/lifegame", &lifegame.LifegameController{})
}