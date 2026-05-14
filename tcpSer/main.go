package main

import "pomeloServe/framework/net"

func main() {
	server := net.NewTcpServer(":8080")
	server.Run()
	select {}
}
