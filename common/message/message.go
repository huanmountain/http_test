package message

const (
	LoginMesType = "LoginMes"
	LoginResMesType = "LoginResMes"

)

type Message struct {
	//大写的原因需要序列化???
	Type string `json:"type"`	//消息类型
	Date string	 `json:"date"`	//消息内容

}

type LoginMes struct {
	UserId int		`json:"user_id"`//用户id
	UserPwd string	`json:"user_pwd"`//用户密码
	UserName string `json:"user_name"`//用户名
}

type LoginResMes struct {
	Code int      `json:"code"`  //返回状态码
	Error string  `json:"error"`  //返回错误信息
}