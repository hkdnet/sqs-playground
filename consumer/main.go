package main

import (
	"fmt"
	"log"
	"time"

	"github.com/hkdnet/sqs-playground/gateway"
)

const sqsURL = "http://sqs:9324"
const queueName = "q"

type MessageViewer struct {
	client *gateway.SQSClient
}

func (v *MessageViewer) Process(msg gateway.Message) error {
	fmt.Printf("%s: %s\n", msg.MessageID, msg.Body)
	return v.client.DeleteMessage(msg.ReceiptHandle)
}

func main() {
	client, err := gateway.NewClient(queueName, gateway.NewLocalSessionCreator(sqsURL))
	if err != nil {
		log.Fatalf("new client: %s\n", err)
	}

	viewer := &MessageViewer{client: client}
	for {
		err = client.ReceiveMessage(viewer)
		if err != nil {
			log.Fatalf("recv message: %s\n", err)
		}

		time.Sleep(5 * time.Second)
	}
}
