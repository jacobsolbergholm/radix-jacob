# Build stage
FROM --platform=$BUILDPLATFORM docker.io/golang:1.24.5-alpine3.21 AS builder
ARG TARGETARCH
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=${TARGETARCH}

WORKDIR /src
#COPY go.mod go.sum ./
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -ldflags="-s -w" -o /build/radix-jacob

# Final stage, ref https://github.com/GoogleContainerTools/distroless/blob/main/base/README.md for distroless
FROM gcr.io/distroless/static
WORKDIR /app
COPY --from=builder /build/radix-jacob .
USER 1000
ENTRYPOINT ["/app/radix-jacob"]