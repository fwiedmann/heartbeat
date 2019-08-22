FROM alpine:3.10.2

WORKDIR /code

COPY heartbeat .

USER 1000

ENTRYPOINT [ "./heartbeat" ]