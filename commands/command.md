Dokumen ini sementara dipertahankan sebagai placeholder.

Struktur aplikasi telah dipindahkan ke pendekatan `cmd/` dan `internal/` agar lebih dekat dengan clean architecture. Jika nanti command generator ingin dihidupkan lagi, sebaiknya gunakan pola berikut:

```text
cmd/
  cli/
internal/
  delivery/cli/
```

Dengan pola itu:

- `cmd/cli` menjadi entry point command line
- `internal/delivery/cli` berisi definisi command Cobra
- business logic tetap dipanggil melalui `usecase`, bukan langsung dari command
