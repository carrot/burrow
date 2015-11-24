package models

import (
	"database/sql"
	"errors"
	db "github.com/carrot/burrow/db/postgres"
	"time"
)

type Topic struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func AllTopics(limit int64, offset int64) ([]Topic, error) {
	// Query DB
	database := db.Get()
	rows, err := database.Query("SELECT * FROM topics ORDER BY created_at LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Converting rows into []Topic
	var topics []Topic = []Topic{}
	for rows.Next() {
		t := new(Topic)
		err = t.consumeNextRow(rows)

		if err != nil {
			return nil, err
		}

		topics = append(topics, *t)
	}

	// Checking for any errors during iteration
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return topics, nil
}

func (t *Topic) Load(id int64) error {
	database := db.Get()
	row := database.QueryRow("SELECT * FROM topics WHERE id = $1", id)
	return t.consumeRow(row)
}

func (t *Topic) Save() error {
	// Putting into database
	database := db.Get()
	row := database.QueryRow("INSERT INTO topics VALUES(default, $1, default, default) RETURNING *",
		&t.Name,
	)

	// Updating values to match database
	err := t.consumeRow(row)
	if err != nil {
		return err
	}

	return nil
}

func (t *Topic) Update() error {
	// Updating in database
	database := db.Get()
	row := database.QueryRow("UPDATE topics SET name=$1 WHERE id=$2 RETURNING *",
		&t.Name,
		&t.Id,
	)

	// Updating values to match database
	err := t.consumeRow(row)
	if err != nil {
		return err
	}

	return nil
}

func (t *Topic) Delete() error {
	// Deleting from database
	database := db.Get()
	res, err := database.Exec("DELETE FROM topics WHERE id=$1", t.Id)
	if err != nil {
		return err
	}

	numRows, _ := res.RowsAffected()
	if numRows != 1 {
		return errors.New("Nothing was deleted")
	}

	return nil
}

func (t *Topic) consumeRow(row *sql.Row) error {
	return row.Scan(
		&t.Id,
		&t.Name,
		&t.CreatedAt,
		&t.UpdatedAt,
	)
}

func (t *Topic) consumeNextRow(rows *sql.Rows) error {
	return rows.Scan(
		&t.Id,
		&t.Name,
		&t.CreatedAt,
		&t.UpdatedAt,
	)
}
