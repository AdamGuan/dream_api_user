package controllers

import (
	"dream_api_user/models"
	"github.com/astaxie/beego"
	"dream_api_user/helper"
	"github.com/astaxie/beego/config" 
	"github.com/astaxie/beego/utils"
	"github.com/astaxie/beego/logs"
)

type ErrorController struct {
	beego.Controller
}

//json echo
func (u0 *ErrorController) jsonEcho(datas map[string]interface{},u *ErrorController) {
//	u.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
	u.Ctx.Output.ContentType("application/json; charset=utf-8")
	
	datas["responseMsg"] = ""
	appConf, _ := config.NewConfig("ini", "conf/app.conf")
	debug,_ := appConf.Bool(beego.RunMode+"::debug")
	if debug{
		datas["responseMsg"] = models.ConfigMyResponse[helper.IntToString(datas["responseNo"].(int))]
	}
	u.Data["json"] = datas
	u.ServeJson()
}

func (u *ErrorController) Error401() {
	//ini return
	u.Ctx.Output.SetStatus(401)
	datas := map[string]interface{}{"responseNo": 401}
	//return
	u.jsonEcho(datas,u)
}

func (u *ErrorController) Error403() {
	//ini return
	u.Ctx.Output.SetStatus(403)
	datas := map[string]interface{}{"responseNo": 403}
	//return
	u.jsonEcho(datas,u)
}

func (u *ErrorController) Error404() {
	//ini return
	u.Ctx.Output.SetStatus(404)
	datas := map[string]interface{}{"responseNo": 404}
	//return
	u.jsonEcho(datas,u)
}

func (u *ErrorController) Error500() {
	//ini return
	u.Ctx.Output.SetStatus(500)
	datas := map[string]interface{}{"responseNo": 500}
	//return
	u.jsonEcho(datas,u)
}

func (u *ErrorController) Error503() {
	//ini return
	u.Ctx.Output.SetStatus(503)
	datas := map[string]interface{}{"responseNo": 503}
	//log
	//"Body", u.Ctx.Input.Body(), 
	str := "503\n"+utils.GetDisplayString("Body", u.Ctx.Input.CopyBody(), "IP", u.Ctx.Input.IP(), "Uri", u.Ctx.Input.Uri(),"Method", u.Ctx.Input.Method())
	log := logs.NewLogger(10000)
	log.SetLogger("file", `{"filename":"test.log"}`)
	log.Error(str)
	//return
	u.jsonEcho(datas,u)
}