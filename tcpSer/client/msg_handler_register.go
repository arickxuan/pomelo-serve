package main

import "pomeloServe/proto/pd"

func (c *Client) MessageHandlerRegister() {
	c.messageHandlers[pd.MessageId_SCLogin] = c.OnLoginRsp
	//c.messageHandlers[pb.MessageId_SCAddFriend] = c.OnAddFriendRsp
	//c.messageHandlers[pb.MessageId_SCDelFriend] = c.OnDelFriendRsp
	//c.messageHandlers[pb.MessageId_SCSendChatMsg] = c.OnSendChatMsgRsp

}
