package raft

import (
	"fmt"
	"os"
)

func readKeyin() {
	keyin := make(chan []byte)
	go func() {
		for {
			data := make([]byte, 1)
			os.Stdin.Read(data)
			keyin <- data
		}
	}()
	for {
		select {
		case key := <-keyin:
			if key[0] != 10 {
				fmt.Println(key)
			}
		}
	}
}
