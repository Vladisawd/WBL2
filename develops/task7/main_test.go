package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestOr(t *testing.T) {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	finish := time.Since(start)
	norma := time.Second * 2

	assert.Greater(t, norma, finish)
}
