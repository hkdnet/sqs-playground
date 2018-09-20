package main

import (
	"fmt"
	"log"

	"github.com/hkdnet/sqs-playground/gateway"
)

const sqsURL = "http://sqs:9324"
const queueName = "q"

type MessageViewer struct{}

func (v MessageViewer) Process(msg gateway.Message) error {
	fmt.Printf("%s: %s\n", msg.MessageID, msg.Body)
	return nil
}

func main() {
	client, err := gateway.NewClient(queueName, gateway.NewLocalSessionCreator(sqsURL))
	if err != nil {
		log.Fatalf("new client: %s\n", err)
	}

	err = client.ReceiveMessage(MessageViewer{})
	if err != nil {
		log.Fatalf("recv message: %s\n", err)
	}
	log.Println("done")
}
