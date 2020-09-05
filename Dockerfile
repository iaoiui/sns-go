FROM golang:alpine as builder
WORKDIR /go
COPY ./vendor/ ./vendor/
COPY go.mod go.sum ./
ENV GOPATH=""
COPY ./app ./app/
RUN go build -o main app/*

FROM alpine:latest
ENV GOTRACEBACK=single
COPY --from=builder /go/main .
CMD ["./main"]
