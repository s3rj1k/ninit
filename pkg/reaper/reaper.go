package reaper

import (
	"context"
	"sync"
	"time"
)

const cooldownTime = 250 * time.Millisecond

// Run starts goroutine that will reap zombie processes when appropriate signal is sent.
func Run(ctx context.Context, wg *sync.WaitGroup) <-chan Message {
	out := make(chan Message, 1)

	wg.Add(1)

	go worker(ctx, wg, out)

	return out
}
