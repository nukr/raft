package raft

import (
	"fmt"
	"time"
)

// Raft ...
type Raft struct {
	config                 Config
	state                  State
	peers                  []string
	electionTimeoutElapsed int
}

// Run ...
func (r *Raft) Run() {
	fmt.Println(r.state)
	// fmt.Println(r.config.electionTimeout)
	// r.countdown = int(r.config.electionTimeout)
	// go r.ticker()
	// go r.resetCountdown()
}

// NewRaft ...
func NewRaft(config Config, peers []string) *Raft {
	return &Raft{
		config: config,
		state:  0,
		peers:  peers,
	}
}

func (r *Raft) ticker() {
	ticker := time.NewTicker(time.Millisecond)
	for {
		r.electionTimeoutElapsed += int(time.Millisecond)
		if r.electionTimeoutElapsed <= 0 {
			ticker.Stop()
			r.state.setState(1)
		}
	}
}

func (r *Raft) resetTicker() {
	for {
		fmt.Println("sleep for 20ms", r.countdown)
		time.Sleep(20 * time.Millisecond)
		fmt.Println("reset countdown", r.countdown)
		r.countdown = int(r.config.ElectionTimeout)
	}
}
