CONDITIONS
    Kondisi(condition) digunakan untuk mengatur alur suatu program. analogikan seperti lampu lalulintas, ketika kondisi lampu berwarna merah kita berhenti, ketika lampunya hijau kita jalan. ini yang dimaksud mengatur alur.
    Pada Go, terdapat 2 cara menggunakan kondisi yaitu memakai "if", dan "switch".
    "if" dan "switch" bisa digabungkan menjadi "nested conditions".

LOOPING
    Looping atau perulangan merupakan suatu proses berulang yang dimana proses tersebut akan berhenti jika memenuhisuatu kondisi.
    Bahasa Go hanya memiliki satu looping yaitu looping dengan menggunakan keyword "for" atau yang kita kenal dengan sebutan "for loop".
    Pastikan saat membuat loop jangan sampai membuat infinity loop. hal itu akan membuat program crash.

ARRAY
    Array adalah kumpulan dari data-data. Anggap saja himpunan. Data-data dalam array haruslah bertipe sama. Array pada go memiliki panjang yang tetap. Panjang ini menentukan jumlah data yang ada pada Array.
    
    Membuat array ada 2 cara, 
		    a. Membuat Array Kosong

			      var nama_array [panjang_array/jumlah_data]tipe_data
            
        b. Membuat array langsung terisi
        
            var nama_array = [panjang_array/jumlah_data]tipe_data{data1, data2, data3, ....}
            
    Modify Data Through Index (Mengubah data melalui indeks)
        Kita dapat mengakses dan mengubah data pada array melalui indeks.
            nama_array[no.indeks]
    
    Looping pada array
        Looping pada array ada dua cara, yaitu menggunakan range dan len.
        
SLICE
    Slice merupakan suatu tipe data yang mirip dengan tipe data array, yaitu untuk menyimpan satu atau lebih data. Berbeda dengan array, slice tidak memiliki sifat fixed-length  yang berarti panjang dari slice tidak tetap. Berarti, ketika membuat slice, jumlah elemen/data yang ada dalam slice tidak ditentukan.
    
    var nama_array = []tipe_data{data1, data2, data3,....}
    
    var fruit = []string{"apel", "pisang", "jambu"}
    
Membuat slice juga bisa menggunakan "make". Hasilnya akan berupa slice kosong.
      
    var angka = make([]int, 3) =======> [0 0 0]
    
Jika kita ingin mengubah elemen pada slice, sama seperti array, kita dapat menggunakan indeksnya.

      nama_slice[no.indeks]
      
Ada juga fungsi "append", digunakan untuk menambahkan elemen pada slice, ataupun menggabungkan dua slice atau lebih.

Jika kita ingin menyalin slice1 ke slice2, gunakan fungsi "copy". fungsi "copy" akan mengembalikan nilai berupa banyaknya jumlah elemen yang di-copy, yaitu nilai paling rendah dari "len" masing-masing slice.

Ada cara agar kita dapat mendapatkan element-element dari sebuah slice dan kita juga bisa menentukan element dari index keberapa yang ingin kita dapatkan. 
Caranya adalah dengan menggunakan slicing.

Setiap kita membuat suatu slice pada bahasa Go , secara otomatis Go akan membuat suatu array tersembunyi yang disebut dengan Backing array. Backing array akan bertugas untuk menyimpan elemen pada slice, bukan slice nya sendiri. Bahasa Go mengimplementasikan slice sebagai sebuah struktur data yang disebut dengan slice header.
menggunakan append dapat mambuat backing array baru.

            

    
