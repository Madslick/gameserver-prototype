package internal

import (
    "github.com/name5566/leaf/log"
    //"github.com/name5566/leaf/gate"
    "reflect"
    "github.com/madslick/gameserver-prototype/msg"
    "github.com/madslick/gameserver-prototype/game/internal/objects"

)

func init() {
    // Register the handler of `Hello` message to `game` module handleHello
    handler(&msg.Hello{}, handleHello)
}

func handler(m interface{}, h interface{}) {
    skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleHello(args []interface{}) {

    // Send "Hello"
    m := args[0].(*msg.Hello)
    // The receiver
    //a := args[1].(gate.Agent)

    // The content of the message
    log.Debug("hello %v", m.Name)

    // Reply with a `Hello`
    handleBroadcast(&msg.Hello{
        Name: "client",
    })
}

func handleBroadcast(cmd interface{}){
	log.Debug("%v",len(objects.Sessions))
    for _,obj:=range objects.Sessions {
        log.Debug("%v",len(obj.Players))
    }
	//for a:=range &objects.Sessions{
	//	log.Debug("%v\n",a.RemoteAddr().String())
	//	a.WriteMsg(cmd)
	//}


}
