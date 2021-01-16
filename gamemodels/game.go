package gamemodels

import (
	"encoding/json"
	"fmt"
	"github.com/Lcfling/AutoGame/models/games"
	"github.com/Lcfling/AutoGame/utils"
	"github.com/idoubi/goz"
	"strconv"
	"strings"
	"time"
)

//var GameInfo
var s = "JSESSIONID=aaaivkaPuu2gfdgWuMJBD; x-session-token=TlkQNftlh1WIHrbouGvkr4udU%2B1Uli%2FqKhaNQUaUWQAS5sZ7GEq5Sw%3D%3D; checkCode=d9d35671-f110-4ddd-a582-918a8a250651"

type PlatformCert struct {
	Cookie string
	Name   string
	Id     string //平台id
}

type Pingguo struct {
	Cookie     string
	checkCode  string
	JSESSIONID string
	Gamelist   map[string]interface{}
	Token      string
	handle     chan int64
	TaskId     int64
	Issue      string
	Config     games.Task
	TaskNums   int64
	PlayCates  map[string]interface{}
	orderNo    map[string]string
}

func (p *Pingguo) GetGameList() (data []map[string]interface{}) {
	cli := goz.NewClient()
	resp, err := cli.Get("https://www.pg8088.com/static/data/gamedatas-v2.json?0.713297844145778", goz.Options{
		Headers: map[string]interface{}{
			"Content-Type": "application/x-www-form-urlencoded",
		},
	})
	body, err := resp.GetBody()

	var res map[string]interface{}
	err = json.Unmarshal([]byte(body), &res)
	fmt.Println(res, err)
	gameMap := (res["gameMap"].(map[string]interface{}))
	p.PlayCates = res["playCates"].(map[string]interface{})
	for _, v := range gameMap {

		/*NewValue:=v.(map[string]interface{})
		CateMap:=playCates[strconv.FormatInt(int64(NewValue["playCateId"].(float64)),10)].(map[string]interface{})
		NewValue[]*/
		data = append(data, v.(map[string]interface{}))
	}
	//fmt.Println(body)
	return
}

