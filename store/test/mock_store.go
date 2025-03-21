package store

import (
	"context"
	"log"
	"time"
)

func getMockStore() Store {
}

// Add a functiuon to create a mock database using test containers
func CreateMockDatabase(dbConfig *databaseConfig) *PostgresContainer {
	log.Printf("This function was called")
	// Create a new postgres test container.
	ctx := context.Background()
	dbName := "users"
	dbUser := "user"
	dbPassword := "password"
	postgresContainer, err := postgres.Run(ctx,
		"postgres:16-alpine",
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPassword),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)
	defer func() {
		if err := testcontainers.TerminateContainer(postgresContainer); err != nil {
			log.Printf("Test Container Deref function called")
			log.Printf("failed to terminate container: %s", err)
		}
	}()
	if err != nil {
		log.Printf("failed to start container: %s", err)
		return
	}
	return postgresContainer
}
