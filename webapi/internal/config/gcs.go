package config

import (
	"context"

	"cloud.google.com/go/storage"
	"github.com/adafatya/micro-services/webapi/pkg/util"
	"google.golang.org/api/option"
)

func NewGCSBucket() *storage.BucketHandle {
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(util.GetEnv("GCS_KEY_PATH", "")))
	if err != nil {
		panic(err)
	}

	return client.Bucket(util.GetEnv("GCS_BUCKET", ""))
}
