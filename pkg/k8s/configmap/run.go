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

	go func(
		ctx context.Context,
		wg *sync.WaitGroup,
		c Config,
		cmWatch <-chan *Object,
	) {
		for {
			select {
			case <-ctx.Done():
				wg.Done()

				return

			case obj := <-cmWatch:
				if obj.IsModified() || obj.IsAdded() {
					obj.log.Infof("%s\n", obj.String())

					if err := obj.Write(c.GetK8sBaseDirectory()); err != nil {
						obj.log.Errorf("%v\n", err)
					}
				}

				if obj.IsDeleted() {
					obj.log.Infof("%s\n", obj.String())

					if err := obj.RemoveFilesFromDir(c.GetK8sBaseDirectory()); err != nil {
						obj.log.Errorf("%v\n", err)
					}
				}
			}
		}
	}(ctx, wg, c, cmWatch)

	return nil
}
