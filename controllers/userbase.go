package controllers

import (
	//"github.com/virteman/OPMS/initial"

	"github.com/astaxie/beego"
)

type UserBaseController struct {
	beego.Controller
	IsLogin bool
	//UserInfo string
	UserUserId   int64
	UserUsername string
	UserAvatar   string
}

func (this *UserBaseController) Prepare() {
	token := this.Ctx.Request.Header.Get("token")
	useridstr := this.Ctx.Request.Header.Get("userid")
	if token == "" || useridstr == "" {
		this.Data["json"] = map[string]interface{}{"code": 2, "message": "登录效验失败-效验数据为空", "data": nil}
		this.ServeJSON()
	}



}
func (this *UserBaseController) SendMsg(msg string) {
	publish <- newEvent(4, "", msg)
}
