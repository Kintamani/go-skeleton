Pindah ke direktori myapp/.

Jalankan perintah go build untuk command tertentu:

```bash
go build -o myapp-makefile cmd/makefile.go
```
Ini akan menghasilkan binary bernama myapp-makefile. Anda bisa menjalankannya seperti aplikasi biasa:
```bash
./myapp-makefile make:file MyNewFile
```
Dengan menjalankan perintah ini, file binary akan mengeksekusi command make:file dan menghasilkan file Go baru sesuai dengan nama yang Anda berikan (MyNewFile.go).