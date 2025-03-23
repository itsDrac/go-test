package store

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/InstaUpload/user-management/types"
	"github.com/InstaUpload/user-management/utils"
	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
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
	mockStore, err := CreateMockDatabase(testDbConfig)
	if err != nil {
		log.Fatalf("Error creating mock database %v", err)
	}
	exitCode := m.Run()
	KillPostgresContainer()
	os.Exit(exitCode)

}

// Should I create a function to test all the interface of Store?
