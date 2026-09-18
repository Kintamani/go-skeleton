# Go Skeleton

Skeleton project Go untuk membangun REST API dengan struktur yang mengarah ke **Clean Architecture**. Repo ini sudah menyiapkan fondasi dasar: HTTP server dengan Echo, konfigurasi environment, koneksi PostgreSQL, migration, seeder, logging, dan Docker untuk development.

## Status

Struktur proyek ini masih berupa skeleton yang terus disederhanakan dan dirapikan. Boundary antar layer sudah mulai dibentuk, tetapi belum selengkap implementasi clean architecture yang matang. Karena itu, dokumentasi ini dibuat untuk mencerminkan kondisi repo saat ini, bukan struktur ideal yang belum sepenuhnya ada.

## Tech Stack

- Go `1.22+`
- Echo
- PostgreSQL
- SQLX
- Logrus
- Docker Compose

## Struktur Proyek

```text
.
├── cmd/
│   ├── api/               # Entry point aplikasi HTTP
│   ├── scheduler/         # Placeholder executable untuk scheduler
│   ├── seed/              # Executable untuk menjalankan seed
│   └── worker/            # Placeholder executable untuk background worker
├── db/
│   ├── migrations/        # File migration SQL
│   └── seeders/           # Seeder database
├── deployments/
│   ├── docker/            # Dockerfile dan bootstrap container
│   ├── helm/              # Placeholder chart Helm
│   ├── kubernetes/        # Placeholder manifest Kubernetes
│   └── terraform/         # Placeholder infrastructure as code
├── docs/
│   └── swagger/           # Dokumen Swagger/OpenAPI dasar
├── internal/
│   ├── app/               # Wiring dependency aplikasi
│   ├── config/            # Loader environment dan konfigurasi aplikasi
│   ├── delivery/          # HTTP, gRPC, websocket, dan worker delivery
│   ├── entity/            # Entity domain
│   ├── infrastructure/    # Implementasi framework, DB, logger, queue, redis
│   ├── mapper/            # Placeholder mapper lintas layer
│   ├── model/             # Model input/output use case yang aktif dipakai
│   ├── repository/        # Repository contract dan implementasi aktif
│   └── usecase/           # Business use case
├── pkg/                   # Placeholder shared package
├── scripts/               # Placeholder helper scripts
├── docker-compose.yml
└── Taskfile.yml           # Shortcut task untuk migration
```

## Alur Aplikasi

Alur request saat ini sederhana:

1. `cmd/http/main.go` memuat environment dan menjalankan executable HTTP server (atau `cmd/grpc/main.go` untuk gRPC server).
2. `internal/app/http.go` melakukan wiring dependency aplikasi HTTP.
3. `internal/delivery/http/route/route.go` mendaftarkan route.
4. Handler memanggil use case.
5. Use case memanggil repository interface.
6. Implementasi repository ada di layer `internal/infrastructure`.

Endpoint bawaan yang tersedia saat ini:

```http
GET /api/v1/ping
GET /api/v1/examples
```

Response:

```text
pong
```

## Menjalankan Proyek

### Prasyarat

- Go `1.22+`
- Docker dan Docker Compose

### Setup

1. Clone repository:

```bash
git clone https://github.com/Kintamani/go-skeleton.git
cd go-skeleton
```

2. Siapkan environment:

```bash
cp .env.example .env
```

3. Jalankan service:

```bash
docker compose up --build
```

Secara default aplikasi akan berjalan di `http://localhost:8181`.

File `docker-compose.yml` tetap berada di root agar workflow local development tetap sederhana, sedangkan Dockerfile dan bootstrap container disimpan di `deployments/docker`.

## Migration

Migration bisa dijalankan langsung dengan `migrate` atau melalui `task`.

Menjalankan migration:

```bash
docker compose exec skeleton-golang task migrate
```

Membuat migration dengan sequence number:

```bash
docker compose exec skeleton-golang task create-migration-sequence name=create_examples_table
```

Membuat migration dengan timestamp:

```bash
docker compose exec skeleton-golang task create-migration-datetime name=create_examples_table
```

Detail tambahan ada di [db/migrations/migration.md](/Users/camel/Projects/Go/go-skeleton/db/migrations/migration.md).

Contoh file migration yang bisa dijadikan acuan:

- `202605210001_create_roles.up.sql`
- `202605210002_create_users.up.sql`

## Seeder

Seeder disiapkan di folder `db/seeders/`. Contoh yang saat ini tersedia:

- `example_seeder.go`
- `role_seeder.go`
- `admin_seeder.go`

Menjalankan example seeder:

```bash
go run ./cmd/seed -seed example -total 10
```

Menjalankan role seeder:

```bash
go run ./cmd/seed -seed role
```

Menjalankan admin seeder:

```bash
go run ./cmd/seed -seed admin
```

`admin` seeder akan memastikan role dasar tersedia lebih dulu, jadi aman dijalankan langsung pada database kosong.

Mengosongkan data example:

```bash
go run ./cmd/seed -seed clear-example
```

## Testing

Menjalankan seluruh test:

```bash
go test ./...
```

## Catatan Arsitektur

Struktur repo ini sekarang sudah lebih dekat dengan referensi `khannedy/golang-clean-architecture`:

- entry point dipindah ke `cmd/`
- scaffold executable lain sudah disiapkan di `cmd/` untuk base project
- business flow dipusatkan ke `internal/usecase`
- `internal/config/` menjadi satu-satunya lokasi config aktif
- `internal/model/` dipakai untuk model input/output use case yang aktif
- repository sekarang disederhanakan di satu area `internal/repository`
- HTTP handler hanya menangani delivery concern
- implementasi DB, logger, dan HTTP server ada di `internal/infrastructure`
- file output build dan cache development tidak lagi menjadi bagian struktur source

Masih ada ruang pengembangan, misalnya:

- menambah request/response model per fitur yang lebih lengkap
- memisahkan transaction boundary bila use case bertambah kompleks
- menambah unit test di layer use case dan repository

## Commit Convention

Proyek ini mengikuti [Conventional Commits](https://www.conventionalcommits.org/).

- `feat:` untuk fitur baru
- `fix:` untuk perbaikan bug
- `refactor:` untuk perubahan struktur internal
- `test:` untuk test
- `docs:` untuk dokumentasi
- `style:` untuk perubahan formatting/gaya

## Lisensi

Proyek ini menggunakan lisensi [MIT](./LICENSE).
