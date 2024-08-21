package repository

import (
	"context"
	"io"
	"mime/multipart"

	"cloud.google.com/go/storage"
)

type GCSRepository struct {
	Bucket *storage.BucketHandle
}

func NewGCSRepository(bucket *storage.BucketHandle) *GCSRepository {
	return &GCSRepository{
		Bucket: bucket,
	}
}

func (g *GCSRepository) Upload(ctx context.Context, file multipart.File, name string) error {
	writer := g.Bucket.Object(name).NewWriter(ctx)
	writer.ACL = []storage.ACLRule{{Entity: storage.AllUsers, Role: storage.RoleReader}}
	if _, err := io.Copy(writer, file); err != nil {
		return err
	}

	if err := writer.Close(); err != nil {
		return err
	}

	return nil
}
