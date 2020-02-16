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
)

type HBWebSite struct {
	Host string

	sessionId string
}

// 发送GET请求
// url：         请求地址，不需要host部分
// response：    请求返回的内容
func (h *HBWebSite) Get(url string) (string, error) {
	// 超时时间：10秒
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(h.Host + ":" + url)
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
func (h *HBWebSite) Post(url string, data interface{}, contentType string) (string, error) {
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

	resp, err := client.Post(h.Host+":"+url, contentType, bytes.NewBuffer(jsonStr))
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
func (h *HBWebSite) PostWitchCookie(url string, data string, contentType string) (string, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("POST", h.Host+":"+url, strings.NewReader(data))
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
func (h *HBWebSite) GetUserPassSd() (user, pwd, sid string, error error) {
	rsp, err := http.Get(h.Host + ":" + UrlGetUserPassSd)
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
	return user, pwd, sid, nil
}

// 2-获取会话ID
func (h *HBWebSite) GetSessionID() (sessionId string, error error) {
	client := &http.Client{Timeout: 10 * time.Second}
	jsonStr := []byte("")
	rsp, err := client.Post(h.Host+":"+UrlGetSessionID, "application/json", bytes.NewBuffer(jsonStr))
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

	return sessionId, nil
}

// 3-获取验证码
func (h *HBWebSite) GetQrCode(sessionId string) (string, error) {
	h.sessionId = sessionId
	return h.PostWitchCookie(UrlGetQrCodeHeight, `{"storagePosition": "database"}`, "application/json")
}

// 4-验证码校验
func (h *HBWebSite) CheckQrCode(sessionId, timeTemple string, xLenght int) (string, error) {
	h.sessionId = sessionId
	data := fmt.Sprintf(`{"timeTemple": "%s", "xLenght": %d, "storagePosition": "database"}`, timeTemple, xLenght)
	logger.Sugar.Info(data)
	return h.PostWitchCookie(UrlSetQrCodeHeight, data, "application/json")
}

// 5-登录
func (h *HBWebSite) Login(sessionId, user, pwd, sid string) (string, error) {
	jsonData := fmt.Sprintf(`{"suLoginid": "%s","suPasswd": "%s","sdId": "%s"}`, user, pwd, sid)
	h.sessionId = sessionId
	return h.PostWitchCookie(UrlLogin, jsonData, "application/json")
}
