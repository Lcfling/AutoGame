package gamemodels

var userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36"
var Sitemap SitHandel

type SitHandel map[string]GameCrl

func init() {
	Sitemap = make(map[string]GameCrl)
	Sitemap["pingguo"] = &Pingguo{}

}

/*type GameCrl interface {
	Dologin()
	GetGamelist()		//获取游戏列表
	GetPlayMethd()		//获取玩法列表
	GetUserInfo()		//获取用户信息
	GetValidateCode()	//获取验证码
	Bet()
}*/

type GameCrl interface {
	DoLogin(map[string]string) error
	GetGameList() []map[string]interface{}            //获取游戏列表
	GetPlayMethd(data map[string]string) []PlayMathds //获取玩法列表
	GetUserInfo() map[string]string                   //获取用户信息
	GetValidateCode() []byte                          //获取验证码
	Bet(map[string]string) map[string]interface{}
	Begin()
	Stop()
}
type PlayMathds struct {
	GameId  string
	Betfree float64
	Playid  string
	Name    string
	Code    string
	MaxM    int64
	MinM    int64
}
