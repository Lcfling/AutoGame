package controllers

import (
	//"github.com/virteman/OPMS/initial"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	IsLogin bool
	//UserInfo string
	UserUserId   int64
	UserUsername string
	UserAvatar   string
}

func (this *BaseController) Prepare() {
	this.Ctx.Input.URL()
	userLogin := this.GetSession("userLogin")
	if userLogin == nil {
		this.IsLogin = false
		/*_, ok := this.Ctx.Input.Session("userLogin").(string)
		if !ok && !(this.Ctx.Request.RequestURI == "/login" || strings.Contains(this.Ctx.Request.RequestURI, "/register")) {
			this.Ctx.Redirect(302, "/login")
		}*/
	} else {
		this.IsLogin = true
		/*tmp := strings.Split((this.GetSession("userLogin")).(string), "||")

		userid, _ := strconv.Atoi(tmp[0])
		longid := int64(userid)
		this.Data["LoginUserid"] = longid
		this.Data["LoginUsername"] = tmp[1]
		this.Data["LoginAvatar"] = tmp[2]

		this.UserUserId = longid
		this.UserUsername = tmp[1]
		this.UserAvatar = tmp[2]*/

		//this.Data["PermissionModel"] = this.GetSession("userPermissionModel")
		//this.Data["PermissionModelc"] = this.GetSession("userPermissionModelc")

		//消息
	}
	this.Data["IsLogin"] = this.IsLogin
}
func (this *BaseController) SendMsg(msg string) {
	publish <- newEvent(4, "", msg)
}
