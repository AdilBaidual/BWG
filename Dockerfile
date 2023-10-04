FROM golang:1.21.1-bookworm

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN apt-get update

RUN go mod download
RUN go build -o api ./cmd/main.go

EXPOSE 8000

CMD ["./api"]
