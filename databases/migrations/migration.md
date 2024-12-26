### Migration
Untuk mengatur migrasi database, gunakan perintah berikut:

#### Membuat file migration
*untuk sequence number*
```bash
docker compose exec skeleton-golang migrate create -ext sql -dir databases/migrations -seq create_example_table
```

*untuk datetime*
```bash
docker compose exec skeleton-golang migrate create -ext sql -dir databases/migrations create_example_table
```

*atau menggunakan command taskfile (sequence)*
```bash
docker compose exec skeleton-golang task create-migration-sequence name=create_example_table
```

*menggunakan command taskfile (datetime)*
```bash
docker compose exec skeleton-golang task create-migration-datetime name=create_example_table
```

> **_NOTE:_**  jika mengalami error saat menjalankan perintah *sequence* `Next sequence number 20241002005308 too large`, hapus dahulu file migration dengan format *datetime* yang sudah ada di folder `databases/migrations` dan jalankan perintah di atas.

**File migration UP** adalah file yang harus diisi dengan perubahan yang ingin ditambahkan

**File migration DOWN** adalah file yang berisikan kode untuk mengembalikan perubahan yang kita lakukan di file up, karena misal terjadi masalah di aplikasi, namun database migration terlanjur dijalankan, kita bisa melakukan rollback dengan cara menjalankan file down, karena berisikan kode untuk mengembalikan perubahan di file up


#### Menjalankan file migration
```bash
docker compose exec skeleton-golang migrate -path databases/migrations -database "postgres://DB_USERNAME:DB_PASSWORD@DB_HOST:DB_PORT_CONTAINER/DB_NAME?sslmode=disable" up
```
*atau menggunakan command taskfile*
```bash
docker compose exec skeleton-golang task migrate
```

contoh :

```bash
docker compose exec skeleton-golang migrate -path databases/migrations -database "postgres://skeleton:skeleton@skeleton-db:5432/skeleton_db?sslmode=disable" up
```