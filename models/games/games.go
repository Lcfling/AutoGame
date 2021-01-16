package games

import (
	"github.com/Lcfling/AutoGame/models"
	"github.com/astaxie/beego/orm"
)

type Platform struct {
	Id       int64
	Url      string
	Name     string
	Configs  string
	Platform string
}

func (this *Platform) TableName() string {
	return models.TableName("platform")
}
func init() {
	orm.RegisterModel(new(Platform))
}
func ListPlatform() []Platform {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable(models.TableName("platform"))
	cond := orm.NewCondition()
	qs = qs.SetCond(cond)
	var res []Platform
	//qs = qs.OrderBy("-id")
	qs.Limit(100, 0).All(&res)
	return res
}
