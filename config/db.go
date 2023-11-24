package config

import (
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"echo/model"
)

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "123456"
	dbname   = "universal"
)

var (
	db  *gorm.DB
	err error
)

func Connect() {

	psqlInfo := fmt.Sprintf(`
	host=%s 
	port=%d 
	user=%s `+`
	password=%s 
	dbname=%s 
	sslmode=disable`,
		host, port, user, password, dbname)

	_ = psqlInfo

	dbUrl := os.Getenv("DATABASE_URL")

	//ini koneksi menggunakan database/sql
	// db, err = sql.Open("postgres", psqlInfo)

	//ini koneksi menggunakan gorm
	db, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database")

	db.Debug().AutoMigrate(&model.Employee{}, &model.Item{})

}

func GetDB() *gorm.DB {
	return db
}
