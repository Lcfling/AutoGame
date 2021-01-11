package users

import "github.com/Lcfling/AutoGame/controllers"


//登录服务器  获取服务许可证
type LoginController struct {
	controllers.IndexController
}


//登录游戏服务器 获取游戏资料
type GameLoginController struct {
	controllers.IndexController
}
