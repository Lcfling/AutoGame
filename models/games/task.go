package games

import (
	"github.com/Lcfling/AutoGame/models"
	"github.com/astaxie/beego/orm"
)

type Task struct {
	Id          int64
	Pid         int64
	Platform    string
	Playmethd   string
	GameId      string
	Money       int64
	Add         float64
	Nums        int64
	Creatime    int64
	Lastbettime int64
}

func (this *Task) TableName() string {
	return models.TableName("task")
}
func init() {
	orm.RegisterModel(new(Task))
}

func AddTask(task Task) (id int64, err error) {
	return
}
func UpdateLastbetTime(id int64) error {
	return nil
}
func GetTask(id int64) Task {

	return Task{}
}
