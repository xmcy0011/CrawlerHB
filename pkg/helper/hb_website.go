package helper

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/CrawlerHB/pkg/logger"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"
)

var (
	UrlGetUserPassSd   = "/onlinemonitor/index.html"
	UrlGetSessionID    = "/onlinemonitor/mon/psBaseInfo/systemName"
	UrlGetQrCodeHeight = "/onlinemonitor/mon/getHeight"
	UrlSetQrCodeHeight = "/onlinemonitor/mon/checkLenght"
	UrlLogin           = "/onlinemonitor/mon/sysuser/login"
	UrlPortInfo        = "/onlinemonitor/mon/portInfo/get"
	UrlGetPortList     = "/onlinemonitor//mon/portInfo/findSelectListByPsIdPortTypeCode"
	UrlDataQueryHb     = "/onlinemonitor/mon/dataQueryHb/page/search/chart"
	UrlDataChatHeader  = "/onlinemonitor/mon/dataQueryHb/search/header"
)

// 监测点信息
type PortInfo struct {
	Color    string `json:"color"`
	PsInfoId string `json:"psinfoid"` // ??
	Name     string `json:"name"`     // 名称
	Id       string `json:"id"`
	Dgi      string `json:"dgi_mn"`                // ?
	Code     string `json:"monitorpointtype_code"` // ??
}

// 历史数据的表头
type PortChatHeader struct {
	Id          string `json:"id"`
	Name        string `json:"name"` //监测时间
	Type        string `json:"type"`
	Show        bool   `json:"show"`
	Render      string `json:"render"`
	Width       string `json:"width"`
	ExportWidth string `json:"exportWidth"`
	Checked     bool   `json:"checked"`
	Ordering    bool   `json:"ordering"`
	Children    string `json:"children"`
}

type PortChatData struct {
	//"date": "2020-02-01 00:00:00",
	//"03-zsavg": "62.30$50.482499999999995",
	//"S06-avg": "-$1.0",
	//"03-avg": "62.30$50.48291666666666",
	//"02-zsavg": "1.08$3.3470833333333334",
	//"S03-avg": "63.11$58.86958333333331",
	//"01-zsavg": "2.33$1.6291666666666667",
	//"S01-avg": "13.35$11.805000000000001",
	//"S05-avg": "3.12$1.493333333333333",
	//"S07-avg": "-$1.0",
	//"01-avg": "2.33$1.6291666666666667",
	//"02-avg": "1.08$3.3470833333333334",
	//"S08-avg": "-87.87$-77.08208333333333",
	//"S04-avg": "-$1.0",
	//"S02-avg": "7.10$6.945416666666667"

	//"id": "01-avg",
	//"name": "烟尘实测浓度(mg/m³)",
	//"type": null,
	//"show": true,
	//"render": null,
	//"width": null,
	//"exportWidth": null,
	//"checked": true,
	//"ordering": false,
	//"children": null
}

type HBWebSite struct {
	host      string
	sessionId string
	user      string
	pwd       string
	sid       string
}

