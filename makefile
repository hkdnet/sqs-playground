default: all

all: bin/consumer bin/producer elasticmq/jar/elasticmq-server-0.14.5.jar

bin/consumer: consumer/main.go
	GOOS=linux GOARCH=386 go build -o $@ consumer/main.go
bin/producer: producer/main.go
	GOOS=linux GOARCH=386 go build -o $@ producer/main.go
elasticmq/jar/elasticmq-server-0.14.5.jar:
	mkdir -p elasticmq/jar
	curl -o $@ https://s3-eu-west-1.amazonaws.com/softwaremill-public/elasticmq-server-0.14.5.jar

.PHONY: default all
