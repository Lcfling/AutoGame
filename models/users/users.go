package users

type Users struct {
	Id       int64         `orm:"pk;column(userid);"`
	Token    string
	Pid      int64
	Areaid   int64
	Pareaid  int64
	Username string
	Password string
	Avatar   string
	Status   int
}


