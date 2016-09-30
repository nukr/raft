package raft

import (
	"errors"
	"math/rand"
	"time"
)

func random(n, m int) (int, error) {
	if m < n {
		n, m = m, n
	}
	if m == n {
		return 0, errors.New("invalid argument")
	}
	rand.Seed(time.Now().Unix())
	return rand.Intn(m-n) + n, nil
}
