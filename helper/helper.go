package helper

import (
	"crypto/md5"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	//"net/url"
	"crypto/sha1"
	"github.com/astaxie/beego/config"
	"regexp"
	"time"
)

var MyLog *logs.BeeLogger

func init() {
	//初始化log
	appConf, _ := config.NewConfig("ini", "conf/app.conf")
	debug, _ := appConf.Bool(beego.RunMode + "::debug")
	if debug {
		MyLog = logs.NewLogger(10000)
		MyLog.SetLogger("file", `{"filename":"log.log"}`)
		MyLog.EnableFuncCallDepth(true)
		//MyLog.Debug("debug test1")
		// MyLog.Error("error")
	}
}

//类型转化 string  to int
func StrToInt(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

//类型转化 string  to float64
func StrToFloat64(str string) float64 {
	f, _ := strconv.ParseFloat(str, 64)
	return f
}

//类型转化 int to string
func IntToString(i int) string {
	return fmt.Sprintf("%d", i)
}

//类型转化 int64 to string
func Int64ToString(i int64) string {
	return fmt.Sprintf("%d", i)
}

//获取16位Guid
func GetGuid() string {
	f, _ := os.OpenFile("/dev/urandom", os.O_RDONLY, 0)
	b := make([]byte, 16)
	f.Read(b)
	f.Close()
	uuid := fmt.Sprintf("%x", b)
	return uuid[0:16]
}

//创建密码
func CreatePwd(num int) string {
	return GetGuid()[0:num]
}

//leanCloud curl
func CurlLeanCloud(requestUri string, method string, requestData map[string]string, appId string, appKey string) map[string]interface{} {
	geturl := requestUri
	req, _ := http.NewRequest(method, geturl, nil)
	data, _ := json.Marshal(requestData)
	req.Body = ioutil.NopCloser(strings.NewReader(string(data)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("User-Agent", "SSTS Browser/1.0")
	//req.Header.Add("X-AVOSCloud-Application-Id", "l7dn2hfzry3yxdtfhnhciy0jt00dqd3ysbees8pjdhonwjb4")
	req.Header.Add("X-AVOSCloud-Application-Id", appId)
	//req.Header.Add("X-AVOSCloud-Application-Key", "76nib8wzexxcgaccrl9e4955ll8q11w1jwq36ofpx6u340q2")
	req.Header.Add("X-AVOSCloud-Application-Key", appKey)
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	resp, _ := client.Do(req)
	defer resp.Body.Close()
	bodyByte, _ := ioutil.ReadAll(resp.Body)
	p := map[string]interface{}{}
	json.Unmarshal(bodyByte, &p)
	return p
}

//记录短信发送的 curl
func CurlSmsLog(requestUri string, method string, requestData map[string]string) (map[string]interface{}, map[string][]string) {
	fmt.Println(requestUri)
	fmt.Println(method)
	fmt.Println(requestData)

	geturl := requestUri
	req, _ := http.NewRequest(method, geturl, nil)
	data, _ := json.Marshal(requestData)
	req.Body = ioutil.NopCloser(strings.NewReader(string(data)))
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	resp, _ := client.Do(req)
	defer resp.Body.Close()
	bodyByte, _ := ioutil.ReadAll(resp.Body)
	p := map[string]interface{}{}
	json.Unmarshal(bodyByte, &p)
	return p, resp.Header
}

//检查签名
func CheckSign(sign string, token string) bool {
	//sign = timestamp+md5(token+timestamp)
	if len(sign) == 46 && len(token) == 32 {
		timestamp := sign[0:14]
		//检测是否超时
		//		appConf, _ := config.NewConfig("ini", "conf/app.conf")
		//		debug,_ := appConf.Bool(beego.RunMode+"::debug")
		//		if !debug{
		if 1 != 1 {
			nowTime, _ := strconv.Atoi(time.Now().Format("20060102150405"))
			requestTime, _ := strconv.Atoi(timestamp)
			timedistince := nowTime - requestTime
			if timedistince > 60*5 {
				return false
			}
		}

		md5Str := sign[14:]
		str := token + timestamp
		md5Str2 := fmt.Sprintf("%x\r\n", md5.Sum([]byte(str)))
		md5Str2 = strings.TrimSpace(md5Str2)
		if string(md5Str) == string(md5Str2) {
			return true
		}
	}
	return false
}

//手机号码有效性验证
func CheckMPhoneValid(phone string) bool {
	matched, err := regexp.MatchString("^1[3|4|5|6|7|8][0-9]{9}$", phone)
	if err == nil && matched {
		return true
	}
	return false
}

//密码有效性验证
func CheckPwdValid(pwd string) bool {
	//pwd = strings.TrimSpace(pwd)
	matched, err := regexp.MatchString("^\\w{6,40}$", pwd)
	if err == nil && matched {
		return true
	}
	return false
}

//md5
func Md5(str string) string {
	md5Str := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return md5Str
}

//get now datatime
func GetNowDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func GetDateTimeBeforeMinute(num int) string {
	return time.Now().Add(-time.Minute * time.Duration(num)).Format("2006-01-02 15:04:05")
}

func GetDateTimeAfterMinute(num int) string {
	return time.Now().Add(time.Minute * time.Duration(num)).Format("2006-01-02 15:04:05")
}

func Split(str string, flag string) []string {
	return strings.Split(str, ",")
}

func JoinString(list []string, flag string) string {
	result := ""
	if len(list) > 0 {
		for _, v := range list {
			result += v + ","
		}
		result = strings.Trim(result, ",")
	}
	return result
}

func StringInArray(value string, list []string) bool {
	result := false
	for _, item := range list {
		if value == item {
			result = true
			break
		}
	}
	return result
}

func Sha1(str string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(str)))
}
