FROM alpine:3.7
LABEL MAINTAINER="Ted <ski2per@gmail.com>"

COPY fruitninja /usr/bin/
EXPOSE 8080
ENTRYPOINT ["fruitninja"]