//获取游戏玩法
func (p *Pingguo) GetPlayMethd(data map[string]string) []PlayMathds {

	//https://www.pg8088.com/static/data/80-playdatas-v2.json?_t=1610677087687
	gameId := data["gameId"]
	times := time.Now().UnixNano()
	tstr := strconv.FormatInt(times, 10)
	url := "https://www.pg8088.com/static/data/" + gameId + "-playdatas-v2.json?_t=" + tstr
	cli := goz.NewClient()
	resp, err := cli.Get(url, goz.Options{
		Cookies: "JSESSIONID=" + p.JSESSIONID + ";checkCode=" + p.checkCode,

		Headers: map[string]interface{}{
			"accept":       "application/json, text/javascript, */*; q=0.01",
			"Content-Type": "application/x-www-form-urlencoded",
			"user-agent":   userAgent,
		},
	})
	if err != nil {
		fmt.Println("http get methd error:", err.Error())
		return nil
	}
	body, _ := resp.GetBody()
	var jsonRes map[string]map[string]map[string]interface{}
	err = json.Unmarshal(body, &jsonRes)
	if err != nil {
		fmt.Println("json to map error:", err.Error())
		return nil
	}
	//fmt.Println("body map data:",jsonRes)
	//开始整理数据
	var result []PlayMathds
	for _, v := range jsonRes["plays"] {
		if v["odds"].(float64) > 1 && v["odds"].(float64) < 2.5 {
			var play PlayMathds
			cateId := strconv.FormatInt(int64(v["playCateId"].(float64)), 10)
			maps := (p.PlayCates[cateId]).(map[string]interface{})
			cateName := maps["name"].(string)
			play.Name = "[" + cateName + "]" + v["name"].(string)
			play.Betfree = v["odds"].(float64)
			if v["code"] != nil {
				play.Code = v["code"].(string)
			}

			play.GameId = strconv.FormatInt(int64(v["gameId"].(float64)), 10)
			play.MaxM = int64(v["maxMoney"].(float64))
			play.MinM = int64(v["minMoney"].(float64))
			play.Playid = strconv.FormatInt(int64(v["id"].(float64)), 10)

			result = append(result, play)
		}
	}

	return result
}
func (p *Pingguo) DoLogin(user map[string]string) error {
	cli := goz.NewClient()
	Data := map[string]interface{}{
		"account":  user["account"],
		"password": utils.Md5(user["password"]),
		"loginSrc": 0,
		"valiCode": user["code"],
	}
	resp, err := cli.Get("https://www.pg8088.com/api/login.do", goz.Options{
		Cookies: "JSESSIONID=" + p.JSESSIONID + ";checkCode=" + p.checkCode,

		Headers: map[string]interface{}{
			"accept":       "application/json, text/javascript, */*; q=0.01",
			"Content-Type": "application/x-www-form-urlencoded",
			"user-agent":   userAgent,
		},
		Query: Data,
	})
	fmt.Println("err", err)
	body, _ := resp.GetBody()

	//var res map[string]interface{}
	//err = json.Unmarshal([]byte(body), &res)
	fmt.Println("Login", string(body), err)
	return err
}
func (p *Pingguo) GetValidateCode() []byte {
	cli := goz.NewClient()
	p.JSESSIONID = "aaaivkaPuu2gfdgWuMJBD"
	//checkCode="d9d35671-f110-4ddd-a582-918a8a250651"
	Data := map[string]interface{}{
		"t": time.Now().UnixNano(),
	}
	//var cookieJar, _ = browsercookie.Chrome("https://www.pg8088.com/api/getValidateCode.do")
	//fmt.Println("cookieJar:=",cookieJar)
	resp, _ := cli.Get("https://www.pg8088.com/api/getValidateCode.do", goz.Options{
		/*Cookies:map[string]string{
			"JSESSIONID":JSESSIONID,
		},*/
		Cookies: "JSESSIONID=" + p.JSESSIONID + ";",

		Headers: map[string]interface{}{
			"Content-Type": "application/x-www-form-urlencoded",
			"user-agent":   userAgent,
		},
		Query: Data,
	})

	cookie := resp.GetHeaders()
	fmt.Println("cookie:=", cookie)
	header := resp.GetHeader("set-cookie")
	fmt.Println("set-cookie:=", header)
	for _, v := range header {
		//fmt.Println("v=",v,"k=",k)
		arr := strings.Split(v, ";")
		for _, str := range arr {
			fmt.Println("str=", str)
			arrs := strings.Split(str, "=")
			if arrs[0] == "checkCode" && arrs[1] != "" {
				fmt.Println("找到了 code:", arrs[1])
				p.checkCode = strings.Replace(arrs[1], " ", "", -1)
			}
		}
	}
	body, _ := resp.GetBody()
	//var res map[string]interface{}
	//err = json.Unmarshal([]byte(body), &res)
	//fmt.Println("GetValidateCode",body,err)
	return body
}

func (p *Pingguo) Bet(Betinfo map[string]string) map[string]interface{} {
	cli := goz.NewClient()
	Data := map[string]interface{}{
		"gameId":            "80",
		"totalNums":         "1",
		"totalMoney":        "20",
		"betSrc":            "0",
		"turnNum":           "210112506",
		"betBean[0].playId": "8014101",
		"betBean[0].money":  "20",
	}
	resp, err := cli.Post("https://www.pg8088.com/bet/bet.do", goz.Options{
		Cookies: "JSESSIONID=aaaivkaPuu2gfdgWuMJBx; x-session-token=ftw9BYgIuj11NxRKiJa6ICouIKXkVUiJI3PLM72I8Yz7q7kQUgdsoA%3D%3D;",

		Headers: map[string]interface{}{
			"Accept":       "application/json, text/javascript, */*; q=0.01",
			"Content-Type": "application/x-www-form-urlencoded",
			"user-agent":   userAgent,
		},
		FormParams: Data,
	})
	fmt.Println("err", err)
	body, err := resp.GetBody()

	var res map[string]interface{}
	err = json.Unmarshal([]byte(body), &res)
	fmt.Println("Login", string(body), err)
	if res["success"].(bool) {
		p.GetOrder()
	}

	return nil
}

