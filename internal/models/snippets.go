package models

import (
	"database/sql"
	"time"
)

// holds data for an individual snippet
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// snippet model which wraps a sql.DB connection pool
type SnippetModel struct {
	DB *sql.DB
}

// will insert a new snippet into the database
func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires)
 	VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	// Exec() method is used to execute the statement. 1st parameter is the SQL statement,
	// title, content & expires.
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	// LastInsertId() is used to get the id of
	// the newly inserted record in snippets table
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// returns a specific snippet based on its id
func (m *SnippetModel) Get(id int) (*Snippet, error) {
	return nil, nil
}

// returns most recently created snippets
func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}
