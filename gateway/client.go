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

type Message struct {
	MessageID     string
	Body          string
	ReceiptHandle string
}

type MessageProcessor interface {
	Process(m Message) error
}

func (c *SQSClient) ReceiveMessage(processor MessageProcessor) error {
	resp, err := c.svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl:            c.queueURL,
		MaxNumberOfMessages: aws.Int64(1),
	})

	if err != nil {
		return errors.Wrap(err, "recv message")
	}

	for _, m := range resp.Messages {
		msg := Message{
			MessageID:     *m.MessageId,
			Body:          *m.Body,
			ReceiptHandle: *m.ReceiptHandle,
		}

		err = processor.Process(msg)
		if err != nil {
			return errors.Wrap(err, "processing")
		}
	}

	return nil
}

// メッセージを送信する
func (c *SQSClient) SendMessage(groupID string, body string) error {
	params := &sqs.SendMessageInput{
		MessageBody:            aws.String(body),
		QueueUrl:               c.queueURL,
		MessageDeduplicationId: aws.String(groupID),
		MessageGroupId:         aws.String(groupID),
	}

	_, err := c.svc.SendMessage(params)
	if err != nil {
		return errors.Wrap(err, "send message")
	}

	return nil
}

// DeleteMessage trys to delete a message by receiptHandle.
func (c *SQSClient) DeleteMessage(receiptHandle string) error {
	params := &sqs.DeleteMessageInput{
		QueueUrl:      c.queueURL,
		ReceiptHandle: aws.String(receiptHandle),
	}

	_, err := c.svc.DeleteMessage(params)
	if err != nil {
		return errors.Wrap(err, "delete message")
	}

	return nil
}
