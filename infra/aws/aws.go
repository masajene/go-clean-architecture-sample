package aws

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

var (
	cfg aws.Config
)

func init() {
	var err error
	cfg, err = config.LoadDefaultConfig(
		context.Background(),
		config.WithRegion("ap-northeast-1"),
	)
	if err != nil {
		panic(err)
	}
}
