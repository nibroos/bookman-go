<h1 align="center">
    <br>
  Backend Book Management
  <br>
</h1>

## 🚀 Quick Start
### Developement Environment
On `/` dir, Run `make copy-env`, Modify to suit your environment, focus on these key, you can leave others as it is. The key name is explanatory itself.
```bash
# MAIN APP PORT
GATEWAY_PORT=4050

# DATABASE
POSTGRES_PORT=5432
```

> Without docker, you need to install [air-verse](https://github.com/air-verse/air) to activate the hot reloading.

### 🐳 Docker :: Container Platform

[Docker](https://docs.docker.com/get-docker/) Install.

- On the root folder, Starts the containers in the background and leaves them running : `docker-compose -f docker/docker-compose-dev.yml up --build -d`
- Stops containers and removes containers, networks, volumes, and images : `docker-compose down`

## 🛎 Available Commands each Service

Change bash directory to each service.
> ${arg} means replace all of it match your args without space
- Run export path : `export PATH="$PATH:$(go env GOPATH)/bin"`
- Create mirgration : `make migrate-create name=${your_migration_name}`
- Run migration : `make migrate-up`
- Stepback migraiton: `make migrate-down`
- Generate proto file, leave the proto args blank if you want to generate all proto file: `make proto ${your-proto.proto}`. If its fail, run this command on specific service. for example, in /service/ run bash `export PATH="$PATH:$(go env GOPATH)/bin"`
- Create seeder : `make seed-create name=${your_seeder_name}`
- Run seeder : `make seed-run file=${your_seeder_name}.sql`

## 💎 The Package Features

<p>
  <img src="https://img.shields.io/badge/-Docker-2496ED?style=for-the-badge&logo=Docker&logoColor=fff" />&nbsp;&nbsp;
  <img src="https://img.shields.io/badge/-NGINX-269539?style=for-the-badge&logo=NGINX&logoColor=fff" />
  <img src="https://img.shields.io/badge/-Go-1185F4?style=for-the-badge&logo=Go&logoColor=fff" />
</p>
<p>
<img src="https://img.shields.io/badge/-PostgreSQL-336791?style=for-the-badge&logo=PostgreSQL&logoColor=fff" />&nbsp;&nbsp;
</p>

## 📔 Notes & Issues

#### dial tcp: lookup postgres: no such host
Change the makefile DB_HOST to `localhost` if run in local env, when running on docker, change it to `postgres`.

#### run multiple seeder in one execution
You can run multiple seeder references in the seeder_controller.go file.

#### error running migration fix migration
Change the 'version' column name on schema_migrations to latest succeed migration, change the 'dirty' column to false, then run the migration again

#### postgre extensions
`CREATE EXTENSION IF NOT EXISTS pgcrypto;`

### 📗 API Document
All endpoints stored in  `-.json`


## 📝 Notes
- Projek ini sebagian besar hanya copy dari projek yang sudah saya kerjakan, hanya saya rubah entitasnya, struktur terinspirasi dari laravel & nest js, menggunakan MVC. [Referensi Projek E-Learning](git@github.com:nibroos/bookman-go.git)
- Hanya ada CRUD rest API, autentikasi & otorisasi JWT.
- migrasi dan seeder sudah ada sesuai arahan diatas, kalau gagal, bisa eksekusi manual krn pakai .sql kalau perlu.
- /api/v1/user, /api/v1/author dll, bisa dilihat di routes.go. [Host production elearning](https://api-elearning-service.nibros.tech/api/v1)
- Docker development & deploy manual, belum upload docker registry. developement: docker compose -f docker/docker-compose-dev.yml up -d
- belum ada grpc, unit test, replikasi
- rencananya saya mau implement grpc untuk mengambil list relasi, misal pada list book ada relasi dg author, disitu saya ambil list array untuk diambil dari author service. tapi sebenarnya bisa juga diimplementasikan menggunakan jsonb yg diisi pada saat user input buku baru, lalu ditampilkan tanpa grpc
- stock management, rencananya saya implement seperti in out barang pada umumnya ya, lalu ada stok opname di akhir bulan atau simulasi menggunakan range date
- soal scalling kalau yang dimaksud replikasi server, diprojek lain saya yg struktur nya spt ini sudah bisa dengan docker compose up --scale user-service=2. nanti nya akan otomatis scalling dengan bantuan docker nginx gateway. bisa dilihat nginx.conf
- jujur, saya baru membaca challange nya hari selasa, untuk pengerjaannya pagi 2024-11-26, hari jam 3. jadi mohon maaf krn kesibukan di kantor saya sekarang hasilnya kurang maksimal.
- terima kasih atas kesempatan ini, saya berharap bisa bergabung dengan tim anda, dan saya akan belajar lebih keras lagi.