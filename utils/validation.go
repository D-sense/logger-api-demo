package utils

import (
	"net/http"
)

type ValidationErr string

const (
	ErrMissingImage     = ValidationErr("Image is missing")
	DataValidationErr   = http.StatusUnprocessableEntity
)

func ValidateImageInput(image string) (err map[string]ValidationErr) {
	err = make(map[string]ValidationErr)

	if image == "" {
		err["Email"] = ErrMissingImage
	}

	return err
}