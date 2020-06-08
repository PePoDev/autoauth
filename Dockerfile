# build stage
FROM golang:alpine AS builder
WORKDIR /code
ADD . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -i -a -o /opt/autoauth .

# final stage
FROM alpine:latest
COPY --from=builder /opt/autoauth /opt
RUN mkdir -p /etc/autoauth && echo autoauth: > /etc/autoauth/config.yaml
VOLUME [ "/etc/autoauth" ]
CMD ["sh", "-c","/opt/autoauth start -f /etc/autoauth/config.yaml"]