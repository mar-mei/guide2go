FROM debian:10-slim

LABEL version="2.0"

COPY guide2go /usr/local/bin/guide2go

RUN apt-get update && apt-get install ca-certificates -y && apt autoclean

CMD ["guide2go"]