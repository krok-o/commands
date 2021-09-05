FROM golang:latest as builder

WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY main.go .
COPY pkg/ pkg/
COPY cmd/ cmd/
RUN go build -o issue-commenter main.go


FROM ubuntu:latest
LABEL "name"="Issue Commenter Command"
LABEL "maintainer"="Gergely Brautigam <gergely@gergelybrautigam.com>"
LABEL "version"="0.0.1"

WORKDIR /app
COPY --from=builder /build/issue-commenter .

RUN \
  apt-get update && \
  apt-get install -y ca-certificates openssl git && \
  update-ca-certificates && \
  rm -rf /var/lib/apt

ENTRYPOINT ["/app/issue-commenter"]
