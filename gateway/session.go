package gateway

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func NewLocalSessionCreator(sqsURL string) *LocalSessionCreator {
	return &LocalSessionCreator{sqsURL: sqsURL}
}

type LocalSessionCreator struct{ sqsURL string }

func (c *LocalSessionCreator) Create() *session.Session {
	cred := credentials.NewStaticCredentials("x", "x", "x") // dummy
	return session.New(&aws.Config{
		Credentials: cred,
		Region:      aws.String("elasticmq"),
		Endpoint:    aws.String(c.sqsURL),
	})
}

type SessionCreator interface {
	Create() *session.Session
}
