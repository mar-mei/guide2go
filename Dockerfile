FROM nginx:1.22.0
WORKDIR /app/

ENV images_path="/data/images"

COPY guide2go /app/guide2go

COPY nginx.conf /tmp/nginx.conf
RUN envsubst < /tmp/nginx.conf > /etc/nginx/nginx.conf
RUN apt update && apt-get install cron
COPY cronjob /tmp/cronjob
RUN contab /tmp/cronjob
CMD ["/usr/sbin/nginx", "-g", "daemon off;"]
