FROM golang:latest as builder

WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY main.go .
COPY cmd/ cmd/
RUN go build -o output-command main.go


FROM ubuntu:latest
LABEL "name"="Output Command"
LABEL "maintainer"="Gergely Brautigam <gergely@gergelybrautigam.com>"
LABEL "version"="0.0.2"

WORKDIR /app
COPY --from=builder /build/output-command .

RUN \
  apt-get update && \
  apt-get install -y ca-certificates openssl git && \
  update-ca-certificates && \
  rm -rf /var/lib/apt

ENTRYPOINT ["/app/output-command"]
