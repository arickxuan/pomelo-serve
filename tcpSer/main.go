package main

import "framework/net"

func main() {
	server := net.NewTcpServer(":8080")
	server.Run()
	select {}
}
