# documentation_odoo — Gin Web Server

> Documentation Portal untuk Titus Dirgantoro, di-serve menggunakan **Go + Gin** dengan struktur **golang-standards/project-layout**.

---

## 📁 Struktur Project

```
documentation_odoo/
├── cmd/
│   └── server/
│       └── main.go             ← Entry point aplikasi
├── configs/
│   └── config.go               ← Konfigurasi dari environment variables
├── internal/
│   ├── handler/
│   │   └── static.go           ← HTTP handlers (Index, Health, NotFound)
│   └── server/
│       └── server.go           ← Setup Gin engine & routing
├── web/                        ← Semua static HTML content
│   ├── index.html
│   ├── public/                 ← Dokumen publik
│   ├── impact/                 ← Dokumen Impact (password protected)
│   └── owner/                  ← Dokumen Owner Only (password protected)
├── go.mod
├── go.sum
├── .env.example
└── README.md
```

> **`internal/`** — package yang tidak bisa di-import dari luar module (Go enforced).  
> **`cmd/`** — entry point binary, minimal dan bersih.  
> **`configs/`** — semua konfigurasi terpusat di sini.  
> **`web/`** — static assets dan HTML, dipisah dari Go source code.

---

## 🚀 Cara Menjalankan

### Development
```bash
go run ./cmd/server/
```

### Build binary (production)
```bash
go build -o doc-server.exe ./cmd/server/
./doc-server.exe
```

Server berjalan di: **http://localhost:8080**

---

## ⚙️ Environment Variables

| Variable   | Default   | Keterangan                    |
|------------|-----------|-------------------------------|
| `PORT`     | `8080`    | Port server                   |
| `GIN_MODE` | `debug`   | Mode Gin: `debug` / `release` |

---

## 🌐 Routes

| Method | Path          | Keterangan                                |
|--------|---------------|-------------------------------------------|
| GET    | `/`           | Halaman utama portal                      |
| GET    | `/health`     | Health check endpoint                     |
| GET    | `/public/*`   | Static files — Public docs                |
| GET    | `/impact/*`   | Static files — Impact docs (🔒 protected) |
| GET    | `/owner/*`    | Static files — Owner-only docs (🔒 protected) |
