FROM golang:latest AS builder
WORKDIR /app
COPY . .
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags "-w -s" -o server ./cmd

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/server .
ENTRYPOINT ["./server"]