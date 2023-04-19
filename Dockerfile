FROM golang:1.18-alpine3.15 as builder

# Set up dependencies
ENV PACKAGES make gcc git libc-dev linux-headers bash
ENV GO111MODULE  on
ENV BINARY_NAME project-name

ARG GITUSER=""
ARG GITPASS=""
ARG GOPRIVATE=gitlab.bianjie.ai
ARG GOPROXY=http://192.168.0.60:8081/repository/go-bianjie/,http://nexus.bianjie.ai/repository/golang-group,https://goproxy.cn,direct

COPY  . $GOPATH/src
WORKDIR $GOPATH/src

# Install minimum necessary dependencies, build binary
# RUN apk add --no-cache $PACKAGES && make all
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && \
    apk add --no-cache $PACKAGES &&  \
    git config --global url."https://${GITUSER}:${GITPASS}@gitlab.bianjie.ai".insteadOf "https://gitlab.bianjie.ai" && make all

FROM alpine:3.15
ENV BINARY_NAME  project-name
WORKDIR /root/
COPY --from=builder /go/src/project-name /usr/local/bin
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories

CMD ["project-name", "start"]
