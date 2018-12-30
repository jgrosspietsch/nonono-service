FROM alpine:3.8

COPY ./nonono-service /opt/nonono-service
EXPOSE 8888

ENTRYPOINT /opt/nonono-service