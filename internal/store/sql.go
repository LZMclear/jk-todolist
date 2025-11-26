package store

import "database/sql"

func InitDB(db *sql.DB) error {
	// create table if not exists (MySQL compatible)
	schema := `
	CREATE TABLE IF NOT EXISTS tasks (
		id BIGINT AUTO_INCREMENT PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		description TEXT,
		category VARCHAR(100) DEFAULT '',
		completed TINYINT(1) NOT NULL DEFAULT 0,
		due_date DATETIME NULL,
		created_at DATETIME NOT NULL,
		updated_at DATETIME NOT NULL
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
	`
	if _, err := db.Exec(schema); err != nil {
		return err
	}
	return nil
}
