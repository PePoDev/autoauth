# build stage
FROM golang:latest AS builder
ADD . /app/backend
WORKDIR /app/backend
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /autoauth .

# final stage
FROM alpine:latest
COPY --from=builder /autoauth ./
RUN chmod +x ./autoauth
ENTRYPOINT ["./autoauth start"]