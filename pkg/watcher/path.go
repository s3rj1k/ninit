package watcher

import (
	"context"
	"strings"
	"sync"
	"time"
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

// Path runs changes watcher for specified path using fast recursive file hashing.
func Path(ctx context.Context, wg *sync.WaitGroup, path string, interval time.Duration, pause <-chan bool) <-chan Message {
	if strings.TrimSpace(path) == "" || interval == 0 {
		return nil
	}

	msg := make(chan Message, 1)

	wg.Add(1)

	go worker(ctx, wg,
		&workerConfig{
			ch:       msg,
			interval: interval,
			path:     path,
			pause:    pause,
		},
	)

	return msg
}
