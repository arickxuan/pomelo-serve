package main

import (
	"encoding/binary"
	"fmt"
	"framework/net"

	pb "proto/pb"
)

type Client struct {
	cli             *net.Client
	inputHandlers   map[string]InputHandler
	messageHandlers map[pb.MessageId]MessageHandler
	console         *ClientConsole
	chInput         chan *InputParam
	packer          net.IPacker
}

func NewClient() *Client {
	c := &Client{
		cli:             net.NewClient(":8080"),
		inputHandlers:   map[string]InputHandler{},
		messageHandlers: map[pb.MessageId]MessageHandler{},
		console:         NewClientConsole(),
		packer:          &net.NormalPacker{ByteOrder: binary.BigEndian},
	}
	c.cli.OnMessage = c.OnMessage
	c.cli.ChMsg = make(chan *net.ProtoMessage, 1)
	c.chInput = make(chan *InputParam, 1)
	c.console.chInput = c.chInput
	return c
}

func (c *Client) Run() {
	go func() {
		for {
			select {
			case input := <-c.chInput:
				fmt.Printf("cmd:%s,param:%v  <<<\t \n", input.Command, input.Param, c.inputHandlers)
				inputHandler := c.inputHandlers[input.Command]
				if inputHandler != nil {
					inputHandler(input)
				}
			}
		}
	}()
	go c.console.Run()
	go c.cli.Run()
}

func (c *Client) OnMessage(packet *net.ClientPacket) {
	if handler, ok := c.messageHandlers[pb.MessageId(packet.Msg.ID)]; ok {
		handler(packet)
	}
}
