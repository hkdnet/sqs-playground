FROM openjdk:8-alpine

COPY jar/elasticmq-server-0.14.5.jar /
COPY custom.conf /
ENTRYPOINT ["/usr/bin/java", "-Dconfig.file=custom.conf", "-jar", "/elasticmq-server-0.14.5.jar"]

EXPOSE 9324

CMD ["-help"]
