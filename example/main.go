package main

import (
	"flag"
	"time"

	"github.com/nukr/raft"
)

var (
	peersFlag  peers
	serverPort *int
)

func init() {
	serverPort = flag.Int("port", 3333, "server port")
	flag.Var(&peersFlag, "join", "eg. 1.2.3.4:5678")
}

func main() {
	flag.Parse()
	config := raft.Config{
		ElectionTimeout: 150 * time.Millisecond,
	}
	r := raft.NewRaft(config, peersFlag)

	r.Run()

	time.Sleep(30 * time.Second)
}
