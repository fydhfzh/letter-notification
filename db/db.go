package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fydhfzh/letter-notification/dto"
	"github.com/fydhfzh/letter-notification/entity"
	"github.com/fydhfzh/letter-notification/pkg/helpers"
	"github.com/fydhfzh/letter-notification/pkg/mailer"
	"github.com/fydhfzh/letter-notification/repository/letter_repository/letter_pg"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db    *gorm.DB
	dbErr error
)

func InitializeDB() {
	handleDBConnection()
	createRequiredTable()
	seedAdminData()
	seedSubditData()
	rescheduleAllMail()
}

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}
}

func handleDBConnection() {
	PORT := os.Getenv("PG_PORT")
	USER := os.Getenv("PG_USER")
	PASSWORD := os.Getenv("PG_PASSWORD")
	HOST := os.Getenv("PG_HOST")
	DBNAME := os.Getenv("PG_DBNAME")

	psqlInfo := fmt.Sprintf("port=%s user=%s password=%s host=%s dbname=%s sslmode=disable", PORT, USER, PASSWORD, HOST, DBNAME)

	db, dbErr = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{TranslateError: true})

	if dbErr != nil {
		log.Fatal(dbErr)
	}
}

func createRequiredTable() {
	err := db.AutoMigrate(entity.User{})

	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(entity.Subdit{})

	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(entity.Letter{})

	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(entity.UserLetter{})

	if err != nil {
		log.Fatal(err)
	}

}

func seedAdminData() {
	admin := entity.User{
		Name:        "admin",
		Email:       "admin@gmail.com",
		Password:    "admin123",
		PhoneNumber: "0812345678",
		Role:        "admin",
	}

	hashed, err := helpers.HashPassword(admin.Password)

	if err != nil {
		log.Panic("Error on password hashing")
	}

	admin.Password = hashed

	_ = db.Create(&admin)
}

func seedSubditData() {

	subdits := []entity.Subdit{
		{
			Name: "Kawasan Khusus Lingkup I",
		},
		{
			Name: "Fasilitasi Masalah Pertanahan",
		},
		{
			Name: "Kawasan Khusus Lingkup II",
		},
		{
			Name: "Administrasi Kawasan Perkotaan",
		},
		{
			Name: "Batas Negara dan Pulau-Pulau Terluar",
		},
	}

	result := db.Create(&subdits)

	if err := result.Error; err != nil {
		log.Panic(err)
	}

}

func rescheduleAllMail() {
	var letters []entity.Letter

	// Reschedule email before one day
	currTime := time.Now().Add(-24 * time.Hour).Format(time.RFC3339)

	result := db.Where("datetime > ?", currTime).Where("is_notified = ?", false).Find(&letters)

	if err := result.Error; err != nil {
		log.Panic(err)
	}

	for _, letter := range letters {
		var users []entity.User

		result := db.Where("subdit_id = ?", letter.ToSubditID).Find(&users)

		if err := result.Error; err != nil {
			log.Panic(err)
		}

		scheduler := dto.SendLetterToMailScheduler{
			About:      letter.About,
			Datetime:   letter.Datetime,
			Recipients: users,
		}

		fmt.Println("Resending email to assigned users")
		mailer.SetSchedule(scheduler, letter_pg.NewLetterRepository(db))

		letter.IsNotified = true

		result = db.Save(&letter)

		if err := result.Error; err != nil {
			log.Panic(err)
		}
	}
}

func GetDBInstance() *gorm.DB {
	return db
}
