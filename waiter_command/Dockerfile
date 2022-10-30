FROM golang:latest as builder

WORKDIR /build
COPY go.mod .
RUN go mod download

COPY main.go .
RUN go build -o waiter main.go


FROM ubuntu:latest
LABEL "name"="Waiter Command - Simply waits until the configured period of time."
LABEL "maintainer"="Gergely Brautigam <gergely@gergelybrautigam.com>"
LABEL "version"="0.0.1"

WORKDIR /app
COPY --from=builder /build/waiter .

RUN \
  apt-get update && \
  apt-get install -y ca-certificates openssl git && \
  update-ca-certificates && \
  rm -rf /var/lib/apt

ENTRYPOINT ["/app/waiter"]
