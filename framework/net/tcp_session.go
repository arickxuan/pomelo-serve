package net

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"strings"
	"time"
)

type TcpSession struct {
	UId            int64
	Conn           net.Conn
	IsClose        bool
	packer         IPacker
	WriteCh        chan *ProtoMessage
	IsPlayerOnline bool
	MessageHandler func(packet *SessionPacket)
	//
}

func NewTcpSession(conn net.Conn) *TcpSession {
	return &TcpSession{Conn: conn, packer: &NormalPacker{ByteOrder: binary.BigEndian}, WriteCh: make(chan *ProtoMessage, 1)}
}

func (s *TcpSession) Run() {
	go s.Read()
	go s.Write()

}

func (s *TcpSession) Read() {
	for {
		//fmt.Println("start read")
		err := s.Conn.SetReadDeadline(time.Now().Add(time.Second))
		if err != nil {
			fmt.Println(err)
			return
		}
		message, err := s.packer.Unpack(s.Conn)
		if err != nil {
			switch {
			case errors.Is(err, io.EOF):
				// 正常关闭
				fmt.Println("Connection closed by peer gracefully")
				return

			case strings.Contains(err.Error(), "connection reset"):
				// 强制重置
				fmt.Println("Connection reset by peer (client crashed or force closed)")
				return

			case strings.Contains(err.Error(), "timeout"):
				// 超时
				fmt.Println("Connection timeout")
				return

			case strings.Contains(err.Error(), "refused"):
				// 连接拒绝
				fmt.Println("Connection refused")
				return

			default:
				if netErr, ok := err.(net.Error); ok {
					fmt.Printf("Other network error: %v\n", netErr)
					return
				}
				// 协议错误，可以继续
				fmt.Printf("Protocol error: %v\n", err)
				continue
			}
		}
		//fmt.Println("receive message:", string(message.Data))
		s.MessageHandler(&SessionPacket{
			Msg:  message,
			Sess: s,
		})
		//s.WriteCh <- &ProtoMessage{
		//	ID:   555,
		//	Data: []byte("hi"),
		//}
	}

}

func (s *TcpSession) Write() {
	for {
		select {
		case resp := <-s.WriteCh:
			s.send(resp)
		}
	}
}

func (s *TcpSession) SendMsg(msg *ProtoMessage) {
	s.WriteCh <- msg
}

func (s *TcpSession) send(message *ProtoMessage) {
	err := s.Conn.SetWriteDeadline(time.Now().Add(time.Second))
	if err != nil {
		fmt.Println(err)
		return
	}
	bytes, err := s.packer.Pack(message)
	if err != nil {
		fmt.Println(err)
		return
	}
	s.Conn.Write(bytes)

}
