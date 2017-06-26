package msg

import (
	//"github.com/name5566/leaf/network"
	"github.com/name5566/leaf/network/json"
)

//var Processor network.Processor
var Processor = json.NewProcessor()

//var SignUp_processor = json.NewProcessor()

func init() {

	///注册json消息
	Processor.Register(&Ok{})
	Processor.Register(&SignUp{})
	Processor.Register(&SignIn{})
	Processor.Register(&State{})
	Processor.Register(&Up{})
	Processor.Register(&Right{})
	Processor.Register(&Left{})
	Processor.Register(&Down{})
	Processor.Register(&Command{})
	Processor.Register(&UpLoad{})
}

///结构体定义了一个JSON消息格式

///测试消息结构
type Ok struct {
	Name string
}





///注册消息结构
type SignUp struct {

	Name string `json:"name"`
	Password string `json:"password"`
}

///登录消息结构
type SignIn struct {
	Name string `json:"name"`
	Password string `json:"password"`
}

///用来同步用户的个人资料
type UpLoad struct {
	ID int `json:"id"`
	Data UserData `json:"data"`
}

///状态消息(向客户端发送)
type State struct {
	Name string `json:"name"`
}

///测试的向前开车消息
type Up struct {
	Direction float32
}

///向左转
type Left struct {
	Direction float32
}

///向右转
type Right struct {
	Direction float32
}

type Down struct {
	Direction float32
}


///定义了具体的命令
const (
	UpCom = iota
	DownCom
	LeftCom
	RightCom
)

type Command struct {
	CarID int `json:"car_id"`
	ID int `json:"id"`
	Cmd int `json:"cmd"`
	Val float32 `json:"val"`
}