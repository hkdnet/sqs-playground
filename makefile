default: all

all: bin/consumer bin/producer

GATEWAYS = $(wildcard gateway/*.go)

bin/consumer: consumer/main.go $(GATEWAYS)
	GOOS=linux GOARCH=386 go build -o $@ consumer/main.go
bin/producer: producer/main.go $(GATEWAYS)
	GOOS=linux GOARCH=386 go build -o $@ producer/main.go

.PHONY: default all
