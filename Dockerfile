FROM golang:1.17.1-bullseye
WORKDIR /go/src/learn-github-actions
COPY . .
RUN go mod download
RUN go install -v
RUN go test -v
CMD ["learn-github-actions"]