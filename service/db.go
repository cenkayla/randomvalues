package service

import (
	"encoding/json"
	"net/http"

	"github.com/cenkayla/randomvalues/models"
	"github.com/google/uuid"
)

func (s *Service) generateValue(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	v := models.Value{}
	err := json.NewDecoder(r.Body).Decode(&v)

	v.Name = models.GenerateURL(v.Number)
	v.UUID = uuid.Must(uuid.NewRandom())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = s.DB.Create(&v).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("Succesfully generated a value")
}

func (s *Service) retrieveValue(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	v := models.Value{}
	err := json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = s.DB.First(&v).Where("uuid = ?", v.UUID).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(&v)
}
