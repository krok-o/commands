FROM golang:latest as builder

WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY main.go .
COPY pkg/ pkg/
COPY cmd/ cmd/
RUN go build -o hugo-builder main.go


FROM ubuntu:latest
LABEL "name"="Hugo Builder Command"
LABEL "maintainer"="Gergely Brautigam <gergely@gergelybrautigam.com>"
LABEL "version"="0.0.1"

WORKDIR /app
COPY --from=builder /build/hugo-builder .

RUN \
  apt-get update && \
  apt-get install -y ca-certificates openssl git && \
  update-ca-certificates && \
  rm -rf /var/lib/apt

RUN curl -O -L https://github.com/gohugoio/hugo/releases/download/v0.83.1/hugo_extended_0.83.1_Linux-64bit.deb && dpkg -i hugo_extended_0.83.1_Linux-64bit.deb

ENTRYPOINT ["/app/hugo-builder"]
