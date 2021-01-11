package gamemodels

import (
	"encoding/json"
	"fmt"
	"github.com/idoubi/goz"
)

//var GameInfo

type PlatformCert struct {
	Cookie string
	Name 	string
	Id 		string //平台id
}

func GetGameList (url string) map[string]interface{}{
	cli:= goz.NewClient()
	resp, err := cli.Get("https://www.pg8088.com/static/data/gamedatas-v2.json?0.713297844145778", goz.Options{
		Headers: map[string]interface{}{
			"Content-Type": "application/x-www-form-urlencoded",
		},
		FormParams: map[string]interface{}{
			"key1": "value1",
			"key2": []string{"value21", "value22"},
			"key3": "333",
		},
	})
	body,err:=resp.GetBody()

	var res map[string]interface{}
	err = json.Unmarshal([]byte(body), &res)
	fmt.Println(res,err)
	//fmt.Println(body)
	return res
}
//获取游戏玩法
func GetPlayMethd() map[string]interface{}{

	return nil
}