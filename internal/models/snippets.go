package models

import (
	"database/sql"
	"errors"
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
	stmt := `SELECT id, title, content, created, expires FROM snippets
 WHERE expires > UTC_TIMESTAMP() AND id = ?`

	// QueryRow() is used to execute SQL stmt. by passing id as the placeholder variable.
	row := m.DB.QueryRow(stmt, id)

	// Initialize a pointer to a new zeroed Snippet struct
	s := &Snippet{}

	// row.Scan() is used to copy the values from each field in sql.Row
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	// return the Snippet object
	return s, nil
}

// returns most recently created snippets
func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}
