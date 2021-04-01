package configmap

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/s3rj1k/ninit/pkg/log/logger"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
)

const (
	// Field path constants that are specific to the internal API
	// representation.
	// https://github.com/kubernetes/apimachinery/blob/v0.21.0-beta.1/pkg/apis/meta/v1/types.go#L105
	ObjectNameField = "metadata.name"
)

func Watch(ctx context.Context, wg *sync.WaitGroup, log logger.Logger, namespace, cmName string, interval time.Duration) (<-chan *Object, error) {
	out := make(chan *Object, 1)

	restConfig, err := rest.InClusterConfig()
	if err != nil {
		return nil, fmt.Errorf("unable to get cluster config: %w", err)
	}

	clientset, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return nil, fmt.Errorf("unable to get cluster config: %w", err)
	}

	watchlist := cache.NewListWatchFromClient(
		clientset.CoreV1().RESTClient(),
		corev1.ResourceConfigMaps.String(),
		namespace,
		// https://github.com/kubernetes/kubernetes/issues/43299
		fields.OneTermEqualSelector(ObjectNameField, cmName),
	)

	_, controller := cache.NewInformer(
		watchlist,
		&corev1.ConfigMap{},
		interval,
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				if msg := addEvent(log, obj); msg != nil {
					out <- msg
				}
			},
			DeleteFunc: func(obj interface{}) {
				if msg := deleteEvent(log, obj); msg != nil {
					out <- msg
				}
			},
			UpdateFunc: func(oldObj, newObj interface{}) {
				if msg := updateEvent(log, oldObj, newObj); msg != nil {
					out <- msg
				}
			},
		},
	)

	stop := make(chan struct{})

	wg.Add(1)

	go func(wg *sync.WaitGroup, stop chan struct{}) {
		controller.Run(stop)
		wg.Done()
	}(wg, stop)

	wg.Add(1)

	go func(ctx context.Context, wg *sync.WaitGroup, stop chan struct{}) {
		<-ctx.Done()
		close(stop)
		wg.Done()
	}(ctx, wg, stop)

	return out, nil
}
