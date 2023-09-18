package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/haroldpacha/customer-management/entity"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDatabaseConnection() *gorm.DB {
	errEnv := godotenv.Load()

	if errEnv != nil {
		panic("failed to load env")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&entity.Customer{})

	if err = db.AutoMigrate(&entity.User{}); err == nil && db.Migrator().HasTable(&entity.User{}) {
		if err := db.First(&entity.User{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			user := entity.User{Name: "admin", Email: "admin@demo.com", Password: "$2a$04$aCy5gZxoqaVoj64WxJPQfOuAbn2y2.urXTf0Ny.6REOIcggdpCU52"}
			db.Create(&user)
		}
	}

	if err = db.AutoMigrate(&entity.Department{}); err == nil && db.Migrator().HasTable(&entity.Department{}) {
		if err := db.First(&entity.Department{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			var departments = []entity.Department{
				{Name: "Amazonas"},
				{Name: "Ancash"},
				{Name: "Apurimac"},
				{Name: "Arequipa"},
				{Name: "Ayacucho"},
				{Name: "Cajamarca"},
				{Name: "Callao"},
				{Name: "Cusco"},
				{Name: "Huancavelica"},
				{Name: "Huanuco"},
				{Name: "Ica"},
				{Name: "Junin"},
				{Name: "La libertad"},
				{Name: "Lambayeque"},
				{Name: "Lima"},
				{Name: "Loreto"},
				{Name: "Madre de dios"},
				{Name: "Moquegua"},
				{Name: "Pasco"},
				{Name: "Piura"},
				{Name: "Puno"},
				{Name: "San martin"},
				{Name: "Tacna"},
				{Name: "Tumbes"},
				{Name: "Ucayali"},
			}
			db.Create(departments)
		}
	}

	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()

	if err != nil {
		panic("failed to close database connection")
	}
	dbSQL.Close()
}
