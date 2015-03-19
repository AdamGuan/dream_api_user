package controllers

import (
	"dream_api_user/models"
	"github.com/astaxie/beego"
	"net/http"
	"dream_api_user/helper"
	"github.com/astaxie/beego/config" 
//	"fmt"
//	"strings"
//	"os"
)

//用户
type ConsumerController struct {
	beego.Controller
}

//上传文件用的size接口
type Sizer2 interface {
	Size() int64
}

//json echo
func (u0 *ConsumerController) jsonEcho(datas map[string]interface{},u *ConsumerController) {
	if datas["responseNo"] == -6 || datas["responseNo"] == -7 {
		u.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
		u.Ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
	} 
	
	datas["responseMsg"] = ""
	appConf, _ := config.NewConfig("ini", "conf/app.conf")
	debug,_ := appConf.Bool(beego.RunMode+"::debug")
	if debug{
		datas["responseMsg"] = models.ConfigMyResponse[helper.IntToString(datas["responseNo"].(int))]
	}

	u.Data["json"] = datas
	u.ServeJson()
}

//sign check, , token为包名的md5值
func (u0 *ConsumerController) checkSign(u *ConsumerController)int {
	result := -6
	pkg := u.Ctx.Request.Header.Get("pkg")
	sign := u.Ctx.Request.Header.Get("sign")
	var pkgObj *models.MPkg
	if !pkgObj.CheckPkgExists(pkg){
		result = -7
	}else{
		var signObj *models.MSign
		if re := signObj.CheckSign(sign,helper.Md5(pkg)); re == true {
			result = 0
		}
	}
	return result
}

// @Title 注册(手机号码注册)
// @Description 注册(手机号码注册)(token: md5(pkg))
// @Param	mobilePhoneNumber	form	string	true	手机号码
// @Param	pwd					form	string	true	密码
// @Param	num					form	string	true	验证码(经过验证成功后的)
// @Param	sign				header	string	true	签名
// @Param	pkg					header	string	true	包名
// @Success	200 {object} models.MResp
// @Failure 401 无权访问
// @router /phone-register [post]
func (u *ConsumerController) RegisterByPhone() {
	//ini return
	datas := map[string]interface{}{"responseNo": -1}
	//model ini
	var userObj *models.MConsumer
	var smsObj *models.MSms
	//parse request parames
	u.Ctx.Request.ParseForm()
	mobilePhoneNumber := u.Ctx.Request.FormValue("mobilePhoneNumber")
	pwd := u.Ctx.Request.FormValue("pwd")
	num := u.Ctx.Request.FormValue("num")
	pkg := u.Ctx.Request.Header.Get("pkg")
	//check sign
	datas["responseNo"] = u.checkSign(u)
	//检查参数
	if datas["responseNo"] == 0 && helper.CheckMPhoneValid(mobilePhoneNumber) && helper.CheckPwdValid(pwd) {
		datas["responseNo"] = -1
		if smsObj.CheckMsmActionvalid(mobilePhoneNumber,pkg,num) == true{
			parames := make(map[string]string)
			for k,v := range u.Ctx.Request.PostForm{
				parames[k] = v[0]
			}
			parames["mobilePhoneNumber"] = mobilePhoneNumber
			parames["pwd"] = pwd

			res2 := userObj.AddUserByPhone(parames)
			datas["responseNo"] = res2
		}
	}else if datas["responseNo"] == 0{
		datas["responseNo"] = -1
	}
	//return
	u.jsonEcho(datas,u)
}

// @Title 重置密码(利用手机号码重置密码)
// @Description 重置密码(利用手机号码重置密码)(token: md5(pkg))
// @Param	mobilePhoneNumber	form	string	true	手机号码
// @Param	pwd			form	string	true	密码
// @Param	num			form	string	true	验证码(经过验证成功后的)
// @Param	sign		header	string	true	签名
// @Param	pkg			header	string	true	包名
// @Success	200 {object} models.MResp
// @Failure 401 无权访问
// @router /resetpwd [put]
func (u *ConsumerController) ResetPwdByPhone() {
	//ini return
	datas := map[string]interface{}{"responseNo": -1}
	//model ini
	var userObj *models.MConsumer
	var smsObj *models.MSms
	//parse request parames
	u.Ctx.Request.ParseForm()
	mobilePhoneNumber := u.Ctx.Request.FormValue("mobilePhoneNumber")
	pwd := u.Ctx.Request.FormValue("pwd")
	num := u.Ctx.Request.FormValue("num")
	pkg := u.Ctx.Request.Header.Get("pkg")
	//check sign
	datas["responseNo"] = u.checkSign(u)
	//检查参数
	if datas["responseNo"] == 0 && helper.CheckMPhoneValid(mobilePhoneNumber) && helper.CheckPwdValid(pwd) {
		datas["responseNo"] = -1
		if smsObj.CheckMsmActionvalid(mobilePhoneNumber,pkg,num) == true{
			res2 := userObj.ModifyUserPwdByPhone(mobilePhoneNumber,pwd)
			datas["responseNo"] = res2
		}
	}else if datas["responseNo"] == 0{
		datas["responseNo"] = -1
	}
	//return
	u.jsonEcho(datas,u)
}

