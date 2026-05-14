package base

import (
	"pomeloServe/framework/remote"
	"pomeloServe/game/component/proto"
)

type RoomFrame interface {
	GetUsers() map[string]*proto.RoomUser
	GetId() string
	EndGame(session *remote.Session)
	UserReady(uid string, session *remote.Session)
}
