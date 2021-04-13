package configmap

import (
	"context"
	"sync"

	"github.com/s3rj1k/ninit/pkg/log/logger"
)

// Run starts kubernetes ConfigMap event watch and local config update inside container.
func Run(ctx context.Context, wg *sync.WaitGroup, c Config, log logger.Logger) error {
	cmWatch, err := Watch(ctx, wg, log, c.GetK8sNamespace(), c.GetK8sObjectName(), c.GetWatchInterval())
	if err != nil {
		return err
	}

	wg.Add(1)

	log.Tracef("Starting to read channel with kubernetes object (ADDED/MODIFIED/DELETED) events\n")

	go worker(ctx, wg, c, cmWatch)

	return nil
}
