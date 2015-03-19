package main

import (
	_ "dream_api_user/docs"
	_ "dream_api_user/routers"
	"dream_api_user/controllers"

	"github.com/astaxie/beego"
	"runtime"
	"github.com/astaxie/beego/config" 
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	
	appConf, _ := config.NewConfig("ini", "conf/app.conf")
	debug,_ := appConf.Bool(beego.RunMode+"::debug")
	if debug{
		beego.StaticDir["/swagger"] = "swagger"
	}
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}
