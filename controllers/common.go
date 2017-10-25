package controllers

import (
	"github.com/astaxie/beego/logs"

	"io"
	"time"
	"net/http"
	"io/ioutil"
)

func DoReq(method, urlStr string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, urlStr, body)
	if err != nil {
		logs.Error("new req : ", err)
		return nil, err
	}
	client := http.Client{Timeout:time.Second * 30}
	resp, err := client.Do(req)
	if err != nil {
		logs.Error("client do : ", err)
		return nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error("read resp body : ", err)
		return nil, err
	}
	return data, nil
}