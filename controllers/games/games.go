package games

import (
	"github.com/Lcfling/AutoGame/controllers"
	."github.com/Lcfling/AutoGame/gamemodels"
)


//游戏选择
type GamesChooseController struct {
	controllers.BaseController
}

//获取游戏列表
type GamesListController struct {
	controllers.IndexController
}

func (this *GamesListController)Get(){
	//this.GetString("")
	res:=GetGameList("")
	this.Data["json"]=res
	this.ServeJSON()
}

//选择游戏平台
type GamePlatformController struct {
	controllers.BaseController
}

func (this *GamePlatformController) Get() {
	//列出所有游戏平台
}