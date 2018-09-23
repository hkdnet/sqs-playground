package main

import (
	"fmt"
	"log"
	"time"

	"github.com/hkdnet/sqs-playground/gateway"
)

const sqsURL = "http://sqs:9324"
const queueName = "q"

func main() {
	client, err := gateway.NewClient(queueName, gateway.NewLocalSessionCreator(sqsURL))
	if err != nil {
		log.Fatalf("new client: %s\n", err)
	}

	n := 1
	for {
		limit := time.After(3 * time.Second)
		msg := time.Now().Format("15:04:05")

		err = client.SendMessage(fmt.Sprintf("%d", n), msg)
		n++
		if err != nil {
			log.Fatalf("send message: %s\n", err)
		}
		log.Println("Successfully sent a message.")

		<-limit
	}
}
