FROM golang:1.17.1-bullseye
WORKDIR /go/src/learn-github-actions
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ARG BUILD_ID=dev
RUN go install -v -ldflags "-X main.BuildID=${BUILD_ID}"
RUN go test -v
CMD ["learn-github-actions"]