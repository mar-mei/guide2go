FROM nginx:1.22.0

ENV images_path="/data/images"
ENV ip_address="127.0.0.1"

COPY guide2go /usr/local/bin/guide2go

COPY nginx.conf /tmp/nginx.conf
RUN envsubst < /tmp/nginx.conf > /etc/nginx/nginx.conf
RUN apt update && apt-get install cron -y
COPY cronjob /tmp/cronjob
RUN crontab /tmp/cronjob
CMD ["/usr/sbin/nginx", "-g", "daemon off;"]
