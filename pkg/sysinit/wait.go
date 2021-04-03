package sysinit

import (
	"context"
	"sync"

	"github.com/s3rj1k/ninit/pkg/log/logger"
)

// Cleanup is used to cancel context and wait for all goroutines to finish.
// Danger, Will Robinson:
//	this function is intended to be used in `defer` statement of main application
func Cleanup(wg *sync.WaitGroup, cancel context.CancelFunc, log logger.Logger) {
	cancel()
	wg.Wait()

	log.Infof("Coroutine cleanup finished\n")
}
