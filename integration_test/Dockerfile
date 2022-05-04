FROM golang:latest as builder

WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY main.go .
COPY cmd/ cmd/
RUN go build -o integration-test main.go


FROM ubuntu:latest
LABEL "name"="Integration Test Command"
LABEL "maintainer"="Gergely Brautigam <gergely@gergelybrautigam.com>"
LABEL "version"="0.0.2"

WORKDIR /app
COPY --from=builder /build/integration-test .

RUN \
  apt-get update && \
  apt-get install -y ca-certificates openssl git && \
  update-ca-certificates && \
  rm -rf /var/lib/apt

ENTRYPOINT ["/app/integration-test"]
