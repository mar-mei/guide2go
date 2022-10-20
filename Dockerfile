FROM debian:10-slim

ENV IMAGES_PATH=/data/images
ENV PORT=3000

COPY guide2go /usr/local/bin/guide2go

ENTRYPOINT ["tail", "-f", "/dev/null"]


