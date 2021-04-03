module github.com/s3rj1k/ninit/cmd/ninit-k8s-cm

go 1.16

replace (
	github.com/s3rj1k/ninit => ../../
	github.com/s3rj1k/ninit/pkg/log/logger => ../../pkg/log/logger
	k8s.io/klog/v2 => ../../pkg/log/klog/v2
)

require (
	github.com/s3rj1k/ninit v0.0.0-00010101000000-000000000000
	github.com/s3rj1k/ninit/pkg/log/logger v0.0.0-00010101000000-000000000000
	k8s.io/klog/v2 v2.8.0
)
