package games

import (
	"fmt"
	"github.com/Lcfling/AutoGame/controllers"
	"github.com/Lcfling/AutoGame/gamemodels"
	"github.com/Lcfling/AutoGame/models/games"
)

//游戏选择
type GamesChooseController struct {
	controllers.BaseController
}

func (this *GamesChooseController) Get() {
	gameId := this.GetString("gameid")
	platform := this.GetString("platform")
	fmt.Println("gameid=", gameId)
	data := make(map[string]string)
	data["gameId"] = gameId
	res := gamemodels.Sitemap[platform].GetPlayMethd(data)
	this.Data["json"] = map[string]interface{}{"code": 1, "info": "success", "data": res}
	this.ServeJSON()

}

//获取游戏列表
type GamesListController struct {
	controllers.IndexController
}

func (this *GamesListController) Get() {
	//this.GetString("")
	platform := this.GetString("platform")
	res := gamemodels.Sitemap[platform].GetGameList()
	this.Data["json"] = map[string]interface{}{"code": 1, "info": "success", "data": res}
	this.ServeJSON()
}

//选择游戏平台
type GamePlatformController struct {
	controllers.BaseController
}

func (this *GamePlatformController) Get() {
	//列出所有游戏平台

	res := games.ListPlatform()
	this.Data["json"] = map[string]interface{}{"code": 1, "info": "success", "data": res}
	this.ServeJSON()
}
