package service

import (
	"log"
	"net/http"
	"os"

	"github.com/cenkayla/randomvalues/models"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Service struct {
	DB *gorm.DB
}

func (s *Service) Open() error {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error when loading .env file")
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE")), &gorm.Config{})
	if err != nil {
		return err
	}
	s.DB = db
	s.ConfigureRouter()
	return s.DB.AutoMigrate(models.Value{})
}

func (s *Service) ConfigureRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/api/generate", s.generateValue).Methods("POST")
	r.HandleFunc("/api/retrieve", s.retrieveValue).Methods("POST")
	return r
}
