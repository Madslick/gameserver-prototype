package gate

import (
	"github.com/madslick/gameserver-prototype/game"
	"github.com/madslick/gameserver-prototype/msg"
)

func init() {
	msg.Processor.SetRouter(&msg.Hello{}, game.ChanRPC)
}
