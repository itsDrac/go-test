package main

import (
	"fmt"
	"log"

	"github.com/InstaUpload/user-management/store/database"
	"github.com/InstaUpload/user-management/types"
	"github.com/InstaUpload/user-management/utils"
)

func main() {
	dbConfig := types.DatabaseConfig{
		User:         utils.GetEnvString("DATABASEUSER", "user"),
		Password:     utils.GetEnvString("DATABASEPASSWORD", "user"),
		Name:         utils.GetEnvString("DATABASENAME", "user"),
		MaxOpenConns: utils.GetEnvInt("DATABASEOPENCONNS", 5),
		MaxIdleConns: utils.GetEnvInt("DATABASEIDLECONNS", 5),
		MaxIdleTime:  utils.GetEnvString("DATABASEIDLETIME", "1m"),
	}
	connectionString := fmt.Sprintf("postgress://%s:%s@localhost/%s?sslmode=disable", dbConfig.User, dbConfig.Password, dbConfig.Name)
	log.Printf("connection string %s", connectionString)
	testDbConfig.SetConnectionString(connectionString)
	db, err := database.New(&testDbConfig)
	if err != nil {
		log.Fatal("Can not connect to database %v", err)
	}
	store = store.New(db)
	log.Printf("Store created, calling user.GetStr fn: %s", store.GetString())
}
