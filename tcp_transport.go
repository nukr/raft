package raft

import (
	"fmt"
	"log"
	"net"
)

func dial(peers []string) {
	for _, peer := range peers {
		conn, err := net.Dial("tcp", peer)
		if err != nil {
			log.Fatalln(err)
		}
		go handleConnection(conn)
	}
}

func serve(serverPort int) {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", serverPort))
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("server listening port %d\n", serverPort)
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handleConnection(conn)
	}
}
