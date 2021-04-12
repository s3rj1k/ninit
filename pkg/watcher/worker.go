package watcher

import (
	"context"
	"sync"
	"time"

	"github.com/s3rj1k/ninit/pkg/hash"
)

func worker(ctx context.Context, wg *sync.WaitGroup, ch chan<- Message, path string, interval time.Duration, pause <-chan bool) {
	ticker := time.NewTicker(interval)
	ignoreTicks := false

	defer func(wg *sync.WaitGroup, ch chan<- Message, ticker *time.Ticker) {
		// defer inside goroutine works because we return when context is done
		ticker.Stop()
		close(ch)
		wg.Done()
	}(wg, ch, ticker)

	initialHash, err := hash.FromPath(path)
	if err != nil {
		ch <- hashError(path, err)
	}

	for {
		select {
		case <-ctx.Done():
			ch <- shutdown(path)

			return

		case ignoreTicks = <-pause:
			if ignoreTicks {
				ch <- paused(path)
			} else {
				ch <- resumed(path)
			}

		case <-ticker.C:
			if ignoreTicks {
				continue
			}

			t1 := time.Now()

			currentHash, err := hash.FromPath(path)
			if err != nil {
				ch <- hashError(path, err)

				continue
			}

			t2 := time.Now()

			if currentHash != initialHash {
				ch <- change(path, t2.Sub(t1))

				initialHash = currentHash
			}
		}
	}
}