// @Title 登录(利用手机号码登录)
// @Description 登录(利用手机号码登录)(token: md5(pkg))
// @Param	mobilePhoneNumber	path	string	true	手机号码
// @Param	pwd			query	string	true	密码
// @Param	sign		header	string	true	签名
// @Param	pkg			header	string	true	包名
// @Success	200 {object} models.MUserLoginResp
// @Failure 401 无权访问
// @router /login/:mobilePhoneNumber [get]
func (u *ConsumerController) CheckUserAndPwdByPhone() {
	//ini return
	datas := map[string]interface{}{"responseNo": 0}
	//model ini
	var userObj *models.MConsumer
	//parse request parames
	u.Ctx.Request.ParseForm()
	mobilePhoneNumber := u.Ctx.Input.Param(":mobilePhoneNumber")
	pwd := u.Ctx.Request.FormValue("pwd")
	//check sign
	datas["responseNo"] = u.checkSign(u)
	//检查参数
	if datas["responseNo"] == 0 && helper.CheckMPhoneValid(mobilePhoneNumber) && helper.CheckPwdValid(pwd) {
		datas["responseNo"] = -1
		if !userObj.CheckPhoneExists(mobilePhoneNumber){
			datas["responseNo"] = -4
		}else{
			res := userObj.CheckPhoneAndPwd(mobilePhoneNumber,pwd)
			if res{
				datas["responseNo"] = 0
				//获取用户信息
				info := userObj.GetUserInfoByPhone(mobilePhoneNumber)
				if len(info.F_phone_number) > 0{
					datas["F_uid"] = info.F_uid
					datas["F_phone_number"] = info.F_phone_number
					datas["F_crate_datetime"] = info.F_crate_datetime
					datas["F_modify_datetime"] = info.F_modify_datetime
				}
			}else{
				datas["responseNo"] = -9
			}
		}
	}else if datas["responseNo"] == 0{
		datas["responseNo"] = -5
	}
	//return
	u.jsonEcho(datas,u)
}

// @Title 修改密码
// @Description 修改密码(token: md5(pkg))
// @Param	uid				path	string	true	用户ID
// @Param	oldPwd			form	string	true	旧密码
// @Param	newPwd			form	string	true	新密码
// @Param	sign			header	string	true	签名
// @Param	pkg			header	string	true	包名
// @Param	huid			header	string	true	用户ID
// @Success	200 {object} models.MResp
// @Failure 401 无权访问
// @router /pwd/:uid [put]
func (u *ConsumerController) ModifyPwdByUid() {
	//ini return
	datas := map[string]interface{}{"responseNo": -1}
	//model ini
	var userObj *models.MConsumer
	//parse request parames
	u.Ctx.Request.ParseForm()
	uid := u.Ctx.Input.Param(":uid")
	oldPwd := u.Ctx.Request.FormValue("oldPwd")
	newPwd := u.Ctx.Request.FormValue("newPwd")
	//check sign
	datas["responseNo"] = u.checkSign(u)
	//检查参数
	if datas["responseNo"] == 0 && helper.CheckPwdValid(oldPwd) && helper.CheckPwdValid(newPwd) {
		datas["responseNo"] = -1
		if userObj.CheckUserIdAndPwd(uid,oldPwd){
			res2 := userObj.ModifyUserPwdByUid(uid,newPwd)
			datas["responseNo"] = res2
		}else{
			datas["responseNo"] = -8
		}
	}else if datas["responseNo"] == 0{
		datas["responseNo"] = -1
	}
	//return
	u.jsonEcho(datas,u)
}

