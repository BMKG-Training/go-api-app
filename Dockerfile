# Stage 1: build binary
FROM golang:1.26-alpine AS build
WORKDIR /src
COPY . .
RUN go build -o server ./cmd/server

# Stage 2: image akhir yang kecil, hanya berisi binary
FROM alpine:3.22
WORKDIR /app
COPY --from=build /src/server .
ENV PORT=8080
EXPOSE 8080
CMD ["./server"]
