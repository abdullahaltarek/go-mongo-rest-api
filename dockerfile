FROM alpine:latest

ADD ./bin /project

USER root

RUN chmod +x /project/go-mongo-api

ENTRYPOINT /project/go-mongo-api

EXPOSE 8076