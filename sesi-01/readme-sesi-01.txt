Di sesi 1 kita belajar

1. installasi Golang dan VSCode
  
2. Menulis kode pertama
	package main
	
	import "fmt"
	
	func main() {
		fmt.Println("Hello World")
	}
	
  3. Menyiapkan GOPATH dan Go module
	
  4. Tentang fmt.Println dan fmt.Print
		fmt.Println, hasil kode => dengan spasi per variabel yang ditulis
		fmt.Print, hasil kode => tanpa spasi per variabel yang ditulis
	
  5. Membuat Variabel
		variabel, seperti dalam matematika, gunanya untuk menyimpan suatu nilai
		ada tiga(3) cara membuat variabel
			a. var nama_variabel data_type = sesuatu		(tipe data langsung ditetapkan)
			b. var nama_variabel = sesuatu				(tipe data ditetapkan secara implisit)
			c. nama_variabel := sesuatu
		untuk membuat lebih dari satu variabel dalam satu baris kode,
			var variabel1, variabel2, variabel3 = sesuatu1, sesuatu2, sesuatu3
		nama variabel hanya bisa terdiri dari huruf, angka, dan underscore (_). selain itu, nama variabel tidak boleh diawali oleh angka.
	
  6. Tipe Data
		Tipe data mewakili nilai dalam variabel,
			a. int dan uint
				int dan uint adalah angka (number), bedanya adalah int untuk nilai negatif maupun positif, sedangkan int untuk positif saja.
			b. float
				float adalah tipe data untuk angka juga, tapi lebih spesifik ke desimal.
			c. bool
				bool adalah tipe data yang hanya mempunyai 2 nilai, yaitu true dan false.
			d. string
				string adalah tipe data untuk kata ataupun kalimat. ditandai dengan penggunaan tanda kutip ("").
				penggunaan string bisa juga menggunakan backtick(`), gunanya agar simbol-simbol dalam kalimat tidak disalah artikan sebagai kode.
				
					var contoh1 string = " "hai", saya berkata "
					fmt.Println(contoh1)

					var contoh2 string = `"hai", saya berkata "
					fmt.Println(contoh2)
			e. nill dan zero value
				zero value semacam kondisi default suatu tipedata. yaitu ketika suatu variabel telah ditentukan tipe data yang digunakan, tetapi tidak diberi nilai.
					var kata string
					var benarsalah bool
					var nonDesimal int
					var Desimal uint
					fmt.Println(kata, benarsalah, nonDesimal, Desimal)

				output / hasil :
					  false 0 0 =======> (karakter kosong) false 0 0
	
  7. Const
		const merupakan jenis variabel. const merupakan kepanjangan dari constant, yang berarti tidak berubah. artinya, nilai dari const tidak bisa diubah, dan harus langsung ditetapkan ketika const dibuat.
		membuat const sama seperti variabel
			contoh benar
				const pi = 3.14
			contoh salah
				const lamda 		
	
  8. Penggunaan fmt.Printf
		seperti fmt.Println, fmt.Print, fmt.Printf juga berfungsi mencetak/menampilkan nilai. bedanya, fmt.Printf seperti perintah "print dengan format ini".

			var nama = "Dita"
			var usia = 24

			fmt.Printf("Hai, nama saya %s, usia saya %d tahun", nama, usia)

		yang terjadi disini adalah,
			print "Hai, nama saya (nilai string dari variabel nama), usia saya (nilai int dari variabel usia) tahun"
			
			%s adalah kode untuk tipe data string
			%d adalah kode untuk tipe data int

		output pada terminal,
			Hai, nama saya Dita, usia saya 24 tahun
