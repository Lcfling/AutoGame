package routers

import (
	"github.com/Lcfling/AutoGame/controllers"
	"github.com/Lcfling/AutoGame/controllers/games"
	"github.com/astaxie/beego"
)

func init() {
	//beego.Router("/", &users.MainController{})
	beego.Router("/", &games.GamesListController{})


	//websocket
	beego.Router("/ws", &controllers.WebSocketController{})
	beego.Router("/ws/join", &controllers.WebSocketController{}, "get:Join")

	//用户


	//文件管理


}
