FROM golang:1.16-stretch as builder

WORKDIR /simple-qa
COPY go.* ./
RUN go mod download -x
COPY . /simple-qa/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

# Main Image
FROM alpine:3.12

RUN mkdir -p /build

COPY --from=builder /simple-qa/qa-api /build/
COPY --from=builder /simple-qa/migrations /build/migrations
COPY --from=builder /simple-qa/config /build/config

WORKDIR /build

CMD ["./qa-api"]