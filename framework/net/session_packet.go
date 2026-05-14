package net

// import "pomeloServe/framework/net"

type SessionPacket struct {
	Msg  *ProtoMessage
	Sess *TcpSession
}

type ProtoMessage struct {
	ID   uint64
	Data []byte
}
