package migrations

import (
	"github.com/jinzhu/gorm"
	"github.com/mahfuz05/go-banking-restapi/helpers"
)

type User struct {
	gorm.Model
	UserName string
	Email    string
	Password string
}

type Account struct {
	gorm.Model
	Type    string
	Name    string
	Balance uint
	UserID  uint
}

func connectDB() *gorm.DB {
	db, err := gorm.Open("postgress", "host=127.0.0.1 port=5432 user=user dbname=dbname password=password sslmode=disable")
	helpers.HandleError(err)
	return db
}

func createAccounts() {
	db := connectDB()

	users := [2]User{
		{UserName: "Mahfuz", Email: "mahfuzcse05@gmail.com"},
		{UserName: "Kanak", Email: "m@m.com"},
	}

	for i := 0; i < len(users); i++ {
		generatePassword := helpers.HashAndSalt([]byte(users[i].UserName))
		users[i].Password = generatePassword
		db.Create(&users[i])

		account := Account{
			Type:    "Daily Account",
			Name:    string(users[i].UserName + "'s Accounts"),
			Balance: uint(1000 * int(i+1)),
			UserID:  users[i].ID,
		}
		db.Create(&account)
	}
	defer db.Close()
}

func Migrate() {
	db := connectDB()
	db.AutoMigrate(&User{}, &Account{})

	defer db.Close()
	createAccounts()
}
