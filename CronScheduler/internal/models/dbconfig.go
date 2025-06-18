package models

type DBConfig struct {
	DbName string `json:"dbName"`
	DbHost string `json:"dbHost"`
	DbPort string `json:"dbPort"`
	DbUser string `json:"dbUser"`
	DbPass string `json:"dbPass"`
}
