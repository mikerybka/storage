package storage

import (
	"context"
	"fmt"
	"strings"

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

func (c *Client) BucketExists(ctx context.Context, name string) (bool, error) {
	return c.minioClient.BucketExists(ctx, name)
}

func (c *Client) Bucket(ctx context.Context, name string) (*Bucket, error) {
	ok, _ := c.BucketExists(ctx, name)
	if !ok {
		err := c.CreateBucket(ctx, name)
		if err != nil {
			return nil, fmt.Errorf("creating bucket: %s", err)
		}
	}
	return &Bucket{
		minioClient: c.minioClient,
		name:        name,
	}, nil
}

func (c *Client) Put(ctx context.Context, path string, data []byte) error {
	p := parsePath(path)
	if len(p) < 2 {
		return fmt.Errorf("path must be at least 2 parts")
	}
	bucketName := p[0]
	objectName := strings.Join(p[1:], "/")
	bucket, err := c.Bucket(ctx, bucketName)
	if err != nil {
		return err
	}
	return bucket.Put(ctx, objectName, data)
}

func parsePath(s string) []string {
	path := []string{}
	for _, p := range strings.Split(s, "/") {
		if p != "" {
			path = append(path, p)
		}
	}
	return path
}
