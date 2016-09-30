package raft

import "time"

// Config ...
type Config struct {
	ElectionTimeout time.Duration
}
