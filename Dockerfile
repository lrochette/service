FROM golang:1.12-alpine3.10 AS builder

# support go modules
RUN apk add --no-cache git
ENV GO111MODULE=on

WORKDIR /app

# cache packages
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o service


FROM alpine:3.10

RUN apk add --no-cache ca-certificates bash
# copy binary
COPY --from=builder /app /

CMD ./service
