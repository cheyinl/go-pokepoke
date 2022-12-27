package pokepoke

import (
	"time"
)

// SignalCh contain structure for signaling one go routine.
type SignalCh struct {
	c chan time.Time
}

// NewSignalCh create a new instance of SignalCh.
func NewSignalCh() *SignalCh {
	return &SignalCh{
		c: make(chan time.Time, 1),
	}
}

// Wait return the channel to wait for signal from other Go routine.
func (s *SignalCh) Wait() <-chan time.Time {
	return s.c
}

// Poke signals the Go routine waiting for signals.
func (s *SignalCh) Poke() {
	select {
	case s.c <- time.Now():
	default:
	}
}
