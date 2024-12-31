package storage

import (
	"bytes"
	"context"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func NewClient(accessKey, secretKey, endpointURL string) *Client {
	c, err := minio.New(endpointURL, &minio.Options{
		Secure: true,
		Creds: credentials.NewStaticV4(
			accessKey,
			secretKey,
			"",
		),
	})
	if err != nil {
		panic(err)
	}
	return &Client{
		minioClient: c,
	}
}

type Client struct {
	minioClient *minio.Client
}

func (c *Client) CreateBucket(ctx context.Context, name string) error {
	return c.minioClient.MakeBucket(ctx, name, minio.MakeBucketOptions{})
}

func (c *Client) Put(ctx context.Context, bucket, name string, data []byte) error {
	b := bytes.NewBuffer(data)
	_, err := c.minioClient.PutObject(ctx,
		bucket,
		name,
		b,
		int64(b.Len()),
		minio.PutObjectOptions{},
	)
	return err
}
