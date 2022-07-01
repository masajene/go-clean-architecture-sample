package aws

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"io"
)

func GetObject(ctx context.Context, body io.Writer, bucket, key string) (int64, error) {
	input := &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &key,
	}

	client := s3.NewFromConfig(cfg)

	out, err := client.GetObject(ctx, input)
	if err != nil {
		return 0, err
	}

	defer out.Body.Close()

	return io.Copy(body, out.Body)
}

func PutObject(ctx context.Context, body io.Reader, bucket, key, contentType string) error {
	input := &s3.PutObjectInput{
		Bucket:      &bucket,
		Key:         &key,
		Body:        body,
		ContentType: &contentType,
	}
	client := s3.NewFromConfig(cfg)

	_, err := client.PutObject(ctx, input)
	if err != nil {
		return err
	}

	return err
}
