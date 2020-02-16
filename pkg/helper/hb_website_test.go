package helper

import (
	"encoding/json"
	"github.com/CrawlerHB/pkg/logger"
	"testing"
	"time"
)

func TestHBWebSite_GetUserPassSd(t *testing.T) {
	logger.InitLogger("log.log", "DEBUG")

	var hb = &HBWebSite{Host: "http://110.249.223.71:9030"}
	user, pwd, sid, err := hb.GetUserPassSd()
	if err != nil {
		t.Failed()
		return
	}

	logger.Sugar.Infof("user:%s,pwd:%s,sid:%s", user, pwd, sid)
}

func TestHBWebSite_Post(t *testing.T) {
	logger.InitLogger("log.log", "DEBUG")

	var hb = &HBWebSite{Host: "http://110.249.223.71:9030"}
	data, err := hb.Post("/onlinemonitor/mon/psBaseInfo/systemName", nil, "application/json")
	if err != nil {
		t.Failed()
		return
	}

	t.Logf("%s", data)
}

func TestHBWebSite_GetSessionID(t *testing.T) {
	logger.InitLogger("log.log", "DEBUG")

	var hb = &HBWebSite{Host: "http://110.249.223.71:9030"}
	id, err := hb.GetSessionID()
	if err != nil {
		t.Failed()
		return
	}

	t.Logf("sessionId:%s", id)
}

func TestHBWebSite_PostWitchCookie(t *testing.T) {
	logger.InitLogger("log.log", "DEBUG")

	var hb = &HBWebSite{Host: "http://110.249.223.71:9030"}
	id, err := hb.GetSessionID()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("sessionId:%s", id)

	resp, err := hb.GetQrCode(id)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("GetQrCode %s", resp)
	time.Sleep(time.Second * 2)

	type Code struct {
		XLenght    int    `json:"xLenght"`
		YLenght    int    `json:"yLenght"`
		PicNum     string `json:"picNum"`
		TimeTemple string `json:"timeTemple"`
	}
	code := &Code{}
	err = json.Unmarshal([]byte(resp), code)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("xLenght:%d,timeTemple:%s", code.XLenght, code.TimeTemple)

	resp, err = hb.CheckQrCode(id, code.TimeTemple, code.XLenght)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("CheckQrCode %s", resp)
	}
}

func TestHBWebSite_Login(t *testing.T) {
	logger.InitLogger("log.log", "DEBUG")

	var hb = &HBWebSite{Host: "http://110.249.223.71:9030"}
	user, pwd, sid, err := hb.GetUserPassSd()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("GetUserPassSd ,user=%s,pwd=%s,sid=%s", user, pwd, sid)

	id, err := hb.GetSessionID()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("sessionId:%s", id)

	resp, err := hb.GetQrCode(id)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("GetQrCode %s", resp)
	time.Sleep(time.Second * 2)

	type Code struct {
		XLenght    int    `json:"xLenght"`
		YLenght    int    `json:"yLenght"`
		PicNum     string `json:"picNum"`
		TimeTemple string `json:"timeTemple"`
	}
	code := &Code{}
	err = json.Unmarshal([]byte(resp), code)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("xLenght:%d,timeTemple:%s", code.XLenght, code.TimeTemple)

	resp, err = hb.CheckQrCode(id, code.TimeTemple, code.XLenght)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("CheckQrCode %s", resp)
	}

	resp, err = hb.Login(id, user, pwd, sid)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("Login %s", resp)
	}
}
