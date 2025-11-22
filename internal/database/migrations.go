package database

import (
	"database/sql"
	"fmt"
	"log"
	"sort"
	"time"
)

// Migration represents a database schema migration
// Each migration has SQL statements for each supported database type
type Migration struct {
	ID          string            // Unique identifier (e.g., "001_create_users_table")
	Description string            // Human-readable description
	Up          map[string]string // SQL statements for each database type (sqlite, mysql, postgres)
}

// migrationRegistry stores all registered migrations
var migrationRegistry []*Migration

// RegisterMigration adds a migration to the registry
// Called by init() functions in migration files
func RegisterMigration(m *Migration) {
	migrationRegistry = append(migrationRegistry, m)
}

// GetAllMigrations returns all registered migrations in sorted order by ID
func GetAllMigrations() []*Migration {
	// Sort by ID to ensure consistent order
	sorted := make([]*Migration, len(migrationRegistry))
	copy(sorted, migrationRegistry)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].ID < sorted[j].ID
	})
	return sorted
}

// RunMigrationsWithDialect applies all pending migrations for the given dialect
// This is the new migration system that supports multiple databases
// The old RunMigrations() function in database.go is kept for backward compatibility
func RunMigrationsWithDialect(db *sql.DB, dialect Dialect) error {
	// Create schema_migrations table if it doesn't exist
	if err := createSchemaMigrationsTable(db, dialect); err != nil {
		return fmt.Errorf("failed to create schema_migrations table: %w", err)
	}

	// Get list of applied migrations
	appliedMigrations, err := getAppliedMigrations(db)
	if err != nil {
		return fmt.Errorf("failed to get applied migrations: %w", err)
	}

	// Get all registered migrations
	allMigrations := GetAllMigrations()

	// Apply pending migrations
	pendingCount := 0
	for _, migration := range allMigrations {
		// Skip if already applied
		if appliedMigrations[migration.ID] {
			continue
		}

		// Get SQL for current dialect
		sql, ok := migration.Up[dialect.Name()]
		if !ok {
			return fmt.Errorf("migration %s does not support database type: %s", migration.ID, dialect.Name())
		}

		// Execute migration
		log.Printf("Applying migration: %s - %s", migration.ID, migration.Description)
		if _, err := db.Exec(sql); err != nil {
			// Special handling for "already exists" errors
			if isAlreadyExistsError(err, dialect) {
				log.Printf("Migration %s: Object already exists, marking as applied", migration.ID)
				// Mark as applied even though exec failed (idempotency)
				if err := markMigrationAsApplied(db, migration.ID); err != nil {
					return fmt.Errorf("failed to mark migration as applied: %w", err)
				}
				pendingCount++
				continue
			}
			return fmt.Errorf("migration %s failed: %w", migration.ID, err)
		}

		// Mark migration as applied
		if err := markMigrationAsApplied(db, migration.ID); err != nil {
			return fmt.Errorf("failed to mark migration %s as applied: %w", migration.ID, err)
		}

		pendingCount++
	}

	if pendingCount > 0 {
		log.Printf("Applied %d migration(s)", pendingCount)
	} else {
		log.Printf("No pending migrations")
	}

	return nil
}

// createSchemaMigrationsTable creates the schema versioning table
func createSchemaMigrationsTable(db *sql.DB, dialect Dialect) error {
	// Build CREATE TABLE statement using dialect
	sql := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS schema_migrations (
			version %s NOT NULL,
			applied_at %s DEFAULT %s
		)%s`,
		dialect.GetTextType(255), // version as VARCHAR(255)
		dialect.GetTimestampType(),
		dialect.GetCurrentTimestamp(),
		dialect.GetTableCreationSuffix(),
	)

	if _, err := db.Exec(sql); err != nil {
		return fmt.Errorf("failed to create schema_migrations table: %w", err)
	}

	// Create index on version column for faster lookups
	indexSQL := "CREATE INDEX IF NOT EXISTS idx_schema_migrations_version ON schema_migrations(version)"
	if _, err := db.Exec(indexSQL); err != nil {
		// Non-fatal - some databases might not support IF NOT EXISTS for indexes
		log.Printf("Warning: Failed to create index on schema_migrations: %v", err)
	}

	return nil
}

// getAppliedMigrations returns a map of applied migration IDs
func getAppliedMigrations(db *sql.DB) (map[string]bool, error) {
	applied := make(map[string]bool)

	rows, err := db.Query("SELECT version FROM schema_migrations")
	if err != nil {
		return nil, fmt.Errorf("failed to query schema_migrations: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var version string
		if err := rows.Scan(&version); err != nil {
			return nil, fmt.Errorf("failed to scan migration version: %w", err)
		}
		applied[version] = true
	}

	return applied, nil
}

// markMigrationAsApplied records a migration as applied in schema_migrations
func markMigrationAsApplied(db *sql.DB, migrationID string) error {
	_, err := db.Exec("INSERT INTO schema_migrations (version, applied_at) VALUES (?, ?)",
		migrationID, time.Now())
	if err != nil {
		return fmt.Errorf("failed to insert migration record: %w", err)
	}
	return nil
}

// isAlreadyExistsError checks if an error is an "already exists" error
// Different databases return different error messages
func isAlreadyExistsError(err error, dialect Dialect) bool {
	if err == nil {
		return false
	}

	errMsg := err.Error()

	switch dialect.Name() {
	case "sqlite":
		// SQLite errors for existing objects
		return contains(errMsg, "already exists") ||
			contains(errMsg, "duplicate column name")

	case "mysql":
		// MySQL errors
		return contains(errMsg, "already exists") ||
			contains(errMsg, "Duplicate column name") ||
			contains(errMsg, "duplicate key")

	case "postgres":
		// PostgreSQL errors
		return contains(errMsg, "already exists") ||
			contains(errMsg, "duplicate column") ||
			contains(errMsg, "duplicate key value")

	default:
		return false
	}
}

// contains is a case-insensitive substring check helper
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(substr) == 0 ||
		(len(s) > 0 && len(substr) > 0 && indexIgnoreCase(s, substr) >= 0))
}

// indexIgnoreCase performs case-insensitive substring search
func indexIgnoreCase(s, substr string) int {
	sLower := toLower(s)
	substrLower := toLower(substr)
	for i := 0; i <= len(sLower)-len(substrLower); i++ {
		if sLower[i:i+len(substrLower)] == substrLower {
			return i
		}
	}
	return -1
}

// toLower converts string to lowercase (simple ASCII version)
func toLower(s string) string {
	result := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= 'A' && c <= 'Z' {
			c = c + ('a' - 'A')
		}
		result[i] = c
	}
	return string(result)
}

// GetMigrationStatus returns information about migration status
func GetMigrationStatus(db *sql.DB, dialect Dialect) (applied int, pending int, err error) {
	// Ensure schema_migrations table exists
	if err := createSchemaMigrationsTable(db, dialect); err != nil {
		return 0, 0, err
	}

	// Get applied migrations
	appliedMigrations, err := getAppliedMigrations(db)
	if err != nil {
		return 0, 0, err
	}

	// Get all migrations
	allMigrations := GetAllMigrations()

	applied = len(appliedMigrations)
	pending = len(allMigrations) - applied

	return applied, pending, nil
}
