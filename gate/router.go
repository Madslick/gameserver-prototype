package gate

import (
    "github.com/name5566/leafserver/game"
    "github.com/name5566/leafserver/msg"
)


func init() {
    msg.Processor.SetRouter(&msg.Hello{}, game.ChanRPC)
}
