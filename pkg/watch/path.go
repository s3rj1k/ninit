package watch

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/s3rj1k/ninit/pkg/hash"
)

/*
	https://github.com/fsnotify/fsnotify/issues/9#issuecomment-679936703

	Just wanted to add my two cents about polling mechanism. In Kubernetes, when a ConfigMap which is attached as a file on a pod changes,
	the file inside the pod changes as well but its last modified timestamp does not change. Size may not change depending on the change.
	So the GoConvey's polling approach will not work on this scenario.
	Francisco Beltrao from Microsoft offers a different approach which checks the hash of files instead of timesptamp.
	Here's a .NET Core implementation: https://github.com/fbeltrao/ConfigMapFileProvider
	It'll be slower than the aforementioned approach but it covers more ground.
*/

// Message describes output from watch.Path function.
type Message struct {
	IsChanged bool
	Error     error
	Message   string
}

// Path runs changes watcher for specified path using fast recursive file hashing.
func Path(ctx context.Context, wg *sync.WaitGroup, path string, interval time.Duration) <-chan Message {
	if strings.TrimSpace(path) == "" || interval == 0 {
		return nil
	}

	msg := make(chan Message, 1)

	wg.Add(1)

	go func(ctx context.Context, ch chan<- Message, path string, interval time.Duration) {
		ticker := time.NewTicker(interval)

		defer func(wg *sync.WaitGroup, ch chan<- Message, ticker *time.Ticker) {
			// defer inside goroutine works because we return when context is done
			ticker.Stop()
			close(ch)
			wg.Done()
		}(wg, ch, ticker)

		initialHash, err := hash.FromPath(path)
		if err != nil {
			ch <- Message{
				Error: fmt.Errorf("hash compute error, path '%s': %w", path, err),
			}
		}

		for {
			select {
			case <-ctx.Done():
				ch <- Message{
					Message: fmt.Sprintf("path '%s' watch is shutting down", path),
				}
				return

			case <-ticker.C:
				currentHash, err := hash.FromPath(path)
				if err != nil {
					ch <- Message{
						Error: fmt.Errorf("hash compute error, path '%s': %w", path, err),
					}

					continue
				}

				if currentHash != initialHash {
					ch <- Message{
						IsChanged: true,
						Message:   fmt.Sprintf("path '%s' change detected", path),
					}

					initialHash = currentHash
				}
			}
		}
	}(ctx, msg, path, interval)

	return msg
}
