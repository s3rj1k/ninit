package configmap

import (
	"context"
	"sync"
)

func worker(
	ctx context.Context,
	wg *sync.WaitGroup,
	c Config,
	cmWatch <-chan *Object,
) {
	// pause path watch when write/delete operation is in progress
	pause := c.GetPauseChannel()

	for {
		select {
		case <-ctx.Done():
			wg.Done()

			return

		case obj := <-cmWatch:
			if obj.IsModified() || obj.IsAdded() {
				obj.log.Infof("%s\n", obj.String())

				pause <- true

				if err := obj.Write(c.GetK8sBaseDirectory()); err != nil {
					obj.log.Errorf("%v\n", err)
				}

				pause <- false
			}

			if obj.IsDeleted() {
				obj.log.Infof("%s\n", obj.String())

				pause <- true

				if err := obj.RemoveFilesFromDir(c.GetK8sBaseDirectory()); err != nil {
					obj.log.Errorf("%v\n", err)
				}

				pause <- false
			}
		}
	}
}
