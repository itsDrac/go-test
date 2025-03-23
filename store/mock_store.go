package store

import (
	"context"
	"log"
	"time"

	"github.com/InstaUpload/user-management/store/db"
	"github.com/InstaUpload/user-management/types"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/testcontainers/testcontainers-go"
	tcpg "github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func CreateMockStore(dbConfig *types.DatabaseConfig) (Store, error) {
	ctx := context.Background()
	// Create database Test container.
	container := createPostgresContainer(ctx, dbConfig)
	connectionString, err := container.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		log.Fatalf("Error in getting connection string: %v", err)
		return nil, err
	}
	log.Printf("connection string %s", connectionString)
	// Using the connection string create database/sql store.
	dbConfig.SetConnectionString(connectionString)
	db, err := db.New(dbConfig)
	if err != nil {
		log.Fatalf("Error in connecting to test database %v", err)
		return nil, err
	}
	log.Printf("test Database connected")
	// Do migrations in the database.
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Error in setting up migration driver. %v", err)
		return nil, err
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://./migrations",
		"postgres", driver)
	if err != nil {
		log.Fatalf("Error in migrating %v", err)
		return nil, err
	}
	m.Up()
	// To be removed when debuging is completed.
	// defer killPostgresContainer(container)
	// defer db.Close()
	return NewStore(db), nil

}

// Function to run when test are over.
func KillPostgresContainer(container *tcpg.PostgresContainer) {
	if err := testcontainers.TerminateContainer(container); err != nil {
		log.Fatalf("failed to terminate container: %s", err)
	}
}

func createPostgresContainer(ctx context.Context, dbConfig *types.DatabaseConfig) *tcpg.PostgresContainer {
	// Create a new postgres test container.
	dbName := dbConfig.Name
	dbUser := dbConfig.User
	dbPassword := dbConfig.Password
	postgresContainer, err := tcpg.Run(ctx,
		"postgres:16-alpine",
		tcpg.WithDatabase(dbName),
		tcpg.WithUsername(dbUser),
		tcpg.WithPassword(dbPassword),
		testcontainers.WithHostPortAccess(5432),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		log.Printf("failed to start container: %s", err)
		return nil
	}
	return postgresContainer
}
