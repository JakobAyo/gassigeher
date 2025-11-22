package database

func init() {
	RegisterMigration(&Migration{
		ID:          "006_create_system_settings_table",
		Description: "Create system_settings table for runtime configuration",
		Up: map[string]string{
			"sqlite": `
CREATE TABLE IF NOT EXISTS system_settings (
  key TEXT PRIMARY KEY,
  value TEXT NOT NULL,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
`,
			"mysql": "CREATE TABLE IF NOT EXISTS system_settings (\n" +
				"  `key` VARCHAR(255) PRIMARY KEY,\n" +
				"  value TEXT NOT NULL,\n" +
				"  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP\n" +
				") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;",
			"postgres": `
CREATE TABLE IF NOT EXISTS system_settings (
  key VARCHAR(255) PRIMARY KEY,
  value TEXT NOT NULL,
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
`,
		},
	})
}
