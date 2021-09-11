FROM golang:1.17.1-alpine3.14 AS builder
WORKDIR /go/src/learn-github-actions
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ARG BUILD_ID=dev
ENV CGO_ENABLED=0
RUN go install -v -ldflags "-X main.BuildID=${BUILD_ID}"
RUN go test -v

FROM alpine:3.14
COPY --from=builder /go/bin/learn-github-actions /usr/local/bin/learn-github-actions
CMD ["/usr/local/bin/learn-github-actions"]