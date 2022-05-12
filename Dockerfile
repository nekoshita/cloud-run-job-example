FROM golang:1.18.1-alpine AS builder
WORKDIR /usr/local/src/
ADD go.mod ./
RUN go mod download
ADD . .
RUN go build -o /go/bin/app ./main.go

FROM alpine:latest
COPY --from=builder /go/bin/app /
CMD ["/app"]
