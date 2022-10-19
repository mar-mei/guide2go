FROM nginx:1.22.0
WORKDIR /app/
RUN adduser --disabled-password guide2go
ENV images_path="/data/images"

COPY --chown=guide2go:guide2go guide2go /app/guide2go
RUN chown guide2go /app
COPY nginx.conf /tmp/nginx.conf
RUN envsubst < /tmp/nginx.conf > /etc/nginx/nginx.conf
RUN chown guide2go /app/guide2go

ENTRYPOINT ["/usr/sbin/nginx", "-g", "daemon off;"]
USER guide2go