// @Title 验证手机号码是否已注册
// @Description 验证手机号码是否已注册(token: md5(pkg))
// @Param	mobilePhoneNumber	path	string	true	手机号码
// @Param	sign			header	string	true	签名
// @Param	pkg			header	string	true	包名
// @Success	200 {object} models.MResp
// @Failure 401 无权访问
// @router /exists/:mobilePhoneNumber [get]
func (u *ConsumerController) CheckUserExists() {
	//ini return
	datas := map[string]interface{}{"responseNo": -1}
	//model ini
	var userObj *models.MConsumer
	//parse request parames
	u.Ctx.Request.ParseForm()
	mobilePhoneNumber := u.Ctx.Input.Param(":mobilePhoneNumber")
	//check sign
	datas["responseNo"] = u.checkSign(u)
	//检查参数
	if datas["responseNo"] == 0 && helper.CheckMPhoneValid(mobilePhoneNumber) {
		if userObj.CheckPhoneExists(mobilePhoneNumber){
			datas["responseNo"] = -2
		}else{
			datas["responseNo"] = -4
		}
	}else if datas["responseNo"] == 0{
		datas["responseNo"] = -1
	}
	//return
	u.jsonEcho(datas,u)
}

// @Title 获取用户信息
// @Description 获取用户信息(token: md5(pkg))
// @Param	uid	path	string	true	用户ID
// @Param	sign			header	string	true	签名
// @Param	pkg			header	string	true	包名
// @Param	huid		header	string	true	用户ID
// @Success	200 {object} models.MUserInfoResp
// @Failure 401 无权访问
// @router /:uid [get]
func (u *ConsumerController) GetUserInfo() {
	//ini return
	datas := map[string]interface{}{"responseNo": -1}
	//model ini
	var userObj *models.MConsumer
	//parse request parames
	u.Ctx.Request.ParseForm()
	uid := u.Ctx.Input.Param(":uid")
	//check sign
	datas["responseNo"] = u.checkSign(u)
	//检查参数
	if datas["responseNo"] == 0 {
		datas["responseNo"] = -1
		info := userObj.GetUserInfoByUid(uid)
		if len(info.F_phone_number) > 0{
			datas["responseNo"] = 0
			datas["F_uid"] = info.F_uid
			datas["F_phone_number"] = info.F_phone_number
			datas["F_crate_datetime"] = info.F_crate_datetime
			datas["F_modify_datetime"] = info.F_modify_datetime
		}
	}else if datas["responseNo"] == 0{
		datas["responseNo"] = -1
	}
	//return
	u.jsonEcho(datas,u)
}

// @Title 修改用户手机号码
// @Description 修改用户手机号码(token: md5(pkg))
// @Param	mobilePhoneNumber	path	string	true	手机号码(旧的手机号码)
// @Param	newPhone			form	string	true	手机号码(新的手机号码)
// @Param	num					form	string	true	验证码(经过验证成功后的)
// @Param	sign				header	string	true	签名
// @Param	pkg					header	string	true	包名
// @Success	200 {object} models.MModifyPhoneResp
// @Failure 401 无权访问
// @router /phone/:mobilePhoneNumber [put]
func (u *ConsumerController) ModifyPhone() {
	//ini return
	datas := map[string]interface{}{"responseNo": -1}
	//model ini
	var userObj *models.MConsumer
	var smsObj *models.MSms
	//parse request parames
	u.Ctx.Request.ParseForm()
	mobilePhoneNumber := u.Ctx.Input.Param(":mobilePhoneNumber")
	newPhone := u.Ctx.Request.FormValue("newPhone")
	num := u.Ctx.Request.FormValue("num")
	pkg := u.Ctx.Request.Header.Get("pkg")
	//check sign
	datas["responseNo"] = u.checkSign(u)
	//检查参数
	if datas["responseNo"] == 0 && helper.CheckMPhoneValid(mobilePhoneNumber) && helper.CheckMPhoneValid(newPhone) {
		datas["responseNo"] = -1
		if smsObj.CheckMsmActionvalid(newPhone,pkg,num) == true{
			datas["responseNo"] = userObj.ModifyUserPhone(mobilePhoneNumber,newPhone)
		}
	}else if datas["responseNo"] == 0{
		datas["responseNo"] = -10
	}
	//return
	u.jsonEcho(datas,u)
}