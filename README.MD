# Go Skeleton

Skeleton project Go untuk membangun REST API dengan struktur yang mengarah ke **Clean Architecture**. Repo ini sudah menyiapkan fondasi dasar: HTTP server dengan Echo, konfigurasi environment, koneksi PostgreSQL, migration, seeder, logging, dan Docker untuk development.

## Status

Struktur proyek ini masih berupa skeleton yang terus disederhanakan dan dirapikan. Boundary antar layer sudah mulai dibentuk, tetapi belum selengkap implementasi clean architecture yang matang. Karena itu, dokumentasi ini dibuat untuk mencerminkan kondisi repo saat ini, bukan struktur ideal yang belum sepenuhnya ada.

## Tech Stack

- Go `1.22+`
- Echo
- PostgreSQL
- SQLX
- Cobra
- Logrus
- Docker Compose

## Struktur Proyek

```text
.
├── adapters/              # Adapter untuk framework, logger, validator, database, dan CLI
├── api/                   # Routing dan response HTTP
├── bootstrap/             # Wiring aplikasi saat startup
├── commands/              # Eksperimen/generator command berbasis Cobra
├── config/                # Loader environment dan konfigurasi aplikasi
├── databases/
│   ├── migrations/        # File migration SQL
│   └── seeders/           # Seeder database
├── docker/
│   └── configs/           # Dockerfile untuk Go app dan PostgreSQL
├── docs/
│   └── swagger/           # Dokumen Swagger/OpenAPI dasar
├── domains/               # Domain dan placeholder entitas bisnis
├── hotreload/             # Output binary hot reload
├── main.go                # Entry point aplikasi HTTP
├── docker-compose.yml
└── taskfile.yml           # Shortcut task untuk migration
```

## Alur Aplikasi

Alur request saat ini sederhana:

1. `main.go` memuat environment, logger, validator, Echo, dan database.
2. `bootstrap/app.go` melakukan wiring dependency aplikasi.
3. `api/routes/api.go` mendaftarkan route.
4. Response error diformat oleh `api/responses/app.go`.

Endpoint bawaan yang tersedia saat ini:

```http
GET /api/v1/ping
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

Kalau dibandingkan dengan referensi seperti `khannedy/golang-clean-architecture`, repo itu lebih tegas dalam memisahkan delivery, use case, repository, dan gateway. Arah itu bagus, tetapi untuk repo ini lebih aman dilakukan bertahap:

- pertahankan folder yang sudah aktif dipakai
- rapikan naming agar konsisten
- pisahkan use case dan infrastructure saat kebutuhan bisnis mulai nyata
- hindari menambahkan banyak folder kosong hanya demi terlihat "clean architecture"

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
