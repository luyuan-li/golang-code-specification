FROM golang:1.18-alpine3.15 as builder

# Set up dependencies
ENV PACKAGES make gcc git libc-dev linux-headers bash

COPY  . $GOPATH/src
WORKDIR $GOPATH/src

# Install minimum necessary dependencies, build binary
# RUN apk add --no-cache $PACKAGES && make all
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && \
apk add --no-cache $PACKAGES && make all

FROM alpine:3.15
WORKDIR /root/
COPY --from=builder /go/src/project-name /usr/local/bin
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories

CMD ["project-name", "start"]
