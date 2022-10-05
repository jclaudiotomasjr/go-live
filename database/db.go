package database

import (
	"log"

	"github.com/jclaudiotomasjr/go-live/api-go-gin-quiz/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBanco() {
	//stringDeConexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err := gorm.Open(sqlite.Open("data/scores.db"), &gorm.Config{})
	//DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com Banco de dados!")
	}
	DB.AutoMigrate(&models.Score{})
}
