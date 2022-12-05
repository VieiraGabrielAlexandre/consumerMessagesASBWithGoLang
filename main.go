package main

import (
	"origomicrosservices.com/consumer-contratos/infraestruture/configs"
	"origomicrosservices.com/consumer-contratos/infraestruture/consumers"
	"origomicrosservices.com/consumer-contratos/infraestruture/database"
)

// Hello returns a greeting for the named person.
func main() {
	configs.Enviroment()
	database.Connect()
	consumers.ExecuteAll()
}
