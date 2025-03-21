package main

import (
	"fmt"
	"log"

	"github.com/instaupload/user-management/types"
	"github.com/instaupload/user-management/utils"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println("Hello from User Management System")
	testDbConfig := types.DatabaseConfig{
		User:     utils.GetEnvString("TESTDATABASEUSER", "user"),
		Password: utils.GetEnvString("TESTDATABASEUSER", "user"),
		Name:     utils.GetEnvString("TESTDATABASEUSER", "user"),
	}
	log.Printf("testDbConfig: %v", testDbConfig)
	// store.CreateMockDatabase(testDbConfig)
}
