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
}

func (v *MessageViewer) Process(msg gateway.Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	ch := make(chan error)
	go func() {
		fmt.Printf("%s: %s\n", msg.MessageID, msg.Body)
		ch <- v.client.DeleteMessage(msg.ReceiptHandle)
	}()

	select {
	case <-ctx.Done():
		return errors.Wrap(ctx.Err(), msg.MessageID)
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
		viewer := &MessageViewer{
			client: client,
		}
		err = client.ReceiveMessage(viewer)
		if err != nil {
			log.Fatalf("recv message: %s\n", err)
		}
	}
}
