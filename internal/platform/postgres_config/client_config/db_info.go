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
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	return &DBInfo{
		DBUser:     dbUser,
		DBPassword: dbPassword,
		DBName:     dbName,
		DBHost:     dbHost,
	}
}
