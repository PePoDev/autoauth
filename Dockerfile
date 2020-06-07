# build stage
FROM golang:alpine AS builder
ADD . /
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -i -a -o /autoauth /

# final stage
FROM alpine:latest
COPY --from=builder /autoauth /
RUN echo autoauth: > /autoauth.yml
CMD ["sh", "-c","/autoauth start"]