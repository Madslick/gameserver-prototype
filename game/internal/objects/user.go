package objects

import (
    "github.com/name5566/leaf/gate"
)
// One struct for one message
// Contains a string member
type User struct {
    Name string
    Points int
    gate.Agent
}
