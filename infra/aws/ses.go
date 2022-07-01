package aws

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

type EmailData struct {
	To      string
	Subject string
	Body    string
}

func (data *EmailData) SendEmail(s *session.Session) error {

	client := ses.New(s)
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{
				aws.String(data.To),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Text: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(data.Body),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String("UTF-8"),
				Data:    aws.String(data.Subject),
			},
		},
		Source: aws.String("aws.mizuno@gmail.com"),
	}
	_, err := client.SendEmail(input)
	fmt.Printf("ses fail send mail to %s", err)
	return err
}
