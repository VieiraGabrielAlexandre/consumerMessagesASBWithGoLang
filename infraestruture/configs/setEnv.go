package configs

import "os"

func Enviroment() {
	os.Setenv("TOPIC_CONTRATOS_FINALIZADOS", "")
	os.Setenv("TOPIC_CONTRATOS_FINALIZADOS_SUBSCRIPTION", "")
	os.Setenv("TOPIC_CONTRATOS_FINALIZADOS_STRING_CONNECTION", "")
	os.Setenv("TOPIC_CONTRATOS_FINALIZADOS_PRIMARY_KEY", "")
	os.Setenv("TOPIC_CONTRATOS_FINALIZADOS_HOST", "")

	os.Setenv("AZURE_CLIENT_ID", "")
	os.Setenv("AZURE_TENANT_ID", "")
	os.Setenv("AZURE_CLIENT_SECRET", "")

	os.Setenv("DB_USER", "root")
	os.Setenv("DB_PASS", "golang")
	os.Setenv("DB_HOST", "127.0.0.1")
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_NAME", "")
}
