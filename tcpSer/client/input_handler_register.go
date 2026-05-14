package main

import "pomeloServe/proto/pd"

func (c *Client) InputHandlerRegister() {
	c.inputHandlers[pd.MessageId_CSLogin.String()] = c.Login
	c.inputHandlers[pd.MessageId_CSAddFriend.String()] = c.AddFriend
	c.inputHandlers[pd.MessageId_CSDelFriend.String()] = c.DelFriend
	c.inputHandlers[pd.MessageId_CSSendChatMsg.String()] = c.SendChatMsg
}

// CSAddFriend 10001
//
