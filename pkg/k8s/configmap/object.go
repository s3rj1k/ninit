package configmap

import (
	"fmt"

	"github.com/s3rj1k/ninit/pkg/log/logger"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/watch"
)

type Object struct {
	*corev1.ConfigMap

	log       logger.Logger
	eventType watch.EventType
}

func (obj *Object) IsAdded() bool {
	return obj.eventType == watch.Added
}

func (obj *Object) IsModified() bool {
	return obj.eventType == watch.Modified
}

func (obj *Object) IsDeleted() bool {
	return obj.eventType == watch.Deleted
}

func (obj *Object) String() string {
	return fmt.Sprintf("ConfigMap '%s/%s' got event '%s'", obj.Namespace, obj.Name, obj.eventType)
}
