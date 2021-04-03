package configmap

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/s3rj1k/ninit/pkg/utils"
)

// RemoveFilesFromDir removes regular files from directory path, without recursion.
// Function can be provided with a list of file name exception,
// this files will not be removed.
func (obj *Object) RemoveFilesFromDir(path string, except ...string) error {
	path = filepath.Clean(path)

	files, err := os.ReadDir(path)
	if err != nil {
		return fmt.Errorf("configMap '%s/%s' event '%s', cleaning path '%s' error: %w",
			obj.Namespace, obj.Name, obj.eventType, path, err)
	}

	files = utils.FilterDirEntries(
		files,
		func(x string) bool {
			return !utils.IsStringInSlice(x, except)
		},
	)

	for _, file := range files {
		if !file.Type().IsRegular() {
			continue
		}

		filePath := filepath.Join(path, file.Name())
		obj.log.Infof("ConfigMap '%s/%s' event '%s', removing file '%s'\n", obj.Namespace, obj.Name, obj.eventType, filePath)

		if err := os.Remove(filePath); err != nil {
			return fmt.Errorf("configMap '%s/%s' event '%s', removing file '%s' error: %w",
				obj.Namespace, obj.Name, obj.eventType, filePath, err)
		}
	}

	return nil
}

// Write syncs files content from kubernetes config map to container local directory.
// No check is preformed on destination file vs source file, content is overwritten.
// Files that are absent in object data key are also removed.
func (obj *Object) Write(basePath string) error {
	// https://kubernetes.io/docs/concepts/configuration/configmap/#configmap-object
	except := make([]string, 0, len(obj.Data)+len(obj.BinaryData))

	for k := range obj.Data {
		except = append(except, k)
	}

	for k := range obj.BinaryData {
		except = append(except, k)
	}

	if err := obj.RemoveFilesFromDir(basePath, except...); err != nil {
		return err
	}

	for k, v := range obj.Data {
		path := filepath.Join(basePath, k)
		obj.log.Infof("ConfigMap '%s/%s' event '%s', writing file '%s'\n", obj.Namespace, obj.Name, obj.eventType, path)

		if err := os.WriteFile(path, []byte(v), 0644); err != nil {
			return fmt.Errorf("configMap '%s/%s' event '%s', writing file '%s' error: %w",
				obj.Namespace, obj.Name, obj.eventType, path, err)
		}
	}

	for k, v := range obj.BinaryData {
		path := filepath.Join(basePath, k)
		obj.log.Infof("ConfigMap '%s/%s' event '%s', writing file '%s'\n", obj.Namespace, obj.Name, obj.eventType, path)

		if err := os.WriteFile(path, v, 0644); err != nil {
			return fmt.Errorf("configMap '%s/%s' event '%s', writing file '%s' error: %w",
				obj.Namespace, obj.Name, obj.eventType, path, err)
		}
	}

	return nil
}
