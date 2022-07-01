package aws

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

func GetSSMParams(s *session.Session, path string) (*string, error) {
	svc := ssm.New(s)

	res, err := svc.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String(path),
		WithDecryption: aws.Bool(true),
	})
	if err != nil {
		fmt.Printf(err.Error())
		return nil, err
	}
	fmt.Printf(*res.Parameter.Value)
	return res.Parameter.Value, err
}
