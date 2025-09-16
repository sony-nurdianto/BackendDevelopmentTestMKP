# D.Skill Test - Golang API

Proyek ini merupakan implementasi API sederhana menggunakan **Golang** dengan **database Point Test C**.  

API yang dibuat meliputi:

1. **Login API**  
   Menghasilkan **JSON** dengan **JWT token** saat login berhasil.

2. **Create Terminal API**  
   Membuat terminal baru di database. Memerlukan **Authorization** menggunakan **JWT token** dari login API.

---

## Teknologi yang Digunakan

- **Golang** - Bahasa pemrograman backend
- **Database** - Point Test C
- **JWT** - Authentication
- **gRPC** 
- **Postman / curl** - Untuk testing API

---

## How To Run Project

1. cd ./skill_test/service/gRPC && podman compose up -d --build
2. Kemudian Tinggal Akses Enpoint endpointnya


