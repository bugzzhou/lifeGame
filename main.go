package main

import (
	"github.com/astaxie/beego"

	_ "beego/routers"
)

func main() {
	beego.Run(":9292")
}
