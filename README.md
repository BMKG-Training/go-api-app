# go-api-app

Simple Go API Application yang menampilkan data dalam bentuk JSON melalui HTTP request.
Tidak ada koneksi ke database maupun API lain — data disimpan statis di memori
([internal/api/data.go](internal/api/data.go)). Hanya memakai standard library Go.

## Requirements

- Go 1.22+

## Menjalankan

```bash
go run ./cmd/server            # default port 8080
PORT=3000 go run ./cmd/server  # port custom
```

Build binary:

```bash
go build -o bin/server ./cmd/server
./bin/server
```

## Docker

```bash
docker build -t go-api-app .
docker run -d --name go-api-app -p 8080:8080 go-api-app

curl localhost:8080/health
```

Ganti port di dalam container dengan env `PORT`, misalnya
`docker run -p 3000:3000 -e PORT=3000 go-api-app`. Nama host di respons bisa
diatur dengan `--hostname`, misalnya `docker run --hostname api-1 ...`.

Image memakai multi-stage build: binary di-build di image `golang`, lalu hanya
binary-nya yang disalin ke image `alpine` yang kecil.

### Docker Compose

```bash
docker compose up -d --build
docker compose logs -f
docker compose down
```

Konfigurasi service ada di [compose.yaml](compose.yaml).

## Endpoint

| Method | Path                  | Keterangan              |
|--------|-----------------------|-------------------------|
| GET    | `/health`             | Health check            |
| GET    | `/api/v1/items`       | Daftar semua item       |
| GET    | `/api/v1/items/{id}`  | Detail item berdasar ID |

Contoh:

```bash
curl localhost:8080/api/v1/items/1
# {
#   "data": {"id":1,"name":"Item Satu","description":"Contoh data pertama"},
#   "host": {"hostname":"my-server","ip_addresses":["192.168.1.10"]}
# }
```

Setiap respons sukses menyertakan field `host` berisi hostname dan daftar IP
address host yang menjalankan aplikasi (loopback & link-local tidak ditampilkan).

Error dikembalikan dalam format JSON: `{"error": "..."}` dengan status
`400` (ID tidak valid), `404` (data/route tidak ditemukan), atau `405` (method salah).

## Konfigurasi

| Env var | Default | Keterangan       |
|---------|---------|------------------|
| `PORT`  | `8080`  | Port HTTP server |

## Test

```bash
go test ./...
```

## Struktur

```
cmd/server/main.go       entrypoint, konfigurasi server & graceful shutdown
internal/api/handler.go  routing, handler, helper JSON, logging middleware
internal/api/data.go     data statis
internal/api/host.go     info hostname & IP address host
internal/api/*_test.go   test
```
