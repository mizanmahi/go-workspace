# syntax=docker/dockerfile:1
FROM golang:1.22-alpine AS build

WORKDIR /repo

# Copy workspace files first
COPY go.work go.work.sum ./

# Copy all module source (needed because workspace links them)
COPY types/ ./types/
COPY api/   ./api/

# Build only the api binary
RUN go build -o /out/api ./api/...

# Minimal runtime image
FROM alpine:3.19
COPY --from=build /out/api /api
EXPOSE 8080
ENTRYPOINT ["/api"]