package hash

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/minio/highwayhash"
)

const hashKey = "31220946A728567B509734212C3295856D76134229E959910805438F52169117"

// FromPath returns hash of hashes for all files inside path.
// This function is inspired by https://pkg.go.dev/golang.org/x/mod/sumdb/dirhash
func FromPath(path string) (string, error) {
	files, err := getListOfFilesFromPath(path)
	if err != nil {
		return "", fmt.Errorf("hash error, path '%s': %w", path, err)
	}

	key, err := hex.DecodeString(hashKey)
	if err != nil {
		return "", fmt.Errorf("hash error, path '%s': %w", path, err)
	}

	h, err := highwayhash.New(key)
	if err != nil {
		return "", fmt.Errorf("hash error, path '%s': %w", path, err)
	}

	for _, file := range files {
		if strings.Contains(file, "\n") {
			return "", fmt.Errorf("hash error, path '%s': filenames with newlines are not supported", path)
		}

		r, err := os.OpenFile(file, os.O_RDONLY, 0)
		if err != nil {
			return "", fmt.Errorf("hash error, path '%s': %w", path, err)
		}

		hf, err := highwayhash.New(key)
		if err != nil {
			return "", fmt.Errorf("hash error, path '%s': %w", path, err)
		}

		_, err = io.Copy(hf, r)
		_ = r.Close()

		if err != nil {
			return "", fmt.Errorf("hash error, path '%s': %w", path, err)
		}

		_, err = fmt.Fprintf(h, "%x  %s\n", hf.Sum(nil), file)
		if err != nil {
			return "", fmt.Errorf("hash error, path '%s': %w", path, err)
		}
	}

	return base64.StdEncoding.EncodeToString(h.Sum(nil)), nil
}
