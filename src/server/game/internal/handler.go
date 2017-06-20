package internal

import (
	"reflect"
	"server/msg"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"

)

///一切就绪,在game模块中处理Ok消息

func init() {
	//向当前模块(game 模块)注册Ok 消息的消息处理函数 handleOk
	handler(&msg.Ok{},handleOk)
	handler(&msg.Up{},handleUp)
}

func handler(m interface{},h interface{})  {

	skeleton.RegisterChanRPC(reflect.TypeOf(m),h)
}

///Ok 的消息处理函数
func handleOk(args []interface{}) {
	///收到的 Ok 消息
	m := args[0].(*msg.Ok)

	//log.Debug("%v",m)
	///消息的发送者
	a := args[1].(gate.Agent)

	log.Debug("Ok %v", m.Name)

	///给发送者回应一个Ok消息
	a.WriteMsg(&msg.Ok{
		Name: "client",
	})
}

func handleUp(args []interface{})  {
	///收到的 up消息
	 m:=args[0].(*msg.Up)

	//获取网关,即客户端地址,消息的发送着

	a:=args[1].(gate.Agent)


	///此处可以处理相应得逻辑

	///将命令消息广播转发到其他的分组,设置相应的速度
	log.Debug("%v %v",m.Direction,a)

	///向客户端发送,确认接收的消息

	handleBroadcast()

}

///广播
func handleBroadcast(){
	log.Debug("%v",len(agents))
	for a:=range agents{
		log.Debug("%v\n",a.RemoteAddr().String())
		a.WriteMsg(&msg.State{Name:"broadcast"})
	}
}