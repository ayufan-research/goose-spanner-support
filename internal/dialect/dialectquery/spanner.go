package dialectquery

import "fmt"

type Spanner struct{}

var _ Querier = (*Spanner)(nil)

func (p *Spanner) CreateTable(tableName string) string {
	q := `CREATE TABLE %s (
		version_id INT64 NOT NULL,
		is_applied BOOL NOT NULL,
		tstamp TIMESTAMP DEFAULT (CURRENT_TIMESTAMP()),
	) PRIMARY KEY(version_id)`
	return fmt.Sprintf(q, tableName)
}

func (p *Spanner) InsertVersion(tableName string) string {
	q := `INSERT INTO %s (version_id, is_applied) VALUES (?, ?)`
	return fmt.Sprintf(q, tableName)
}

func (p *Spanner) DeleteVersion(tableName string) string {
	q := `DELETE FROM %s WHERE version_id=?`
	return fmt.Sprintf(q, tableName)
}

func (p *Spanner) GetMigrationByVersion(tableName string) string {
	q := `SELECT tstamp, is_applied FROM %s WHERE version_id=? ORDER BY tstamp DESC LIMIT 1`
	return fmt.Sprintf(q, tableName)
}

func (p *Spanner) ListMigrations(tableName string) string {
	q := `SELECT version_id, is_applied from %s ORDER BY version_id DESC`
	return fmt.Sprintf(q, tableName)
}
