FROM alpine

RUN apk update
RUN apk add go
RUN go install github.com/mikerybka/generate-server@latest

RUN mkdir /app
WORKDIR /app
RUN go mod init app
RUN /root/go/bin/generate-server github.com/mikerybka/storage.Server >> main.go
RUN go mod tidy
RUN go build -o /bin/app main.go

ENTRYPOINT ["/bin/app"]
