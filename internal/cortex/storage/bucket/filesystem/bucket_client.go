// Copyright (c) The Cortex Authors.
// Licensed under the Apache License 2.0.

package filesystem

import (
	"github.com/thanos-io/thanos/pkg/objstore"
	"github.com/thanos-io/thanos/pkg/objstore/filesystem"
)

// NewBucketClient creates a new filesystem bucket client
func NewBucketClient(cfg Config) (objstore.Bucket, error) {
	return filesystem.NewBucket(cfg.Directory)
}
