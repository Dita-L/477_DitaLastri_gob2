package main

import (
	"fmt"
	"gorm/database"
	"gorm/models"
)

func main() {
	database.StartDB()

	//createUser("johndoe@gmail.com")
	//getUserByID(1)
	//updateUserByID(1, "johnjon@gmail.com")
	//createProduct(1, "YLO", "YYYY")
	//getUsersWithProducts()
	deleteProductById(1)
}

/*func createUser(email string) {
	db := database.GetDB()
	//mengandung referensi database

	User := models.User{
		Email: email,
	}

	err := db.Create(&User).Error
	//untuk melakukan create data, gunakan method Create
	//method Create memerlukan data sebagai argumen
	//data tersebut harus memiliki tipedata yang sama dengan yang ingin dibuat
	//kita ingin membuat user, makadari itu data yang diberikan pada method Create adalah data yang bertipe struct User.

	// method untuk CRUD memiliki properti Error.

	if err != nil {
		fmt.Println("Error creating user data:", err)
		return
	}

	fmt.Println("New User Data:", User)
}*/

/*func getUserByID(id uint) {
	db := database.GetDB()

	user := models.User{}

	err := db.First(&user, "id = ?", id).Error
	//Untuk mendapatkan data dari tabel, dapat menggunakan method First.
	//method First menerima 3 parameter. parameter1 adalah pointer terhadap yang dicari.
	//parameter ke 2 adalah condition dari query
	//parameter 3 adalah data dari conditionnya

	//method First akan mengemballikan ErrRecordNotFound jika tidak ada data yang ditemukan.

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("User data not found")
			return
		}
		print("Error finding user:", err)
	}

	fmt.Printf("User Data: %+v \n", user)

}*/

/*func updateUserByID(id uint, email string) {
	db := database.GetDB()

	user := models.User{}

	err := db.Model(&user).Where("id = ?", id).Updates(models.User{Email: email}).Error
	//method model digunakan terlebih dahulu agar hasil update dapat langsung di scan
	//untuk membuat conditionnya, gunakan method Where.
	//selain menggunakan method Model, kita juga bisa menggunakan method Table sbb
	// db.Table("users").Where("id = ?").Updates(map[string]interface{}{"email" : "johnjohn@gmail.com"})

	if err != nil {
		fmt.Println("Error updating user data:", err)
		return
	}
	fmt.Printf("Update user's email %+v \n", user.Email)
}*/

/*func createProduct(userID uint, brand string, name string) {
	db := database.GetDB()

	Product := models.Product{
		UserID: userID,
		Brand:  brand,
		Name:   name,
	}

	err := db.Create(&Product).Error

	if err != nil {
		fmt.Println("Error creating product data:", err.Error())
		return
	}
	fmt.Println("New Product Data:", Product)

}*/

/*func getUsersWithProducts() {
	db := database.GetDB()

	users := models.User{}
	err := db.Preload("Products").Find(&users).Error
	//untuk melakukan eager loading, dapat menggunakan method Preload
	// method preload dengan parameter nama table
	// walaupun nama table seharusnya lowercase semua, untuk parameter preload harus digunakan uppercase

	if err != nil {
		fmt.Println("Error getting user datas with products:", err.Error())
		return
	}

	fmt.Println("User Datas With Products")
	fmt.Printf("%+v", users)
}*/

func deleteProductById(id uint) {
	db := database.GetDB()

	product := models.Product{}
	err := db.Where("id = ?", id).Delete(&product).Error

	if err != nil {
		fmt.Println("Error deleting product:", err.Error())
		return
	}

	fmt.Printf("Product with id %d has been successfully deleted", id)
}
