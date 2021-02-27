FROM golang:1.15 as build
LABEL maintainer="117503445"
ENV GO111MODULE=on
WORKDIR /root/project
ADD go.mod .
ADD go.sum .
RUN go mod download
ADD . .
RUN wget https://goframe.org/cli/linux_amd64/gf && chmod +x gf && ./gf install
RUN gf swagger --pack -y
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -installsuffix cgo -o server_bin
FROM alpine as prod
EXPOSE 80
COPY --from=build /root/project/server_bin /root/server_bin
WORKDIR /root
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' >/etc/timezone
ENTRYPOINT /root/server_bin