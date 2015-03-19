package models

import (
	//"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"time"
	"github.com/astaxie/beego/config" 
)

//key:响应代码，value:响应信息
var ConfigMyResponse map[string]string

func init() {
	dbconf, _ := config.NewConfig("ini", "conf/db.conf")
	maxIdle,_ := dbconf.Int("maxIdle")
	maxConn,_ := dbconf.Int("maxConn")
	userName := dbconf.String(beego.RunMode+"::userName")
	password := dbconf.String(beego.RunMode+"::password")
	dbName := dbconf.String(beego.RunMode+"::dbName")
	orm.RegisterDataBase("default", "mysql", userName+":"+password+"@/"+dbName+"?charset=utf8&loc=Asia%2FShanghai",maxIdle, maxConn)
	orm.DefaultTimeLoc = time.UTC
	appConf, _ := config.NewConfig("ini", "conf/app.conf")
	debug,_ := appConf.Bool(beego.RunMode+"::debug")
	if debug{
		orm.Debug = true
	}
	logFile, _ := os.OpenFile("./db.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	orm.DebugLog = orm.NewLog(logFile)

	getResponseConfig()	
}

//获取config  im
func getResponseConfig() {
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("SELECT * FROM t_config_response").Values(&maps)
	if err == nil && num > 0 {
		ConfigMyResponse = make(map[string]string)
		for _, item := range maps {
			ConfigMyResponse[item["F_response_no"].(string)] = item["F_response_msg"].(string)
		}
	}
}