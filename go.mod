module github.com/s3rj1k/ninit

go 1.16

require (
	github.com/minio/highwayhash v1.0.2
	github.com/s3rj1k/ninit/pkg/log/logger v0.0.0-00010101000000-000000000000
	golang.org/x/sys v0.0.0-20210326220804-49726bf1d181
	k8s.io/api v0.20.5
	k8s.io/apimachinery v0.20.5
	k8s.io/client-go v0.20.5
)

replace github.com/s3rj1k/ninit/pkg/log/logger => ./pkg/log/logger
