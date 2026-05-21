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
│   └── web/               # Entry point aplikasi HTTP
├── config/                # Loader environment dan konfigurasi aplikasi
├── databases/
│   ├── migrations/        # File migration SQL
│   └── seeders/           # Seeder database
├── docker/
│   └── configs/           # Dockerfile untuk Go app dan PostgreSQL
├── docs/
│   └── swagger/           # Dokumen Swagger/OpenAPI dasar
├── internal/
│   ├── app/               # Wiring dependency aplikasi
│   ├── delivery/          # Delivery layer, route, handler, HTTP response
│   ├── entity/            # Entity domain
│   ├── infrastructure/    # Implementasi framework, DB, logger, repository
│   ├── model/             # DTO untuk output/input use case
│   ├── repository/        # Kontrak repository
│   └── usecase/           # Business use case
├── hotreload/             # Output binary hot reload
├── docker-compose.yml
└── taskfile.yml           # Shortcut task untuk migration
```

## Alur Aplikasi

Alur request saat ini sederhana:

1. `cmd/web/main.go` memuat environment dan menjalankan aplikasi.
2. `internal/app/web.go` melakukan wiring dependency aplikasi.
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

Detail tambahan ada di [databases/migrations/migration.md](databases/migrations/migration.md).

## Seeder

Seeder disiapkan di folder `databases/seeders`. Saat ini contoh seeder tersedia untuk tabel `examples`.

## Testing

Menjalankan seluruh test:

```bash
go test ./...
```

## Catatan Arsitektur

Struktur repo ini sekarang sudah lebih dekat dengan referensi `khannedy/golang-clean-architecture`:

- entry point dipindah ke `cmd/`
- business flow dipusatkan ke `internal/usecase`
- kontrak repository dipisah dari implementasinya
- HTTP handler hanya menangani delivery concern
- implementasi DB, logger, dan HTTP server ada di `internal/infrastructure`

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
