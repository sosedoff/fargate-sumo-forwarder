# ------------------------------------------------------------------------------
# Build Phase
# ------------------------------------------------------------------------------
FROM golang:1.14 AS build

ADD . /go/src/github.com/sosedoff/fargate-sumo-forwarder
WORKDIR /go/src/github.com/sosedoff/fargate-sumo-forwarder

RUN \
  GOOS=linux \
  GOARCH=amd64 \
  CGO_ENABLED=0 \
  go build -o /fargate-sumo-forwarder

# ------------------------------------------------------------------------------
# Package Phase
# ------------------------------------------------------------------------------

FROM alpine:3.6

RUN \
  apk update && \
  apk add --no-cache ca-certificates openssl wget && \
  update-ca-certificates

WORKDIR /app

COPY --from=build /fargate-sumo-forwarder /bin/fargate-sumo-forwarder

EXPOSE 5000
CMD ["/bin/fargate-sumo-forwarder"]