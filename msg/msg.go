package msg

import (
    "github.com/name5566/leaf/network/json"
)

// Create a JSON Processor（or protobuf if you like）
var Processor = json.NewProcessor()

func init() {
    // Register message Hello
    Processor.Register(&Hello{})
}

// One struct for one message
// Contains a string member
type Hello struct {
    Name string
}

type Connect struct {
    SessionId int
}
