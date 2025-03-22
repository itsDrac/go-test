package main

import (
	"fmt"
	"log"

	store "github.com/InstaUpload/user-management/store/test"
	"github.com/InstaUpload/user-management/types"
	"github.com/InstaUpload/user-management/utils"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println("Hello from User Management System")
	testDbConfig := types.DatabaseConfig{
		User:         utils.GetEnvString("TESTDATABASEUSER", "user"),
		Password:     utils.GetEnvString("TESTDATABASEPASSWORD", "user"),
		Name:         utils.GetEnvString("TESTDATABASENAME", "user"),
		MaxOpenConns: utils.GetEnvInt("TESTDATABASEOPENCONNS", 5),
		MaxIdleConns: utils.GetEnvInt("TESTDATABASEIDLECONNS", 5),
		MaxIdleTime:  utils.GetEnvString("TESTDATABASEIDLETIME", "1m"),
	}
	log.Printf("testDbConfig: %v", testDbConfig)
	store.CreateMockDatabase(&testDbConfig)
}
