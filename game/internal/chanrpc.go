package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/madslick/gameserver-prototype/game/internal/objects"
	"github.com/name5566/leaf/log"
	"github.com/madslick/gameserver-prototype/msg"

	"fmt"
)

//var agents = make(map[gate.Agent]struct{})

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
}

func rpcNewAgent(args []interface{}) {

	user := objects.User{"Test", 0, args[0].(gate.Agent)}
	sessionId := objects.AddUser(user)
	user.Agent.WriteMsg(&msg.Connect{
		SessionId: sessionId,
	})
	log.Debug(fmt.Sprintf("Session Id %v", sessionId))
}

func rpcCloseAgent(args []interface{}) {

}
