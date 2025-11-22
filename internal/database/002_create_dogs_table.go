package database

func init() {
	RegisterMigration(&Migration{
		ID:          "002_create_dogs_table",
		Description: "Create dogs table with photo support",
		Up: map[string]string{
			"sqlite": `
CREATE TABLE IF NOT EXISTS dogs (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL,
  breed TEXT NOT NULL,
  size TEXT CHECK(size IN ('small', 'medium', 'large')),
  age INTEGER,
  category TEXT CHECK(category IN ('green', 'blue', 'orange')),
  photo TEXT,
  special_needs TEXT,
  pickup_location TEXT,
  walk_route TEXT,
  walk_duration INTEGER,
  special_instructions TEXT,
  default_morning_time TEXT,
  default_evening_time TEXT,
  is_available INTEGER DEFAULT 1,
  unavailable_reason TEXT,
  unavailable_since TIMESTAMP,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_dogs_available ON dogs(is_available, category);
`,
			"mysql": `
CREATE TABLE IF NOT EXISTS dogs (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  breed VARCHAR(255) NOT NULL,
  size VARCHAR(20) CHECK(size IN ('small', 'medium', 'large')),
  age INT,
  category VARCHAR(20) CHECK(category IN ('green', 'blue', 'orange')),
  photo VARCHAR(255),
  special_needs TEXT,
  pickup_location VARCHAR(255),
  walk_route TEXT,
  walk_duration INT,
  special_instructions TEXT,
  default_morning_time VARCHAR(10),
  default_evening_time VARCHAR(10),
  is_available TINYINT(1) DEFAULT 1,
  unavailable_reason TEXT,
  unavailable_since DATETIME,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE INDEX IF NOT EXISTS idx_dogs_available ON dogs(is_available, category);
`,
			"postgres": `
CREATE TABLE IF NOT EXISTS dogs (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  breed VARCHAR(255) NOT NULL,
  size VARCHAR(20) CHECK(size IN ('small', 'medium', 'large')),
  age INTEGER,
  category VARCHAR(20) CHECK(category IN ('green', 'blue', 'orange')),
  photo VARCHAR(255),
  special_needs TEXT,
  pickup_location VARCHAR(255),
  walk_route TEXT,
  walk_duration INTEGER,
  special_instructions TEXT,
  default_morning_time VARCHAR(10),
  default_evening_time VARCHAR(10),
  is_available BOOLEAN DEFAULT TRUE,
  unavailable_reason TEXT,
  unavailable_since TIMESTAMP WITH TIME ZONE,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_dogs_available ON dogs(is_available, category);
`,
		},
	})
}
