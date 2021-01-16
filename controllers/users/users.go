package users

import (
	"fmt"
	"github.com/Lcfling/AutoGame/controllers"
	"github.com/Lcfling/AutoGame/gamemodels"
)

//登录服务器  获取服务许可证
type LoginController struct {
	controllers.IndexController
}

func (this *LoginController) Post() {
	username := this.GetString("username")
	password := this.GetString("password")
	this.SetSession("userLogin", 1)
	this.Data["json"] = map[string]interface{}{"code": 1, "info": "登录成功", "data": username + password}
	this.ServeJSON()
}

//登录游戏服务器 获取游戏资料
type GameLoginController struct {
	controllers.IndexController
}

func (this *GameLoginController) Post() {
	username := this.GetString("username")
	password := this.GetString("password")
	platform := this.GetString("platform")

	code := this.GetString("code")

	fmt.Println("入参：", username, password, code)
	var Loginfo map[string]string
	Loginfo = map[string]string{"account": username, "password": password, "code": code}

	if gamemodels.Sitemap[platform] == nil {
		this.Data["json"] = map[string]interface{}{"code": 0, "info": "平台不存在"}
		this.ServeJSON()
	}
	err := gamemodels.Sitemap[platform].DoLogin(Loginfo)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"code": 0, "info": err.Error(), "data": ""}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 1, "info": "登录成功", "data": ""}
	}

	this.ServeJSON()
}

//登录游戏服务器 获取游戏资料
type GameBetController struct {
	controllers.IndexController
}

func (this *GameBetController) Post() {
	username := this.GetString("username")
	password := this.GetString("password")
	platform := this.GetString("platform")
	code := this.GetString("code")

	fmt.Println("入参：", username, password, code)
	res := gamemodels.Sitemap[platform].Bet(nil)
	this.Data["json"] = map[string]interface{}{"res": res}
	this.ServeJSON()
}

type GetCodeImgController struct {
	controllers.IndexController
}

func (this *GetCodeImgController) Get() {
	platform := this.GetString("platform")
	res := gamemodels.Sitemap[platform].GetValidateCode()
	//this.Data[]
	this.Ctx.Output.ContentType("png")
	this.Ctx.Output.Body(res)
}

type UserInfoController struct {
	controllers.IndexController
}
