FROM golang:1.18-alpine as builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./
RUN go build ./cmd/server/main.go

FROM alpine:3.14
WORKDIR /
COPY --from=builder /app/main /main
EXPOSE 8080
ENTRYPOINT ["/main"]
