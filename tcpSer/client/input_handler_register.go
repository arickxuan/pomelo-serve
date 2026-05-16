package main

import pb "proto/pb"

func (c *Client) InputHandlerRegister() {
	c.inputHandlers[pb.MessageId_CSLogin.String()] = c.Login
	c.inputHandlers[pb.MessageId_CSAddFriend.String()] = c.AddFriend
	c.inputHandlers[pb.MessageId_CSDelFriend.String()] = c.DelFriend
	c.inputHandlers[pb.MessageId_CSSendChatMsg.String()] = c.SendChatMsg
}

// CSAddFriend 10001
//
