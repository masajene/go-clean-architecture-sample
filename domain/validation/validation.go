package validation

import (
	"github.com/go-playground/validator/v10"
	"time"
)

func NewValidation() (*validator.Validate, error) {
	vl := validator.New()
	err := vl.RegisterValidation("date", dateValidation)
	if err != nil {
		return nil, err
	}
	return vl, nil
}

func dateValidation(fl validator.FieldLevel) bool {
	_, err := time.Parse("2006-01-02", fl.Field().String())
	if err != nil {
		return false
	}
	return true
}
