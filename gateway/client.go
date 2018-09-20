package gateway

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/pkg/errors"
)

// SQSClient is a wrapper for SQS
type SQSClient struct {
	svc       *sqs.SQS
	queueName string

	queueURL *string // for aws parameter
}

// NewClient creates a new SQSClient
func NewClient(queueName string, sessionCreator SessionCreator) (*SQSClient, error) {
	session := sessionCreator.Create()
	svc := sqs.New(session)
	o, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{QueueName: aws.String(queueName)})
	if err != nil {
		return nil, errors.Wrap(err, "get queue url")
	}

	return &SQSClient{svc: svc, queueName: queueName, queueURL: o.QueueUrl}, nil
}
