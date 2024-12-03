package validators

import (
	"errors"
	"passport/internal/models"
)

func ValidateHuman(human *models.Human) error {
	if human.FirstName == "" || human.LastName == "" || human.Surname == "" || human.DOB == "" || human.SerialNumber == "" || human.Address == "" || human.CodeStructure == "" {
		return errors.New("invalid human data")
	}

	return nil
}
