FROM golang:1.17.1-bullseye
WORKDIR /go/src/learn-github-actions
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go install -v
RUN go test -v
CMD ["learn-github-actions"]