package database

func init() {
	RegisterMigration(&Migration{
		ID:          "008_insert_default_settings",
		Description: "Insert default system settings (booking advance, cancellation notice, auto-deactivation)",
		Up: map[string]string{
			"sqlite": `
INSERT OR IGNORE INTO system_settings (key, value) VALUES
  ('booking_advance_days', '14'),
  ('cancellation_notice_hours', '12'),
  ('auto_deactivation_days', '365');
`,
			"mysql": "INSERT IGNORE INTO system_settings (`key`, value) VALUES\n" +
				"  ('booking_advance_days', '14'),\n" +
				"  ('cancellation_notice_hours', '12'),\n" +
				"  ('auto_deactivation_days', '365');",
			"postgres": `
INSERT INTO system_settings (key, value) VALUES
  ('booking_advance_days', '14'),
  ('cancellation_notice_hours', '12'),
  ('auto_deactivation_days', '365')
ON CONFLICT (key) DO NOTHING;
`,
		},
	})
}
