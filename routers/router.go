package routers

import "C"
import (
	"github.com/Lcfling/AutoGame/controllers"
	"github.com/Lcfling/AutoGame/controllers/games"
	"github.com/Lcfling/AutoGame/controllers/users"
	"github.com/astaxie/beego"
)

func init() {

	//beego.Router("/", &users.MainController{})
	beego.Router("/", &games.GamesListController{})
	beego.Router("/login", &users.LoginController{})
	beego.Router("/getCode", &users.GetCodeImgController{})
	beego.Router("/game/login", &users.GameLoginController{})

	beego.Router("/game/bet", &users.GameBetController{})
	beego.Router("/game/platform", &games.GamePlatformController{})
	beego.Router("/game/gamelist", &games.GamesListController{})
	beego.Router("/game/playlist", &games.GamesChooseController{})

	//websocket
	beego.Router("/ws", &controllers.WebSocketController{})
	beego.Router("/ws/join", &controllers.WebSocketController{}, "get:Join")

	//用户

	//文件管理

}