//获取订单号
func (p *Pingguo) GetOrder() {
	cli := goz.NewClient()
	Data := map[string]interface{}{
		"t":       time.Now().UnixNano(),
		"settled": "false",
		"page":    "1",
		"rows":    "15",
	}
	resp, err := cli.Get("https://www.pg8088.com/bet/getBetBills.do", goz.Options{
		/*Cookies:map[string]string{
			"JSESSIONID":JSESSIONID,
		},*/
		Cookies: "JSESSIONID=" + p.JSESSIONID + ";",

		Headers: map[string]interface{}{
			"Content-Type": "application/x-www-form-urlencoded",
			"user-agent":   userAgent,
		},
		Query: Data,
	})

	if err != nil {
		fmt.Println("err", err)
		return
	}

	body, _ := resp.GetBody()

	var res map[string]interface{}
	err = json.Unmarshal([]byte(body), &res)
	if err != nil {
		fmt.Println("json to map err:", err)
		return
	}
	betList := res["data"].([]map[string]interface{})
	for _, v := range betList {
		if v["turnNum"].(string) == p.Issue && strconv.FormatInt(int64(v["playId"].(float64)), 10) == p.Config.Playmethd && strconv.FormatInt(int64(v["gameId"].(float64)), 10) == p.Config.GameId {
			p.orderNo[p.Issue] = v["orderNo"].(string)
			break
		}
	}

	return

}
func (p *Pingguo) GetUserInfo() map[string]string {

	return nil
}

//获取下一期
// {
//	"issue": "210116057",
//	"endtime": "2021-01-16 08:39:50",
//	"nums": null,
//	"lotteryTime": "2021-01-16 08:40:00",
//	"preIssue": "210116056",
//	"preLotteryTime": "2021-01-16 08:38:45",
//	"preNum": "01,10,06,04,03,09,08,02,05,07",
//	"n11": null,
//	"n12": null,
//	"n13": null,
//	"n14": null,
//	"n15": null,
//	"n16": null,
//	"preIsOpen": true,
//	"serverTime": "2021-01-16 08:39:03",
//	"gameId": 80
//}

func (p *Pingguo) GetNextIssue() map[string]interface{} {
	config := p.Config
	cli := goz.NewClient()
	Data := map[string]interface{}{
		"gameId": config.GameId,
		"t":      time.Now().UnixNano(),
	}
	resp, _ := cli.Get("https://www.pg8088.com/lottery/getNextIssue.do", goz.Options{
		/*Cookies:map[string]string{
			"JSESSIONID":JSESSIONID,
		},*/
		Cookies: "JSESSIONID=" + p.JSESSIONID + ";",

		Headers: map[string]interface{}{
			"Content-Type": "application/x-www-form-urlencoded",
			"user-agent":   userAgent,
		},
		Query: Data,
	})
	body, err := resp.GetBody()
	var res map[string]interface{}
	err = json.Unmarshal([]byte(body), &res)
	if err != nil {
		fmt.Println("Login", string(body), err)
		return nil
	}
	return res
}

//开始任务
func (p *Pingguo) Begin() {
	p.process()
	//p.Config:=games.GetTask(p.TaskId)
}

func (p *Pingguo) process() {

}

//结束当前任务
func (p *Pingguo) Stop() {

}
func (p *Pingguo) run() {
	/*data:=p.GetNextIssue()
	//config:=games.GetTask(p.TaskId)
	p.Issue=data["preIssue"].(string)
	Servertime:=utils.Time2unix(data["serverTime"].(string))
	openTime:=utils.Time2unix(data["lotteryTime"].(string))
	waitime:=openTime-Servertime
	var proTime int64
	Now:=time.Now().Unix()
	if waitime>20{
		//
		if p.TaskNums<p.Config.Nums{
			status:=p.Bet()//
			if !status{
				return
			}

		}
	}else{
		return
	}
	proTime=time.Now().Unix()
	s:=Now+waitime-proTime+5
	//休眠等待开奖时间
	time.Sleep(time.Duration(s)*time.Second)
	//开始拿结果
	running:=true
	for running{
		res:=p.GetNextIssue()
		if res["preIssue"].(string)==p.Issue{
			if rsta{
				running=false
				p.TaskNums=0
			}else{
				time.Sleep(time.Duration(1)*time.Second)
			}

		}else{
			time.Sleep(time.Duration(2)*time.Second)

		}

	}*/

}
