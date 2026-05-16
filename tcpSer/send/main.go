package main

import (
	"encoding/binary"
	"fmt"
	gonet "net"
	"os"

	"framework/net"
	"proto/pb"

	"google.golang.org/protobuf/proto"
)

func main() {
	// 连接服务器
	conn, err := gonet.Dial("tcp", ":8080")
	if err != nil {
		fmt.Printf("连接失败: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	// 要发送的消息
	message := "Hello, TCP Server!"

	req := &pb.RegisterRequest{
		Account:       message,
		Password:      "123456",
		LoginPlatform: 3,
		SmsCode:       "432",
	}
	msg, err := proto.Marshal(req)
	if err != nil {
		fmt.Println(err)
	}
	packer := &net.NormalPacker{
		ByteOrder: binary.BigEndian,
	}

	pkg := &net.ProtoMessage{
		Data: msg,
		ID:   uint64(pb.MessageId_CSLogin),
	}

	bytes, err := packer.Pack(pkg)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 发送消息
	_, err = conn.Write(bytes)
	if err != nil {
		fmt.Printf("发送失败: %v\n", err)
		return
	}

	// 接收响应
	rec, err := packer.Unpack(conn)
	if err != nil {
		fmt.Println(err)
		return
	}
	if rec != nil {
		fmt.Printf("收到响应: %v\n", rec)
	}

}
