package database

import (
	"database/sql"
	"fmt"
	"strings"
)

// MySQLDialect implements the Dialect interface for MySQL
type MySQLDialect struct{}

// NewMySQLDialect creates a new MySQL dialect
func NewMySQLDialect() *MySQLDialect {
	return &MySQLDialect{}
}

// Name returns the database name
func (d *MySQLDialect) Name() string {
	return "mysql"
}

// GetDriverName returns the Go driver name for sql.Open()
func (d *MySQLDialect) GetDriverName() string {
	return "mysql"
}

// GetAutoIncrement returns the auto-increment syntax for primary keys
func (d *MySQLDialect) GetAutoIncrement() string {
	return "INT AUTO_INCREMENT PRIMARY KEY"
}

// GetBooleanType returns the boolean column type
// MySQL uses TINYINT(1) for booleans
func (d *MySQLDialect) GetBooleanType() string {
	return "TINYINT(1)"
}

// GetBooleanDefault returns the default value for a boolean
func (d *MySQLDialect) GetBooleanDefault(value bool) string {
	if value {
		return "1"
	}
	return "0"
}

// GetTextType returns the text column type
// MySQL requires explicit size limits for indexed VARCHAR fields
func (d *MySQLDialect) GetTextType(maxLength int) string {
	if maxLength == 0 {
		// Unlimited text (for long content)
		return "TEXT"
	}
	// Sized text (for indexed fields like email, name)
	return fmt.Sprintf("VARCHAR(%d)", maxLength)
}

// GetTimestampType returns the timestamp column type
// MySQL uses DATETIME for timestamp storage
func (d *MySQLDialect) GetTimestampType() string {
	return "DATETIME"
}

// GetCurrentDate returns SQL expression for current date
func (d *MySQLDialect) GetCurrentDate() string {
	return "CURDATE()"
}

// GetCurrentTimestamp returns SQL expression for current timestamp
func (d *MySQLDialect) GetCurrentTimestamp() string {
	return "CURRENT_TIMESTAMP"
}

// GetPlaceholder returns the placeholder syntax
// MySQL uses ? for all parameters (same as SQLite)
func (d *MySQLDialect) GetPlaceholder(position int) string {
	return "?"
}

// SupportsIfNotExistsColumn returns whether database supports
// ALTER TABLE ADD COLUMN IF NOT EXISTS
// MySQL 5.7+ does not support IF NOT EXISTS for ADD COLUMN
// MySQL 8.0.29+ does support it, but we'll be conservative
func (d *MySQLDialect) SupportsIfNotExistsColumn() bool {
	return false // Conservative - check column existence before adding
}

// GetInsertOrIgnore returns the SQL for insert-or-ignore semantics
func (d *MySQLDialect) GetInsertOrIgnore(tableName string, columns []string, placeholders string) string {
	columnList := strings.Join(columns, ", ")
	return fmt.Sprintf("INSERT IGNORE INTO %s (%s) VALUES (%s)",
		tableName, columnList, placeholders)
}

// GetAddColumnSyntax returns SQL for adding a column
// MySQL doesn't support IF NOT EXISTS for ALTER TABLE ADD COLUMN
// Caller must handle duplicate column error
func (d *MySQLDialect) GetAddColumnSyntax(tableName, columnName, columnType string) string {
	return fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s %s",
		tableName, columnName, columnType)
}

// ApplySettings applies MySQL-specific settings
func (d *MySQLDialect) ApplySettings(db *sql.DB) error {
	// Set charset to utf8mb4 for full Unicode support (including emoji)
	if _, err := db.Exec("SET NAMES utf8mb4"); err != nil {
		return fmt.Errorf("failed to set charset: %w", err)
	}

	// Set timezone to UTC for consistency
	if _, err := db.Exec("SET time_zone = '+00:00'"); err != nil {
		return fmt.Errorf("failed to set timezone: %w", err)
	}

	// Set SQL mode for strict behavior (recommended for data integrity)
	// TRADITIONAL mode enables strict checking
	if _, err := db.Exec("SET sql_mode = 'TRADITIONAL'"); err != nil {
		// Non-fatal - some MySQL versions may not support all modes
		// Just log and continue
		fmt.Printf("Warning: Failed to set SQL mode: %v\n", err)
	}

	return nil
}

// GetTableCreationSuffix returns MySQL table creation suffix
// Specifies storage engine and character set
func (d *MySQLDialect) GetTableCreationSuffix() string {
	return " ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci"
}

// QuoteIdentifier returns the quoted identifier using backticks
// MySQL uses backticks for identifiers
func (d *MySQLDialect) QuoteIdentifier(identifier string) string {
	// Generally not needed in our queries
	// Return as-is unless identifier is a reserved word
	return identifier
}

// ConvertGoTime returns SQL expression to convert Go time.Time
// MySQL driver handles this automatically
func (d *MySQLDialect) ConvertGoTime(goTime string) string {
	return goTime // Driver handles conversion
}
