FROM debian:10-slim

COPY guide2go /usr/local/bin/guide2go

RUN apt-get update && apt-get install ca-certificates -y && apt autoclean

ENTRYPOINT ["tail", "-f", "/dev/null"]


