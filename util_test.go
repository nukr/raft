package raft

import "testing"

func TestRandom(t *testing.T) {
	r, _ := random(1, 100)
	if !(r >= 1 && r <= 100) {
		t.Errorf("error random value %d", r)
	}
	r, _ = random(400, 200)
	if !(r >= 200 && r <= 400) {
		t.Errorf("error random value %d", r)
	}

	_, err := random(200, 200)
	if err == nil {
		t.Errorf("expected error, but got nil")
	}
}
