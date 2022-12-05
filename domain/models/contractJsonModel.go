package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type ContractJSON struct {
	gorm.Model
	UUID string `json:"uuid" validate:"nonzero"`
	Data string `gorm:"text" json:"data" validate:"nonzero"`
}

func ValidatorContractJSON(contractJson *ContractJSON) error {
	if err := validator.Validate(contractJson); err != nil {
		return err
	}
	return nil
}
