package common

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
	"gopkg.in/redis.v5"

	"io"
	"time"
	"net/http"
	"io/ioutil"
)

type BeeControllers struct {
	beego.Controller
	db *gorm.DB
	r *redis.Client
}

func NewBeeController(db *gorm.DB, r *redis.Client) *BeeControllers {
	return &BeeControllers{
		db:db,
		r:r,
	}
}

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

func ByteString(p []byte) string {
	for i := 0; i < len(p); i++ {
		if p[i] == 0 {
			return string(p[0:i])
		}
	}
	return string(p)
}