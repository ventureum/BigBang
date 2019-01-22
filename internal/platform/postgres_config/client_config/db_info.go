package client_config

import (
	"fmt"
	"os"
)

type DBInfo struct {
	DBUser     string `json:"dbUser"`
	DBPassword string `json:"dbPassword"`
	DBName     string `json:"dbName"`
	DBHost     string `json:"dbHost"`
}

func (dbInfo *DBInfo) ToString() string {
	dbInfoStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		dbInfo.DBHost, dbInfo.DBUser, dbInfo.DBPassword, dbInfo.DBName)
	return dbInfoStr
}

func CreateDefaultDBInfo() *DBInfo {
	env := os.Getenv("ENV")
	dbUser := os.Getenv("DB_USER_" + env)
	dbPassword := os.Getenv("DB_PASSWORD_" + env)
	dbName := os.Getenv("DB_NAME_" + env)
	dbHost := os.Getenv("DB_HOST_" + env)
	return &DBInfo{
		DBUser:     dbUser,
		DBPassword: dbPassword,
		DBName:     dbName,
		DBHost:     dbHost,
	}
}
