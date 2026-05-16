package net

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"google.golang.org/protobuf/proto"

	pb "proto/pb"
)

func NewTcpServerback(port string) {
	// 监听端口
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("监听失败:", err)
		return
	}
	defer listener.Close()

	fmt.Println("服务端启动，监听端口 8080...")

	for {
		// 接受连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("接受连接失败:", err)
			continue
		}
		fmt.Println("conn", conn.RemoteAddr().String())
		go func() {
			newSession := NewTcpSession(conn)
			SessionMgrInstance.AddSession(newSession)
			newSession.Run()
			SessionMgrInstance.DelSession(newSession.UId)
		}()

		// 为每个连接启动一个 goroutine 处理
		//go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Printf("客户端 %s 已连接\n", conn.RemoteAddr())

	reader := bufio.NewReader(conn)

	for {
		// 读取客户端消息（按换行符分隔）
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("客户端 %s 断开连接\n", conn.RemoteAddr())
			return
		}

		// 去除换行符
		message = strings.TrimSpace(message)
		fmt.Printf("收到消息: %s\n", message)

		req := &pb.RegisterRequest{}
		err = proto.Unmarshal([]byte(message), req)
		if err != nil {
			fmt.Println("解析请求失败:", err)
			return
		}
		fmt.Printf("收到消息: %v\n", req)

		// 处理消息
		response := "服务器收到: " + message + "\n"

		// 发送响应
		_, err = conn.Write([]byte(response))
		if err != nil {
			fmt.Println("发送消息失败:", err)
			return
		}
	}
}

type Server struct {
	tcpListener net.Listener
	Handlers    map[pb.MessageId]func(message *SessionPacket)
	//OnSessionPacket func(packet *SessionPacket)
}

func (mm *Server) SendMsg(id uint64, session *TcpSession, message proto.Message) {
	bytes, err := proto.Marshal(message)
	if err != nil {
		return
	}
	rsp := &ProtoMessage{
		ID:   id,
		Data: bytes,
	}
	session.SendMsg(rsp)
}

func (mm *Server) OnSessionPacket(packet *SessionPacket) {
	fmt.Println("msg id ", packet.Msg.ID, " id ", pb.MessageId(packet.Msg.ID), mm.Handlers)
	if handler, ok := mm.Handlers[pb.MessageId(packet.Msg.ID)]; ok {
		handler(packet)
		return
	}
}

func (mm *Server) HandlerRegister() {
	mm.Handlers[pb.MessageId_CSLogin] = mm.CSLogin
	//mm.Handlers[pb.MessageId_CSAddFriend] = mm.UserLogin
}

func (mm *Server) CSLogin(message *SessionPacket) {

	msg := &pb.RegisterRequest{}
	err := proto.Unmarshal(message.Msg.Data, msg)
	if err != nil {
		return
	}
	fmt.Println("[MgrMgr.CreatePlayer]", msg)
	mm.SendMsg(uint64(pb.MessageId_CSLogin), message.Sess, &pb.RegisterRequest{
		Account:       "arick",
		Password:      "666",
		LoginPlatform: 4,
		SmsCode:       "777",
	})

}

func NewTcpServer(address string) *Server {
	resolveTCPAddr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		panic(err)
	}
	tcpListener, err := net.ListenTCP("tcp", resolveTCPAddr)
	if err != nil {
		panic(err)
	}
	fmt.Println("listen tcp:", resolveTCPAddr)
	s := &Server{}
	s.tcpListener = tcpListener
	s.Handlers = make(map[pb.MessageId]func(message *SessionPacket))
	s.HandlerRegister()
	return s

}

func (s *Server) Run() {
	for {
		conn, err := s.tcpListener.Accept()
		if err != nil {
			if _, ok := err.(net.Error); ok {
				fmt.Println(err)
				continue
			}
		}
		fmt.Println("conn ", conn.RemoteAddr().String())
		go func() {
			newSession := NewTcpSession(conn)
			newSession.MessageHandler = s.OnSessionPacket
			SessionMgrInstance.AddSession(newSession)
			newSession.Run()
			SessionMgrInstance.DelSession(newSession.UId)
		}()
	}
}
