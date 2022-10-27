PANDUAN PENGGUNAAN
=========================================================================================================
INSTALASI YANG DIPERLUKAN

1. Install Golang
2. Install Code Editor, contohnya VScode
3. Install Postman
4. Install xampp

Semuanya ada di internet
==========================================================================================================
PEMBUATAN KODE

Membuat Database MyGram / Final
    >> menggunakan mySql, xampp.
        >> digunakan karena mengakses database postgres dari pgAdmin memakan waktu.
        >> config db ada di assignment2
        >> install gorm serta drivernya
            >>"github.com/go-sql-driver/mysql" dan "github.com/jinzhu/gorm"

Membuat Models
    >> Karena ada validasi, import "github.com/asaskevich/govalidator"
    >> struct yang dibuat:
        >>User
        >>Photo
        >>Comment
        >>Socmed
    >> masing-masing dibuat function BeforeCreate untuk validasi

Membuat Koneksi Database
    >> Karena database yang digunakan adalah mysql, contoh koneksi bisa dilihat di assignment2

Membuat Controller dan Middleware

    >>Catatan untuk membuat controller
        >>Mengakses data Photo, Comment, Socmed
            >>perlu Autentikasi - jwt
            >>install github.com/dgrijalva/jwt-go dan golang.org/x/crypto
        >>Memodifikasi (Update dan Delete)
            >>perlu Autorisasi
        >>Autentikasi dan Autorisasi dibuat sebagai Middleware
    
    >>Membuat Middleware mengikuti contoh di sesi 10
    >>Membuat Controller (kebanyakan) mengikuti contoh assignment2 dan sesi 10
        >>import "github.com/gin-gonic/gin"

Membuat Router
    >> menggunakan grouping untuk mempermudah.
    >> ada contoh di sesi 10

Membuat main.go
    >>untuk run
============================================================================================================
POSTMAN

Ketentuan request (header, authorization, body) sudah ditentukan dalam modul, jadi mengikuti saja.
Agar dapat dilihat contohnya, saya simpan collection dari postman nya.
============================================================================================================