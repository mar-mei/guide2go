FROM nginx:1.22.0

COPY guide2go /config/guide2go

COPY nginx.conf /etc/nginx/nginx.conf

CMD [ "nginx" ]