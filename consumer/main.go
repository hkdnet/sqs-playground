package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hkdnet/sqs-playground/gateway"
	"github.com/pkg/errors"
)

const sqsURL = "http://sqs:9324"
const queueName = "q"

type MessageViewer struct {
	client *gateway.SQSClient

	ctx context.Context
}

func (v *MessageViewer) Process(msg gateway.Message) error {
	ch := make(chan error)
	go func() {
		fmt.Printf("%s: %s\n", msg.MessageID, msg.Body)
		ch <- v.client.DeleteMessage(msg.ReceiptHandle)
	}()

	select {
	case <-v.ctx.Done():
		return errors.Wrap(v.ctx.Err(), msg.MessageID)
	case err := <-ch:
		return err
	}
}

func main() {
	client, err := gateway.NewClient(queueName, gateway.NewLocalSessionCreator(sqsURL))
	if err != nil {
		log.Fatalf("new client: %s\n", err)
	}

	for {
		func() {
			limit := time.After(5 * time.Second)
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			viewer := &MessageViewer{
				client: client,
				ctx:    ctx,
			}
			err = client.ReceiveMessage(viewer)
			if err != nil {
				log.Fatalf("recv message: %s\n", err)
			}

			<-limit
		}()
	}
}
