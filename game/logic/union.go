package logic

import (
	"pomeloServe/core/models/entity"
	"pomeloServe/core/service"
	"pomeloServe/framework/msError"
	"pomeloServe/framework/remote"
	"pomeloServe/game/component/room"
	"pomeloServe/game/models/request"
	"sync"
)

type Union struct {
	sync.RWMutex
	Id       int64
	m        *UnionManager
	RoomList map[string]*room.Room
}

func (u *Union) CreateRoom(service *service.UserService, session *remote.Session, req request.CreateRoomReq, userData *entity.User) *msError.Error {
	//1. 需要创建一个房间 生成一个房间号
	roomId := u.m.CreateRoomId()
	newRoom := room.NewRoom(roomId, req.UnionID, req.GameRule, u)
	u.RoomList[roomId] = newRoom
	return newRoom.UserEntryRoom(session, userData)
}

func (u *Union) DismissRoom(roomId string) {
	u.Lock()
	defer u.Unlock()
	delete(u.RoomList, roomId)
}
func NewUnion(m *UnionManager) *Union {
	return &Union{
		RoomList: make(map[string]*room.Room),
		m:        m,
	}
}
