
FROM golang:1.19.3
RUN mkdir /app
WORKDIR /app
COPY . /app
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod download

RUN go build main.go
EXPOSE 8080
CMD ["./main"]
