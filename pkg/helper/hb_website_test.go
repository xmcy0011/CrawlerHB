package helper

import (
	"encoding/json"
	"github.com/CrawlerHB/pkg/logger"
	"testing"
	"time"
)

func TestHBWebSite_GetUserPassSd(t *testing.T) {
	logger.InitLogger("log.log", "DEBUG")

	var hb = &HBWebSite{host: "http://110.249.223.71:9030"}
	user, pwd, sid, err := hb.getUserPassSd()
	if err != nil {
		t.Failed()
		return
	}

	logger.Sugar.Infof("user:%s,pwd:%s,sid:%s", user, pwd, sid)
}

func TestHBWebSite_Post(t *testing.T) {
	logger.InitLogger("log.log", "DEBUG")

	var hb = &HBWebSite{host: "http://110.249.223.71:9030"}
	data, err := hb.post("/onlinemonitor/mon/psBaseInfo/systemName", nil, "application/json")
	if err != nil {
		t.Failed()
		return
	}

	t.Logf("%s", data)
}

func TestHBWebSite_GetSessionID(t *testing.T) {
	logger.InitLogger("log.log", "DEBUG")

	var hb = &HBWebSite{host: "http://110.249.223.71:9030"}
	id, err := hb.getSessionID()
	if err != nil {
		t.Failed()
		return
	}

	t.Logf("sessionId:%s", id)
}

func TestHBWebSite_PostWitchCookie(t *testing.T) {
	logger.InitLogger("log.log", "DEBUG")

	var hb = &HBWebSite{host: "http://110.249.223.71:9030"}
	id, err := hb.getSessionID()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("sessionId:%s", id)

	resp, err := hb.getQrCode(id)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("getQrCode %s", resp)
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

	resp, err = hb.checkQrCode(id, code.TimeTemple, code.XLenght)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("checkQrCode %s", resp)
	}
}

func TestHBWebSite_Login(t *testing.T) {
	logger.InitLogger("log.log", "DEBUG")

	var hb = &HBWebSite{}
	resp, err := hb.Login("http://110.249.223.71:9030")
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("Login %s", resp)
	}
}
