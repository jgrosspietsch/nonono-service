FROM ubuntu:18.04

COPY ./nonono-service /opt/nonono-service
EXPOSE 8888

ENTRYPOINT /opt/nonono-service