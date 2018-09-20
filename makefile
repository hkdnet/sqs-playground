default: all

all: bin/consumer bin/producer

bin/consumer: consumer/main.go
	GOOS=linux GOARCH=386 go build -o $@ consumer/main.go
bin/producer: producer/main.go
	GOOS=linux GOARCH=386 go build -o $@ producer/main.go

.PHONY: default all
