package services

import (
	"origomicrosservices.com/consumer-contratos/domain/models"
	"origomicrosservices.com/consumer-contratos/infraestruture/database"
)

func SaveJSON(uuid string, data string) bool {

	contratoFinalizadoJSONStruct := models.ContractJSON{
		UUID: uuid,
		Data: data,
	}

	err := database.DB.FirstOrCreate(&contratoFinalizadoJSONStruct, "uuid = ?", uuid).Error

	return err == nil
}
