# syntax=docker/dockerfile:1

# ---- Build stage ----
FROM golang:1.26-alpine AS build
WORKDIR /src

# Cache module downloads separately from source changes.
COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/server ./cmd/server

# ---- Runtime stage ----
# Distroless static: no shell or package manager, runs as non-root user.
FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=build /out/server /server

ENV PORT=8080
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/server"]
