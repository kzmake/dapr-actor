FROM golang:1.17 as builder

WORKDIR /go/src/app

COPY go.mod .
COPY go.sum .
RUN go mod download

ARG SERVICE_NAME=actor
ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64
COPY api api
COPY $SERVICE_NAME $SERVICE_NAME
RUN go build -o /go/bin/app -ldflags '-s -w' $SERVICE_NAME/main.go

FROM scratch as runner

COPY --from=builder /go/bin/app /app

CMD ["/app"]