// 发送GET请求
// url：         请求地址，不需要host部分
// response：    请求返回的内容
func (h *HBWebSite) get(url string) (string, error) {
	// 超时时间：10秒
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(h.host + ":" + url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// 发送POST请求
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json
func (h *HBWebSite) post(url string, data interface{}, contentType string) (string, error) {
	// 超时时间：10秒
	client := &http.Client{Timeout: 10 * time.Second}

	jsonStr := []byte("")
	if data != nil {
		str, err := json.Marshal(data)
		if err != nil {
			return "", err
		}
		jsonStr = str
	}

	resp, err := client.Post(h.host+":"+url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

// 带cookie的post
func (h *HBWebSite) postWitchCookie(url string, data string, contentType string) (string, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("POST", h.host+":"+url, strings.NewReader(data))
	if err != nil {
		return "", err
	}

	cookie1 := &http.Cookie{Name: "JSESSIONID", Value: h.sessionId, HttpOnly: true}
	req.AddCookie(cookie1)

	req.Header.Add("Content-type", contentType)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

// 1-获取用户名、密码、SID
func (h *HBWebSite) getUserPassSd() (user, pwd, sid string, error error) {
	rsp, err := http.Get(h.host + ":" + UrlGetUserPassSd)
	if err != nil {
		logger.Sugar.Error(err)
		return "", "", "", err
	}

	data, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		logger.Sugar.Error(err)
		return "", "", "", err
	}

	html := string(data)

	r, _ := regexp.Compile("密码:(.*?)</span>")
	if strArr := r.FindAllString(html, -1); len(strArr) > 0 {
		pwd = strings.ReplaceAll(strArr[0], "密码:", "") //还有一个，123
		pwd = strings.ReplaceAll(pwd, "</span>", "")   //还
	} else {
		logger.Sugar.Error("密码匹配错误")
		return "", "", "", errors.New("")
	}

	// 公众用户访问，这2个不会变
	user = "public"
	sid = "d211c9993e4844cebe510db6d4a7bd7d" // 'sdId':

	logger.Sugar.Infof("getUserPassSd ,user=%s,pwd=%s,sid=%s", user, pwd, sid)
	return user, pwd, sid, nil
}

// 2-获取会话ID
func (h *HBWebSite) getSessionID() (sessionId string, error error) {
	client := &http.Client{Timeout: 10 * time.Second}
	jsonStr := []byte("")
	rsp, err := client.Post(h.host+":"+UrlGetSessionID, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		return "", err
	}

	sessionId = ""
	for i := range rsp.Cookies() {
		if rsp.Cookies()[i].Name == "JSESSIONID" {
			sessionId = rsp.Cookies()[i].Value
		}
	}

	if sessionId == "" {
		return "", errors.New("解析JSESSIONID错误")
	}
	h.sessionId = sessionId
	logger.Sugar.Infof("sessionId:%s", sessionId)
	return sessionId, nil
}

// 3-获取验证码
func (h *HBWebSite) getQrCode(sessionId string) (string, error) {
	h.sessionId = sessionId
	resp, err := h.postWitchCookie(UrlGetQrCodeHeight, ` {"storagePosition": "database"}`, "application/json")
	if err == nil {
		logger.Sugar.Infof("getQrCode %s", resp)
	}
	return resp, err
}

// 4-验证码校验
func (h *HBWebSite) checkQrCode(sessionId, timeTemple string, xLenght int) (string, error) {
	h.sessionId = sessionId
	data := fmt.Sprintf(`{"timeTemple": "%s", "xLenght": %d, "storagePosition": "database"}`, timeTemple, xLenght)
	resp, err := h.postWitchCookie(UrlSetQrCodeHeight, data, "application/json")
	if err == nil {
		logger.Sugar.Infof("checkQrCode %s", resp)
	}
	return resp, err
}

// 5-登录
func (h *HBWebSite) login(sessionId, user, pwd, sid string) (string, error) {
	jsonData := fmt.Sprintf(`{"suLoginid": "%s", "suPasswd": "%s", "sdId": "%s"}`, user, pwd, sid)
	h.sessionId = sessionId
	resp, err := h.postWitchCookie(UrlLogin, jsonData, "application/json")
	if err == nil {
		logger.Sugar.Infof("login %s", resp)
	}
	return resp, err
}

// 登录
func (h *HBWebSite) Login(host string) (bool, error) {
	//var hb = &HBWebSite{Host: "http://110.249.223.71:9030"}
	h.host = host
	user, pwd, sid, err := h.getUserPassSd()
	if err != nil {
		return false, err
	}
	h.user = user
	h.pwd = pwd
	h.sid = sid

	id, err := h.getSessionID()
	if err != nil {
		return false, err
	}

	resp, err := h.getQrCode(id)
	if err != nil {
		return false, err
	}
	time.Sleep(time.Second * 2)

	type Code struct {
		XLength    int    `json:"xLenght"`
		YLength    int    `json:"yLenght"`
		PicNum     string `json:"picNum"`
		TimeTemple string `json:"timeTemple"`
	}
	code := &Code{}
	err = json.Unmarshal([]byte(resp), code)
	if err != nil {
		return false, err
	}
	logger.Sugar.Infof("xLength:%d,timeTemple:%s", code.XLength, code.TimeTemple)

	resp, err = h.checkQrCode(id, code.TimeTemple, code.XLength)
	if err != nil {
		return false, err
	}

	resp, err = h.login(h.sessionId, h.user, h.pwd, h.sid+"22")
	if err != nil {
		return false, err
	}

	//{
	//	"success": true,
	//	"flag": "Data",
	//	"msg": "登录成功",
	//	"data": {
	//	"url": "psInfo/psInfo.html",
	//		"result": null,
	//		"name": null,
	//		"regionCode": null,
	//		"sessUser": "4a2c1f4091444292ad7e242c2cc3ef97",
	//		"userLevel": null
	//	}
	//}

	//{
	//	"success": false,
	//	"flag": "Data",
	//	"msg": "登录名无效！",
	//	"data": {
	//	"url": null,
	//		"result": null,
	//		"name": null,
	//		"regionCode": null,
	//		"sessUser": null,
	//		"userLevel": null
	//	}
	//}

	type LoginResult struct {
		Success bool   `json:"success"`
		Msg     string `json:"msg"`
	}
	result := &LoginResult{}
	err = json.Unmarshal([]byte(resp), result)
	if err != nil {
		return false, err
	}

	if result.Success {
		return true, nil
	} else {
		return false, errors.New(result.Msg)
	}
}

// 获取废气排放指示
func (h *HBWebSite) GetPortInfo(factoryId string) (success, data bool, e error) {
	// {"porttypeCode":2,"psid":"39bd36aac0c946b7a55033ec5e3d5f2d"}
	// {
	//    "success": true,
	//    "flag": "Data",
	//    "msg": "成功!",
	//    "data": true
	//}
	jsonData := fmt.Sprintf(`{"porttypeCode":2, "psid":"%s"}`, factoryId)
	resp, err := h.postWitchCookie(UrlPortInfo, jsonData, "application/json")
	if err != nil {
		return false, false, err
	}

	type PortInfo struct {
		Success bool   `json:"success"`
		Msg     string `json:"msg"`
		Data    bool   `json:"data"`
	}
	info := &PortInfo{}
	err = json.Unmarshal([]byte(resp), info)
	if err != nil {
		return false, false, err
	}

	if !info.Success {
		return false, false, errors.New(info.Msg)
	}
	return true, info.Data, nil
}

// 获取工厂所有监测点列表
func (h *HBWebSite) GetFactoryPortList(factoryId string) (*[]PortInfo, error) {
	jsonData := fmt.Sprintf(`{"portTypeCode": 2, "psId": "%s", "startTime": "2019-04-11 00", "endTime": "2019-04-11 14"}`, factoryId)
	resp, err := h.postWitchCookie(UrlGetPortList, jsonData, "application/json")
	if err != nil {
		return nil, err
	}

	type PortListResult struct {
		Success bool        `json:"success"`
		Msg     string      `json:"msg"`
		Data    *[]PortInfo `json:"data"`
	}
	r := &PortListResult{}
	err = json.Unmarshal([]byte(resp), r)
	if err != nil {
		return nil, err
	}
	if r.Success {
		return r.Data, nil
	} else {
		return nil, errors.New(r.Msg)
	}

	//	{
	//		"success": true,
	//		"flag": "ListData",
	//		"msg": "查询成功",
	//		"data": [
	//	{
	//		"color": "#000000",
	//		"psinfoid": "130400000245",
	//		"name": "1号2号焦炉脱硫后排口",
	//		"id": "9",
	//		"dgi_mn": "1304270L0CX029",
	//		"monitorpointtype_code": "2"
	//	},
	//		{
	//			"color": "#000000",
	//			"psinfoid": "130400000245",
	//			"name": "1号焦炉脱硫前排口",
	//			"id": "1",
	//			"dgi_mn": "130427YTJH0001",
	//			"monitorpointtype_code": "5"
	//		},
	//		{
	//			"color": "#000000",
	//			"psinfoid": "130400000245",
	//			"name": "2号焦炉脱硫前排口",
	//			"id": "2",
	//			"dgi_mn": "1304270L0CX028",
	//			"monitorpointtype_code": "5"
	//		},
	//		{
	//			"color": "#000000",
	//			"psinfoid": "130400000245",
	//			"name": "地面站",
	//			"id": "11",
	//			"dgi_mn": "130427YTJH0002",
	//			"monitorpointtype_code": "5"
	//		}
	//]
	//}
}

// 获取监测点历史数据表头
func (h *HBWebSite) GetChatDataHeader(factoryId, portId string) (*[]PortChatHeader, error) {
	jsonData := fmt.Sprintf(`{"colStr":"0", "periodType":"2061", "pollutionType":2, "psId":"%s", "outletNo":"%s","choose":"1"}`, factoryId, portId)
	resp, err := h.postWitchCookie(UrlDataChatHeader, jsonData, "application/json")
	if err != nil {
		return nil, err
	}

	type ChatHeader struct {
		Success bool              `json:"success"`
		Msg     string            `json:"msg"`
		Data    *[]PortChatHeader `json:"data"`
	}
	r := &ChatHeader{}
	err = json.Unmarshal([]byte(resp), r)
	if err != nil {
		return nil, err
	}
	if r.Success {
		return r.Data, nil
	} else {
		return nil, errors.New(r.Msg)
	}
}

// 获取监测点历史数据
func (h *HBWebSite) GetChatData(factoryId, portId string, day time.Time) (*[]PortChatData, error) {
	jsonData := fmt.Sprintf(`{"length":5,"search":{"psId":"%s","pollutionType":"2","outletNo":"%s","colStr":"0","periodType":"2061","fromTime":"%s 00","toTime":"%s 23","header":"","choose":"1","data_source":"1"},"start":"1"
}`, factoryId, portId, day.Format("2006-01-02"), day.Format("2006-01-02"))
	_, err := h.postWitchCookie(UrlDataQueryHb, jsonData, "application/json")
	if err != nil {
		return nil, err
	}
	return nil, nil
}
