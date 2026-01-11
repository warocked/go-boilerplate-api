package main

import (
	"context"
	"flag"
	"log"
	"os"

	"go-boilerplate-api/internal/api/config"
	"go-boilerplate-api/internal/api/db"
)

func main() {
	var (
		command     = flag.String("command", "", "Migration command: up, version, validate, create")
		name        = flag.String("name", "", "Migration name (required for create command)")
		databaseURL = flag.String("database-url", "", "Database URL (overrides DATABASE_URL env var)")
	)
	flag.Parse()

	// Load configuration
	if err := config.LoadEnvFile(); err != nil {
		log.Printf("Warning: Failed to load .env file: %v", err)
	}

	if err := config.LoadAllConfig(); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Use provided database URL or from config
	dbURL := *databaseURL
	if dbURL == "" {
		dbURL = config.DATABASE_URL
	}

	if dbURL == "" {
		log.Fatalf("Database URL is required. Set DATABASE_URL environment variable or use -database-url flag")
	}

	ctx := context.Background()

	switch *command {
	case "up":
		log.Println("Running migrations UP...")
		if err := db.RunMigrations(ctx, dbURL); err != nil {
			log.Fatalf("Migration failed: %v", err)
		}
		log.Println("Migrations completed successfully")

	case "version":
		version, dirty, err := db.GetMigrationVersion(ctx, dbURL)
		if err != nil {
			log.Fatalf("Failed to get migration version: %v", err)
		}
		if dirty {
			log.Printf("Current migration version: %d (DIRTY - manual intervention required)", version)
			os.Exit(1)
		} else {
			log.Printf("Current migration version: %d", version)
		}

	case "validate":
		if err := db.ValidateMigrations(ctx, dbURL); err != nil {
			log.Fatalf("Migration validation failed: %v", err)
		}
		log.Println("Migrations are valid")

	case "create":
		if *name == "" {
			log.Fatal("Migration name is required. Use -name flag")
		}
		log.Printf("Creating migration files: %s", *name)
		// Note: For creating migration files, you may want to use migrate CLI tool:
		// migrate create -ext sql -dir internal/api/db/migrations -seq <name>
		log.Println("Use 'migrate create -ext sql -dir internal/api/db/migrations -seq <name>' to create migration files")
		log.Println("Or manually create <version>_<name>.up.sql and <version>_<name>.down.sql files")

	default:
		log.Fatal("Invalid command. Use: up, version, validate, or create")
	}
}
