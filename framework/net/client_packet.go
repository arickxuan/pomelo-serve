package net

import gonet "net"

type ClientPacket struct {
	Msg  *ProtoMessage
	Conn gonet.Conn
}
