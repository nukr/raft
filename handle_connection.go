package raft

import (
	"fmt"
	"log"
	"net"
	"time"
)

func handleReceive(c net.Conn, connectClose chan bool) {
	for {
		b := make([]byte, 50, 50)
		_, err := c.Read(b)
		if err != nil {
			fmt.Println(err)
		}
		switch int8(b[0]) {
		case 49:
			log.Printf("received request vote %s\n", c.RemoteAddr().String())
			c.Write([]byte("1"))
		case 50:
			log.Printf("received heartbeat from %s\n", c.RemoteAddr().String())
			c.Write([]byte("2"))
		case 0:
			connectClose <- true
		case 10:
		case 13:
		default:
			fmt.Println("unexpected action", b[0])
		}
		time.Sleep(20 * time.Millisecond)
	}
}

func handleSend(c net.Conn) {
}

func handleConnection(c net.Conn) {
	defer c.Close()
	defer fmt.Println("connection closed")
	fmt.Printf("new connection %s => %s\n", c.RemoteAddr().String(), c.LocalAddr().String())
	connectClose := make(chan bool)
	go handleReceive(c, connectClose)
	go handleSend(c)
	<-connectClose
}
