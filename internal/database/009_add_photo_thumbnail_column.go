package database

func init() {
	RegisterMigration(&Migration{
		ID:          "009_add_photo_thumbnail_column",
		Description: "Add photo_thumbnail column to dogs table for optimized thumbnail storage",
		Up: map[string]string{
			"sqlite": `
-- SQLite doesn't support IF NOT EXISTS for ALTER TABLE ADD COLUMN before version 3.35.0
-- The migration runner will handle duplicate column errors gracefully
ALTER TABLE dogs ADD COLUMN photo_thumbnail TEXT;
`,
			"mysql": `
-- MySQL doesn't reliably support IF NOT EXISTS for ALTER TABLE ADD COLUMN
-- The migration runner will handle duplicate column errors gracefully
ALTER TABLE dogs ADD COLUMN photo_thumbnail VARCHAR(255);
`,
			"postgres": `
-- PostgreSQL 9.6+ supports IF NOT EXISTS for ADD COLUMN
ALTER TABLE dogs ADD COLUMN IF NOT EXISTS photo_thumbnail VARCHAR(255);
`,
		},
	})
}
