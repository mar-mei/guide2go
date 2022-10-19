FROM nginx:1.22.0
WORKDIR /app/

ENV images_path="/data/images"

COPY guide2go /app/guide2go

COPY nginx.conf /tmp/nginx.conf
RUN envsubst < /tmp/nginx.conf > /etc/nginx/nginx.conf

CMD ["/usr/sbin/nginx", "-g", "daemon off;"]
