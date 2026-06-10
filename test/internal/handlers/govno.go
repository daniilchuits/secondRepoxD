package handlers

func (manager dbManager) Create() error {

	query := `
	CREATE TABLE IF NOT EXISTS zed(
		id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
		name TEXT
	)
	`

	_, err := manager.db.Exec(query)
	return err
}
