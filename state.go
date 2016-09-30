package raft

import (
	"sync"
)

// State is Follower, Candidate, Leader, Shutdown
type State uint32

const (
	// Follower ...
	Follower State = iota
	// Candidate ...
	Candidate
	// Leader ...
	Leader
	// Shutdown ...
	Shutdown
)

func (s State) String() string {
	switch s {
	case Follower:
		return "Follower"
	case Candidate:
		return "Candidate"
	case Leader:
		return "Leader"
	case Shutdown:
		return "Shutdown"
	default:
		return "Unknown"
	}
}

type state struct {
	state State
	mu    sync.Mutex
}

func (n *state) getState() State {
	n.mu.Lock()
	defer n.mu.Unlock()
	return n.state
}

func (n *state) setState(s State) {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.state = s
}
