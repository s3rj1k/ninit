package watcher

import (
	"context"
	"sync"
	"time"

	"github.com/s3rj1k/ninit/pkg/hash"
)

type workerConfig struct {
	ch       chan<- Message
	pause    <-chan bool
	path     string
	interval time.Duration
}

func worker(ctx context.Context, wg *sync.WaitGroup, wc *workerConfig) {
	ticker := time.NewTicker(wc.interval)
	ignoreTicks := false

	defer func(wg *sync.WaitGroup, ch chan<- Message, ticker *time.Ticker) {
		// defer inside goroutine works because we return when context is done
		ticker.Stop()
		close(ch)
		wg.Done()
	}(wg, wc.ch, ticker)

	initialHash, err := hash.FromPath(wc.path)
	if err != nil {
		wc.ch <- hashError(wc.path, err)
	}

	for {
		select {
		case <-ctx.Done():
			wc.ch <- shutdown(wc.path)

			return

		case ignoreTicks = <-wc.pause:
			if ignoreTicks {
				wc.ch <- paused(wc.path)
			} else {
				wc.ch <- resumed(wc.path)
			}

		case <-ticker.C:
			if ignoreTicks {
				continue
			}

			t1 := time.Now()

			currentHash, err := hash.FromPath(wc.path)
			if err != nil {
				wc.ch <- hashError(wc.path, err)

				continue
			}

			t2 := time.Now()

			if currentHash != initialHash {
				wc.ch <- change(wc.path, t2.Sub(t1))

				initialHash = currentHash
			}
		}
	}
}
