package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	//ada _ karena kita hanya meregistrasi driver Postgresql
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Dita2710"
	dbname   = "db_go"
)

var (
	db  *sql.DB
	err error
)

//variabel global, agar dapat digunakan global, melalui berbagai function

type Employee struct {
	ID        int
	Full_name string
	Email     string
	Age       int
	Division  string
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	//psqlInfo menggabungkan seluruh data postgreSQL menjadi string panjang.
	db, err = sql.Open("postgres", psqlInfo)
	//function Open yang berasal daripackage database/sql, menerima 2 parameter.
	//parameter 1 berupa nama driver yang digunakan ("postgres"),
	//parameter 2 berupa string, yang mengandung informasi tentang bagaimana cara database/sql membangun koneksinya kepada database, sesuai dengan driver yang kita gunakan
	//function Open mengembalikan nilai berupa pointer dari struct sql.DB, struct tersebut memiliki method Ping
	//function Open berfungsi untuk memvalidasi argumen yang diberikan.
	//anggap sebagai proses login.

	if err != nil {
		panic(err)
	}
	defer db.Close()
	//kalau error, close.

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	//dengan memanggil method Ping, kita dapat membangun koneksi ke database
	//sekaligus memeriksa jika string panjang psqlInfo 100%valid

	fmt.Println("Successfully connected to database")

	//CreateEmployee()
	//GetEmployees()
	//UpdateEmployee()
	DeleteEmployee()
}

func CreateEmployee() {
	var employee = Employee{}

	sqlStatement := `
	INSERT INTO employees (full_name, email, age, division)
	VALUES ($1, $2, $3, $4)
	returning *
	`
	//sqlStatement berisi sebuah statement untuk melakukan "create data" pada database.
	//ingat saat baru membuat database, mirip seperti itu.
	//$1,$2,... merupakan representasi dari nilai-nilai yang akan dimasukan nantinya.
	//"returning *" memiliki arti kita ingin mendapat nilai-nilai dari seluruh field yang berasal dari data yang baru dibuat.

	err = db.QueryRow(sqlStatement, "Airell Jordan", "airell@mail.com", 23, "IT").
		Scan(&employee.ID, &employee.Full_name, &employee.Email, &employee.Age, &employee.Division)
	//method QueryRow digunakan untuk mengeksekusi query sql.
	//karena sqlStatement bertujuan untuk create data, maka QueryRow berfungsi untuk membuat data baru dengan nilai yang kita berikan pada parameter 2.
	//parameter ke-2 ini merupakan parameter variadic, dan niilai2 yang kita berikan adalah nilai-nilai yang akan mereplace statement yang kita buat pada VALUES ($1, $2,..)
	//method QueryRow kita chain dengan method Scan.
	//hal ini dilakukan agar kita dapat mendapatkan nilai-nilai balikan dari statement yang telah kita buat.
	//nilai-nilai tersebut akan kita simpan dalam field-field dari variabel employee (var employee Employee{})

	if err != nil {
		panic(err)
	}

	fmt.Printf("New Employee Data : %+v\n", employee)

}

func GetEmployees() {
	var results = []Employee{}

	sqlStatement := `SELECT * from employees`

	rows, err := db.Query(sqlStatement)
	//method Query biasa digunakan untuk mendapatkan banyak data dari suatu table pada database
	//method ini dapat mengembalikan satu atau lebih rows dari suatu table pada database.
	//rows harus ditutup dengan method Close

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		//kita melakukan looping sebanyak data yang kita dapatkan dengan acuan rows.Next.
		//rows.Next akan menghasilkan nilai true selama datanya masih ada.
		//jika sudah tidak ada, maka akan me-return false, dan looping terhenti.

		var employee = Employee{}

		err = rows.Scan(&employee.ID, &employee.Full_name, &employee.Email, &employee.Age, &employee.Division)
		//Scan dilakukan agar kita dapat mendapatkan nilai-nilai balikan dari statement yang telah kita buat.

		if err != nil {
			panic(err)
		}

		results = append(results, employee)
		//setelah proses scan untuk satu row selesai, data dari row tersebut akan kita masukkan dalam variabel results.

	}

	fmt.Println("Employee datas:", results)
}

func UpdateEmployee() {
	sqlStatement := `
	UPDATE employees
	SET full_name = $2, email = $3, division = $4, age = $5
	WHERE id = $1;`
	res, err := db.Exec(sqlStatement, 1, "Airell Jordan Hidayat", "airellhidayat@gmail.com", "CurDevs", 24)
	//kita dapat menggunakan method QueryRow untuk mengupdate data, namun dianjurkan menggunakan method Exec untuk proses insert, update, dan delete.
	//dengan method Exec, kita tidak mendapatkan data yang baru saja diupdate maupun dibuat.
	//kita dapat menggunakan method RowsAffected untuk mengetahui jumlah row atau data yang baru diupdate.

	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()
	//method RowsAffected didapatkan dari interface sql.Result yang merupakan nilai return dari method Exec.

	if err != nil {
		panic(err)
	}

	fmt.Println("Update data amount:", count)
}

func DeleteEmployee() {
	sqlStatement := `
	DELETE from employees
	WHERE id = $1;
	`

	res, err := db.Exec(sqlStatement, 1)
	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("Deleted data amount:", count)
}
