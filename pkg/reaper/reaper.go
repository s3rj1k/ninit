package reaper

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"time"

	"golang.org/x/sys/unix"
)

const cooldownTime = 250 * time.Millisecond

// Run starts goroutine that will reap zombie processes when appropriate signal is sent.
func Run(ctx context.Context, wg *sync.WaitGroup) <-chan Message {
	out := make(chan Message, 1)

	wg.Add(1)

	go func(ctx context.Context, ch chan<- Message) {
		notify := make(chan os.Signal, 1)
		signal.Notify(notify, unix.SIGCHLD)

		defer func(wg *sync.WaitGroup, notify chan<- os.Signal, ch chan<- Message) {
			// defer inside goroutine works because we return when context is done
			signal.Stop(notify)
			close(ch)
			close(notify)
			wg.Done()
		}(wg, notify, ch)

		for {
			select {
			case <-ctx.Done():
				return
			case <-notify:
				for {
					msg, ok := syscallWait()
					if msg != nil {
						ch <- *msg
					}

					if ok {
						break
					}

					if msg == nil {
						time.Sleep(cooldownTime)
					}

					select {
					case <-ctx.Done():
						return
					default:
					}
				}
			}
		}
	}(ctx, out)

	return out
}
