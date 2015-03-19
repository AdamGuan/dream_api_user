package models

import (
	"dream_api_user/helper"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"github.com/astaxie/beego/config" 
)

var SmsMinute int

func init() {
	otherconf, _ := config.NewConfig("ini", "conf/other.conf")
	SmsMinute,_ = otherconf.Int("smsMinute")
}

type MSms struct {
}

//get a msm
func (u *MSms) GetMsm(mobilePhoneNumber string,appId string,appKey string,appName string,appTemplate string) map[string]interface{} {
	return helper.CurlLeanCloud("https://leancloud.cn/1.1/requestSmsCode","POST",map[string]string{"mobilePhoneNumber": mobilePhoneNumber,"template":appTemplate,"appname":appName},appId,appKey);
}

//valid a msm
func (u *MSms) ValidMsm(num string,mobilePhoneNumber string,appId string,appKey string) map[string]interface{} {
	return helper.CurlLeanCloud("https://leancloud.cn/1.1/verifySmsCode/"+num+"?mobilePhoneNumber="+mobilePhoneNumber,"POST",map[string]string{},appId,appKey);
}

//检查是否可以给用户发送短信了
func (u *MSms) CheckMsmRateValid(phone string,pkgName string)bool{
	o := orm.NewOrm()
	var maps []orm.Params
	nowTime := time.Now().Add(-time.Minute * time.Duration(SmsMinute)).Format("2006-01-02 15:04:05")
	num, err := o.Raw("SELECT F_last_timestamp FROM t_sms_rate WHERE F_action=? LIMIT 1", helper.Md5(phone+pkgName)).Values(&maps)
	if err == nil && num > 0 {
		if maps[0]["F_last_timestamp"].(string) <= nowTime{
			return true
		}else{
			return false
		}
	}else{
		return true
	}
	return false
}

//写入短信发送频率表
func (u *MSms) AddMsmRate(phone string,pkgName string){
	//写入数据库
	o := orm.NewOrm()
	o.Raw("replace into t_sms_rate(F_action,F_last_timestamp) values('"+helper.Md5(phone+pkgName)+"','"+time.Now().Format("2006-01-02- 15:04:05")+"')").Exec()
}

//删除短信发送频率表
func (u *MSms) DeleteMsmRate(phone string,pkgName string){
	o := orm.NewOrm()
	o.Raw("UPDATE t_sms_rate SET F_last_timestamp='1001-01-01 00:00:00' WHERE F_action=?",helper.Md5(phone+pkgName)).Exec()
}

//写入t_sms_action_valid
func (u *MSms) AddMsmActionvalid(phone string,pkgName string,sms string){
	//写入数据库
	o := orm.NewOrm()
	o.Raw("replace into t_sms_action_valid(F_action) values('"+helper.Md5(phone+pkgName+sms)+"')").Exec()
}

//删除t_sms_action_valid
func (u *MSms) DeleteMsmActionvalid(phone string,pkgName string,sms string){
	o := orm.NewOrm()
	o.Raw("DELETE FROM t_sms_action_valid WHERE F_action=?",helper.Md5(phone+pkgName+sms)).Exec()
}

//验证t_sms_action_valid
func (u *MSms) CheckMsmActionvalid(phone string,pkgName string,sms string)bool{
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw("SELECT F_action FROM t_sms_action_valid WHERE F_action=? LIMIT 1", helper.Md5(phone+pkgName+sms)).Values(&maps)
	if err == nil && num > 0 {
		u.DeleteMsmActionvalid(phone,pkgName,sms)
		return true
	}
	return false
}