package net

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

type Client struct {
	Address   string
	ChMsg     chan *ProtoMessage
	OnMessage func(message *ClientPacket)
	conn      net.Conn // 无用
	packer    IPacker  //无用
}

func NewClient(address string) *Client {
	return &Client{
		Address: address,
		packer: &NormalPacker{
			ByteOrder: binary.BigEndian,
		},
		ChMsg: make(chan *ProtoMessage, 1),
	}
}

func (c *Client) Run() {
	var err error
	c.conn, err = net.Dial("tcp", c.Address)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("connected to", c.Address)
	go c.Read(c.conn)
	go c.Write(c.conn)

}

func (c *Client) Write(conn net.Conn) {
	tick := time.NewTicker(time.Second)
	for {
		select {
		case <-tick.C:
			c.ChMsg <- &ProtoMessage{
				ID:   111,
				Data: []byte("hello world "),
			}
		case msg := <-c.ChMsg:
			c.Send(conn, msg)
		}
	}
}

func (c *Client) Send(conn net.Conn, message *ProtoMessage) {
	pack, err := c.packer.Pack(message)
	if err != nil {
		fmt.Println(err)
		return
	}
	conn.Write(pack)
}

func (c *Client) Read(conn net.Conn) {
	for {
		message, err := c.packer.Unpack(conn)
		if err != nil {
			fmt.Println(err)
			continue
		}
		c.OnMessage(&ClientPacket{
			Msg:  message,
			Conn: conn,
		})
		fmt.Println("read msg:", string(message.Data))
	}
}
