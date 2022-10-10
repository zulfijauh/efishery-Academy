# Description of Program

Program ini merupakan rancangan untuk membuat order management system (ecommerce) untuk melakukan transaksi jual beli produk yang terintegrasi dengan database perusahaan sehingga dapat membuat, mengubah, mengambil dan menghapus data dari table user(customer) dan product. Dengan program ini juga user dapat melakukan pembelian produk dengan menambahkan produk ke dalam cart dan kemudian melakukan check-out


## Problem and Motivation

Program dibuat untuk mempermudah transaksi jual beli produk antara supplier dan customer secara daring. Dengan system ini, setiap data customer, produk dan transaksi yang terjadi akan tersimpan kedalam database untuk memudahkan perusahaan untuk mengontrol dan memonitor aktivitas transaksi.

## How to run
1. Download project ini dari https://github.com/zulfijauh/efishery-Academy atau dengan mengclone repo ini dengan cara:
```bash
git clone https://github.com/zulfijauh/efishery-Academy
```
Delete seluruh folder kecuali folder "assessment" dimana project ini berada

2. Setup configuration
- Buatlah database (atau bisa gunakan database yang sudah tersedia) pada PostgreSQL
- Buka file config, kemudian pada variable dsn, silahkan rubah sesuai path database yang tadi sudah di buat
```bash
dsn := "host=localhost user=user_name password=password dbname=database_name search_path=schema_jika_ada port=5432 TimeZone=Asia/Jakarta sslmode=disable"
```
- jalankan command berikut:
```bash
go mod download
```
- jika ada package yang belum terdownload, download dengan menjalankan command "go get" contoh:
```bash
 go get -u gorm.io/gorm
```
- untuk menjalankan program, jalankan command:
```bash
 go run .
```
- program anda sudah terhubung dengan database dan bisa melakukan input dan perubahan dangan dengan postman

## How to use
Untuk melakukan CRUD RestAPI dengan mengukan postman, silahkan lihat endpoint yang tersedia pada folder routes, atau import file collection.json di postman untuk lebih detail
