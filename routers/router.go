package routers

import (
	"beego/controller/lifegameController"

	"github.com/astaxie/beego"
)

func init() {

	//views 自用接口
	//beego.Router("/autofresh", &lifegameController.LifegameController{}, "get:Autofresh")
	//beego.Router("/reset", &lifegameController.LifegameController{}, "get:Reset")
	//beego.Router("/test", &lifegameController.LifegameController{}, "get:Test")
	//beego.Router("/getMap", &lifegameController.LifegameController{}, "get:GetMap")
	//beego.Router("/getSlice", &lifegameController.LifegameController{}, "get:GetSlice")

	//fetch专用接口
	beego.Router("/fetch/autofresh", &lifegameController.LifegameController{}, "get:FetchAutofresh")
	beego.Router("/fetch/reset", &lifegameController.LifegameController{}, "get:FetchReset")
	beego.Router("/fetch/testPost", &lifegameController.LifegameController{}, "post:TestPost")
}
