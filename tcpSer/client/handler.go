package main

import (
	"fmt"
	"pomeloServe/framework/net"
	"pomeloServe/proto/pd"

	"google.golang.org/protobuf/proto"
)

type MessageHandler func(packet *net.ClientPacket)

type InputHandler func(param *InputParam)

func (c *Client) Login(param *InputParam) {
	fmt.Println("Login input Handler print")
	fmt.Println(param.Command)
	fmt.Println(param.Param)
	req := &pd.RegisterRequest{
		Account:       "message",
		Password:      "123456",
		LoginPlatform: 3,
		SmsCode:       "432",
	}
	msg, err := proto.Marshal(req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(msg))
	pkg := &net.ProtoMessage{
		Data: msg,
		ID:   uint64(pd.MessageId_CSLogin),
	}

	bytes, err := c.packer.Pack(pkg)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bytes))
	c.cli.ChMsg <- pkg

}

func (c *Client) OnLoginRsp(packet *net.ClientPacket) {

}

func (c *Client) AddFriend(param *InputParam) {

}

func (c *Client) OnAddFriendRsp(packet *net.ClientPacket) {

}

func (c *Client) DelFriend(param *InputParam) {

}

func (c *Client) OnDelFriendRsp(packet *net.ClientPacket) {

}

func (c *Client) SendChatMsg(param *InputParam) {

}

func (c *Client) OnSendChatMsgRsp(packet *net.ClientPacket) {

}
