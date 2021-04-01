package configmap

import (
	"reflect"

	"github.com/s3rj1k/ninit/pkg/log/logger"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/watch"
)

func addEvent(log logger.Logger, obj interface{}) *Object {
	cm, ok := obj.(*corev1.ConfigMap)
	if !ok {
		return nil
	}

	return &Object{
		ConfigMap: cm,

		eventType: watch.Added,
		log:       log,
	}
}

func deleteEvent(log logger.Logger, obj interface{}) *Object {
	cm, ok := obj.(*corev1.ConfigMap)
	if !ok {
		return nil
	}

	return &Object{
		ConfigMap: cm,

		eventType: watch.Deleted,
		log:       log,
	}
}

func updateEvent(log logger.Logger, oldObj, newObj interface{}) *Object {
	oldCM, ok := oldObj.(*corev1.ConfigMap)
	if !ok {
		return nil
	}

	newCM, ok := newObj.(*corev1.ConfigMap)
	if !ok {
		return nil
	}

	if reflect.DeepEqual(oldCM.Data, newCM.Data) && reflect.DeepEqual(oldCM.BinaryData, newCM.BinaryData) {
		return nil
	}

	return &Object{
		ConfigMap: newCM,

		eventType: watch.Modified,
		log:       log,
	}
}
