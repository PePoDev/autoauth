# build stage
FROM golang:alpine AS builder
WORKDIR /go/src/app
ADD . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -i -a -o /go/bin/autoauth .
RUN mkdir -p /etc/autoauth && echo autoauth: > /etc/autoauth/config.yaml

# final stage
FROM gcr.io/distroless/base
COPY --from=builder /go/bin/autoauth /
VOLUME [ "/etc/autoauth" ]
COPY --from=builder /etc/autoauth/config.yaml /etc/autoauth/
CMD ["/autoauth", "start", "-f /etc/autoauth/config.yaml"]