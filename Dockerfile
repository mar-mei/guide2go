FROM golang:alpine3.10 as builder

RUN mkdir /app
COPY *.go /app/
WORKDIR /app
RUN go mod init main
RUN go get
RUN go build -o guide2go

FROM alpine:3.17.2
ENV USER=docker
ENV UID=12345
ENV GID=23456

RUN addgroup "${USER}" -g "${GID}"
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "$(pwd)" \
    --ingroup "$USER" \
    --no-create-home \
    --uid "$UID" \
    "$USER"

RUN mkdir /app
RUN chown ${USER} -R /app
WORKDIR /app
COPY --from=builder --chown="${USER}":"${GID}" /app/guide2go /usr/local/bin/guide2go
COPY --chown="${USER}":"${GID}" sample-config.yaml /app/sample-config.yaml

USER "${USER}"
CMD [ "guide2go", "--config", "/app/config.yaml" ]