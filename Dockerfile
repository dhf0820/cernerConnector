FROM alpine:latest

ADD ./cerner_connector ./cerner_conn
RUN mkdir /root/tmp_images

ADD .env.cerner_conn_docker ./.env
EXPOSE 9200
ENTRYPOINT ["./cerner_conn"]
