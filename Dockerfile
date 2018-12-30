FROM alpine:3.8

COPY ./nonono-service /opt/
EXPOSE 8080

ENTRYPOINT ["/opt/nonono-service"]