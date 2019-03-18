FROM golang:1.12-alpine

ADD . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o server .

ENTRYPOINT [ "./server" ]