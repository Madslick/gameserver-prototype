package objects

import (
    "fmt"
    "github.com/name5566/leaf/log"
)
var Sessions []*Session
var MAX_SESSIONS = 4

// One struct for one message
// Contains a string member
type Session struct {
    Id int
    Players []User
}

func AddUser(player User) int {
    for _, session := range Sessions {
        if len(session.Players) < MAX_SESSIONS {
            session.Players = append(session.Players, player)
            log.Debug(fmt.Sprintf("Added User %v", player))
            return session.Id
        }
    }
    session := Session{ len(Sessions), make([]User, 0) }
    Sessions = append(Sessions, &session)
    session.Players = append(session.Players, player)
    log.Debug(fmt.Sprintf("Added User %v", player))

    return session.Id
}
