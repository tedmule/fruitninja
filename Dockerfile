FROM 192.168.6.99/docker/alpine:3.9
#FROM alpine:3.9
LABEL MAINTAINER="Ted <daddvted@gmail.com>"

COPY ninja /usr/bin/
EXPOSE 8080
ENTRYPOINT ["ninja"]
