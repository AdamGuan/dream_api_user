package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	
	beego.GlobalControllerRouter["dream_api_user/controllers:ConsumerController"] = append(beego.GlobalControllerRouter["dream_api_user/controllers:ConsumerController"],
		beego.ControllerComments{
			"RegisterByPhone",
			`/phone-register`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["dream_api_user/controllers:ConsumerController"] = append(beego.GlobalControllerRouter["dream_api_user/controllers:ConsumerController"],
		beego.ControllerComments{
			"ResetPwdByPhone",
			`/resetpwd`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["dream_api_user/controllers:ConsumerController"] = append(beego.GlobalControllerRouter["dream_api_user/controllers:ConsumerController"],
		beego.ControllerComments{
			"CheckUserAndPwdByPhone",
			`/login/:mobilePhoneNumber`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["dream_api_user/controllers:ConsumerController"] = append(beego.GlobalControllerRouter["dream_api_user/controllers:ConsumerController"],
		beego.ControllerComments{
			"ModifyPwdByUid",
			`/pwd/:uid`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["dream_api_user/controllers:ConsumerController"] = append(beego.GlobalControllerRouter["dream_api_user/controllers:ConsumerController"],
		beego.ControllerComments{
			"CheckUserExists",
			`/exists/:mobilePhoneNumber`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["dream_api_user/controllers:ConsumerController"] = append(beego.GlobalControllerRouter["dream_api_user/controllers:ConsumerController"],
		beego.ControllerComments{
			"GetUserInfo",
			`/:uid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["dream_api_user/controllers:ConsumerController"] = append(beego.GlobalControllerRouter["dream_api_user/controllers:ConsumerController"],
		beego.ControllerComments{
			"ModifyPhone",
			`/phone/:mobilePhoneNumber`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["dream_api_user/controllers:SmsController"] = append(beego.GlobalControllerRouter["dream_api_user/controllers:SmsController"],
		beego.ControllerComments{
			"Smsvalid",
			`/smsvalid/:mobilePhoneNumber`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["dream_api_user/controllers:SmsController"] = append(beego.GlobalControllerRouter["dream_api_user/controllers:SmsController"],
		beego.ControllerComments{
			"RegisterGetSms",
			`/register/:mobilePhoneNumber`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["dream_api_user/controllers:SmsController"] = append(beego.GlobalControllerRouter["dream_api_user/controllers:SmsController"],
		beego.ControllerComments{
			"ResetPwdGetSms",
			`/resetpwd/:mobilePhoneNumber`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["dream_api_user/controllers:SmsController"] = append(beego.GlobalControllerRouter["dream_api_user/controllers:SmsController"],
		beego.ControllerComments{
			"ChangePhoneSms",
			`/phone/:mobilePhoneNumber`,
			[]string{"get"},
			nil})

}
