package handlers

import (
	"carto/internal/db"
	"carto/internal/models"
	"carto/internal/utils"
	"carto/internal/validators"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	validate := validator.New()

	err := validate.RegisterValidation("strongpassword", utils.StrongPasswordFunction)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	var req validators.RegisterRequest

	if err := validate.Struct(req); err != nil {
		validationErrors := make(map[string][]string)
		for _, ve := range err.(validator.ValidationErrors) {
			validationErrors[ve.Field()] = []string{
				fmt.Sprintf("failed %s validation", ve.Tag()),
			}
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"errors": validationErrors,
		})
		return
	}

	var existingUser models.User

	if err := db.DB.First(&existingUser).Error; err != nil {

	}
}
