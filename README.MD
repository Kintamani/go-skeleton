# Skeleton Project Go 

> **_NOTE:_**  Project ini akan selalu di-update, seiring dengan perubahan teknologi dan arsitektur untuk memudahkan pengembangan aplikasi saya pribadi.

Ini adalah skeleton project Golang yang dirancang dengan prinsip **Clean Architecture** untuk mendukung skalabilitas dan maintainability. Proyek ini juga sudah dilengkapi dengan Docker untuk mempermudah development.

## Struktur Proyek

```plaintext
├── adapters/               # Adapters untuk menghubungkan framework dengan infrastructure
├── api/
│   ├── handlers/           # HTTP Handlers untuk routing
│   ├── middlewares/        # Middlewares untuk memproses request
│   ├── requests/           # DTO untuk payload request
│   ├── responses/          # DTO untuk payload response
│   └── routes/             # Definisi URL routes
│       └── router.go       # File utama untuk mengatur routing
├── bootstrap/              # Bootstrap untuk mengatur konfigurasi
│       ├── app.go          # Bootstrap untuk mengatur konfigurasi aplikasi
│       └── ...
├── commands/               # Command untuk menjalankan aplikasi
├── configs/
│   ├── config.go           # File konfigurasi aplikasi
│   └── db.go               # Setup untuk database
├── databases/
│   ├── migrations/         # Migrations untuk database
│   └── seeders/            # Seeder untuk database
├── docker/                 # File konfigurasi Docker
│   └── configs/
│       ├── app/            # Configurasi docker untuk aplikasi
│       ├── db/             # Configurasi docker untuk database
│       ├── echo/           # Configurasi docker untuk echo
│       ├── env/            # Configurasi docker untuk env
│       ├── log/            # Configurasi docker untuk log
│       ├── validation/     # Configurasi docker untuk validation
│       └── ...
├── docs/                   # Dokumentasi API atau dokumentasi proyek
├── domains/                # Layer bisnis utama
│   └── example/            # Entity dan interfaces example
├── hotreload/              # Direktori untuk output dari build Air hot reload
│   ├── logs/               # Log output dari build Air hot reload
│   └── main                # Binary hasil build
├── infrastructures/        # Infrastruktur seperti repository dan external services
│   ├── persistences/       # Implementasi repository (DB)
│   ├── externals/          # Komunikasi dengan layanan eksternal
│   └── loggers/            # Utilities untuk logging
├── interfaces/             # Interface adapters (controllers)
│   ├── controllers/        # Menghubungkan handler ke services
│   └── presenters/         # Mengelola output data (misal JSON)
│       ├── requests/       # Megelola request dari service `JSON`
│       └── responses/      # Mengelola response dari service `JSON`
├── tests/                  # Direktori untuk unit dan integration tests
├── usecases/               # Use case untuk logika bisnis
├── .ait.toml               # Configurasi Air (hot reload)
├── .dockerignore           # File untuk mengabaikan file dalam build Docker
├── .env                    # Environment variables
├── .env.example            # Environment variables
├── .gitignore              # File untuk mengabaikan file push ke repository
├── docker-compose.yml      # Konfigurasi Docker Compose
├── go.mod                  # File Go module
├── go.sum                  # Dependency checksum
├── LICENCE                 # Lesensi MIT
├── main.go                 # Entry point aplikasi
└── README.md               # Dokumentasi proyek

```

### Persyaratan
- **Golang** (minimal versi 1.23)
- **Docker** dan **Docker Compose**
- **PostgreSQL** sebagai database (dapat dikustomisasi)

### Setup Proyek
1. Clone repository ini:
    ```bash
    git clone https://github.com/Kintamani/go-skeleton
    cd skeleton-project
    ```

2. Install dependencies menggunakan Go:
    ```bash
     go mod download
    ```

3. Atur environment variables: Buat file `.env` dari hasil `.env.example` dan isi dengan konfigurasi yang dibutuhkan.

4. Build dan Jalankan aplikasi dengan Docker:
    ```bash
    docker compose up --build
    ```
    Aplikasi akan berjalan di http://localhost:8080.

5. Struktur Direktorinya mengikuti Clean Architecture:
    - **Domain**: Berisi entitas dan logika bisnis utama.
    - **Infrastructure**: Mengelola persistence dan layanan eksternal.
    - **Usecase**: Berisi skenario bisnis atau use case.
    - **API**: Mengelola handler dan middleware untuk HTTP request.

### Pengujian
Untuk menjalankan unit tests, gunakan perintah berikut:
```bash
go test ./...
```

### Command Penggunaan
* [Migration](databases/migrations/migration.md)
* [Command](commands/command.md)

Ada kesalahan sintaksis dalam kueri, saya telah memperbaikinya untuk Anda. Berikut terjemahannya:

### Aturan Umum untuk Commit
Pada projek menggunakan commit konvensional untuk menangani commit git : [Conventional Commits](https://www.conventionalcommits.org) 
* Gunakan `feat: pesan commit` untuk melakukan commit git yang terkait dengan fitur.
* Gunakan `refactor: pesan commit` untuk melakukan commit git yang terkait dengan refactoring kode.
* Gunakan `fix: pesan commit` untuk melakukan commit git yang terkait dengan perbaikan bug.
* Gunakan `test: pesan commit` untuk melakukan commit git yang terkait dengan file test.
* Gunakan `docs: pesan commit` untuk melakukan commit git yang terkait dengan dokumentasi (termasuk file README.md).
* Gunakan `style: pesan commit` untuk melakukan commit git yang terkait dengan gaya kode.


### Pengembangan
Jika Anda ingin mengembangkan lebih lanjut, pastikan untuk selalu menjaga pemisahan tanggung jawab antara layer-layer dalam arsitektur. Setiap penambahan fitur baru harus mengikuti struktur yang sudah ditentukan untuk menjaga kode tetap bersih dan mudah dimaintain.

### Kontribusi
Jika Anda ingin berkontribusi dalam pengembangan proyek ini, silakan fork repository ini dan buat pull request dengan perubahan yang Anda usulkan.

### Lisensi
Proyek ini dilisensikan di bawah [MIT](./LICENSE) License.