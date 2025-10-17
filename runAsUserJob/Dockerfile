# Build stage
FROM --platform=$BUILDPLATFORM docker.io/golang:1.24.5-alpine3.21 AS builder
ARG TARGETARCH
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=${TARGETARCH}

WORKDIR /src
#COPY go.mod go.sum ./
COPY ../go.mod ./
RUN go mod download
COPY . .
RUN go build -ldflags="-s -w" -o /build/runAsUserJob

FROM ubuntu:latest
WORKDIR /app
COPY --from=builder /build/runAsUserJob .
USER 1000
ENTRYPOINT ["/app/runAsUserJob"]