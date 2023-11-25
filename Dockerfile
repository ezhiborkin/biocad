FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN apt-get update 

RUN go mod download
RUN go build -o mas ./cmd/main.go

CMD ["./mas"]