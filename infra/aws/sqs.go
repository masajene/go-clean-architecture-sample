package aws

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/google/uuid"
)

type SQSMessagePayload struct {
	AreaCode               int    `json:"area_code"`
	Title                  string `json:"title"`
	DistributionID         string `json:"distribution_id"`
	DistributionType       string `json:"distribution_type"`
	RowXMLPath             string `json:"row_xml_path"`
	IntegrationType        int    `json:"integration_type"`
	IntegrationDestination string `json:"integration_destination"`
}

func SendSQSMessage(s *session.Session, queueURL string, payload *SQSMessagePayload) error {
	svc := sqs.New(
		s,
	)

	u, err := uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
		return err
	}

	j, _ := json.Marshal(payload)

	// 送信内容を作成
	params := &sqs.SendMessageInput{
		MessageBody:    aws.String(string(j)),
		MessageGroupId: aws.String(u.String()),
		QueueUrl:       aws.String(queueURL),
	}

	sqsRes, err := svc.SendMessage(params)
	if err != nil {
		fmt.Printf(err.Error())
		return err
	}

	fmt.Println("SQSMessageID", *sqsRes.MessageId)
	return nil
}
