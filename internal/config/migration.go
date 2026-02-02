package config

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
)

func RunMigrations(db *sql.DB) error {
	migrationFile := "migrations/001_create_tables.sql"

	if _, err := os.Stat(migrationFile); os.IsNotExist(err) {
		return fmt.Errorf("migration file not found: %s", migrationFile)
	}

	content, err := os.ReadFile(migrationFile)
	if err != nil {
		return fmt.Errorf("failed to read migration file: %v", err)
	}

	lines := strings.Split(string(content), "\n")
	var cleanSQL strings.Builder

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "--") || trimmed == "" {
			continue
		}
		cleanSQL.WriteString(line)
		cleanSQL.WriteString("\n")
	}

	statements := strings.Split(cleanSQL.String(), ";")

	executedCount := 0
	for i, stmt := range statements {
		stmt = strings.TrimSpace(stmt)
		if stmt == "" {
			continue
		}

		fmt.Printf("Executing migration statement %d...\n", i+1)
		_, err := db.Exec(stmt)
		if err != nil {
			return fmt.Errorf("failed to execute migration statement %d: %v\nStatement: %s", i+1, err, stmt)
		}
		executedCount++
	}

	fmt.Printf("✓ Database migrations completed successfully (%d statements executed)\n", executedCount)
	return nil
}
