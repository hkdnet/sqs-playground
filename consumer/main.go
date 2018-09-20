package main

import (
	"fmt"
	"log"

	"github.com/hkdnet/sqs-playground/gateway"
)

const sqsURL = "http://sqs:9324"
const queueName = "q"

func main() {
	_, err := gateway.NewClient(queueName, gateway.NewLocalSessionCreator(sqsURL))
	if err != nil {
		log.Fatalf("new client: %s", err)
	}
	fmt.Println("ok")
}
