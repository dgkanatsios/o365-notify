# build stage
FROM golang:1.10.4-alpine3.8 AS builder
ADD . /src
RUN cd /src && go build -o o365-notify

# final stage
FROM alpine:3.8
WORKDIR /app
RUN apk --no-cache add ca-certificates
COPY --from=builder /src/o365-notify /app/
ENTRYPOINT ./o365-notify