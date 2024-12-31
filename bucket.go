package storage

import (
	"bytes"
	"context"

	"github.com/minio/minio-go/v7"
)

type Bucket struct {
	minioClient *minio.Client
	name        string
}

func (b *Bucket) Put(ctx context.Context, name string, data []byte) error {
	d := bytes.NewBuffer(data)
	_, err := b.minioClient.PutObject(ctx,
		b.name,
		name,
		d,
		int64(d.Len()),
		minio.PutObjectOptions{},
	)
	return err
}