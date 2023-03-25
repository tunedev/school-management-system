package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/tunedev/school-management-system/internal/config"
	"github.com/tunedev/school-management-system/pkg/database"
)

func main() {
	// Load configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	// Connect to database
	migration, err := database.NewDatabase(cfg.DSN())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	fmt.Println("Database connection established successful, running the sql scripts now")

	// get all migration scripts in the /scripts/migration folder
	files, err := filepath.Glob("scripts/migration/*.sql")
	if err != nil {
    log.Fatalf("failed to find migration scripts: %v", err)
	}

	// itereate ove all the migrations scripts and execute them one by one
	for _, file := range files {
		fmt.Printf("Starting migration for =====>>> %v", file)
		script, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatalf("failed to read migration script %s: %v", file, err)
		}
		
		_, err = migration.DB.Exec(string(script))
		if err != nil {
			log.Fatalf("failed to run migration script %s: %v", file, err)
		}
		fmt.Printf("Done running migration for =====>>> %v", file)
	}
	// // Run database migrations
	// migrations, err := ioutil.ReadFile("./scripts/migration/initial_table_setup.sql")
	// if err != nil {
	// 	log.Fatalf("failed to read migration file: %v", err)
	// }

	// _, err = db.DB.Exec(string(migrations))
	// if err != nil {
	// 	log.Fatalf("failed to run database migrations: %v", err)
	// }

	log.Println("database migrations completed successfully")
}
