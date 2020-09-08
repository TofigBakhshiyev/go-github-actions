FROM golang:1.13.6-alpine3.10
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go clean --modcache
RUN go build -o main .
CMD ["/app/main"]