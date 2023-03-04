FROM golang:alpine3.10 as builder

RUN mkdir /app
COPY *.go /app/
WORKDIR /app
RUN go mod init main
RUN go get
RUN go build -o guide2go

FROM debian:10-slim

COPY --from=builder /app/guide2go /usr/local/bin/guide2go
COPY sample-config.yaml /config/sample-config.yaml

RUN apt-get update && apt-get --no-install-recommends -y \
install ca-certificates \
&& apt autoclean \
&& rm -rf /var/lib/apt/lists/*

CMD [ "guide2go", "--config", "/config/config.yaml" ]