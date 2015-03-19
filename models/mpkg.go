package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
}

type MPkg struct {
}

//获取包对应的app config
func (u *MPkg) GetPkgConfig(pkg string)map[string]string{
	result := map[string]string{}
	if len(pkg) > 0{
		o := orm.NewOrm()
			var maps []orm.Params
		num, err := o.Raw("SELECT * FROM t_config_pkg WHERE F_pkg=? LIMIT 1", pkg).Values(&maps)
		if err == nil && num > 0 {
			result["F_app_id"] = maps[0]["F_app_id"].(string)
			result["F_app_key"] = maps[0]["F_app_key"].(string)
			result["F_app_master_key"] = maps[0]["F_app_master_key"].(string)
			result["F_app_msm_template"] = maps[0]["F_app_msm_template"].(string)
			result["F_app_name"] = maps[0]["F_app_name"].(string)
		}
	}
	return result
}

//检查包是否存在
func (u *MPkg) CheckPkgExists(pkg string)bool{
	result := false
	if len(pkg) > 0{
		o := orm.NewOrm()
		var maps []orm.Params
		num, err := o.Raw("SELECT * FROM t_config_pkg WHERE F_pkg=? LIMIT 1", pkg).Values(&maps)
		if err == nil && num > 0 {
			result = true
		}
	}
	return result
}