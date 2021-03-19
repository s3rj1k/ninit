package reaper

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Message describes output from reaper.Path function.
type Message struct {
	Error   error
	Message string
}

// Run starts goroutine that will reap zombie processes when appropriate signal is sent.
func Run(ctx context.Context, wg *sync.WaitGroup) <-chan Message {
	out := make(chan Message, 1)

	wg.Add(1)

	go func(ctx context.Context, ch chan<- Message) {
		notify := make(chan os.Signal, 1)
		signal.Notify(notify, syscall.SIGCHLD)

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
					var status syscall.WaitStatus

					// wait for orphaned zombie process
					// https://man7.org/linux/man-pages/man2/wait.2.html
					pid, err := syscall.Wait4(-1, &status, syscall.WNOHANG|syscall.WCONTINUED, nil)

					if syscall.ECHILD == err {
						// no un-reaped child(ren) exist
						ch <- Message{
							Message: "reaper cleanup: no (more) zombies found",
						}

						break
					}

					switch {
					case pid == 0:
						// one or more child(ren) exist that have not yet changed state
						time.Sleep(250 * time.Millisecond)
					case pid == -1:
						// error from syscall
						ch <- Message{
							Error: fmt.Errorf("reaper error: %w", err),
						}
					case pid > 0:
						// child was reaped
						ch <- Message{
							Message: fmt.Sprintf("reaper cleanup: pid=%d, status=%+v", pid, status),
						}
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
