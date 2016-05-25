package models

import (
	"database/sql"
	"github.com/BrandonRomano/serf"
	db "github.com/carrot/burrow/db/postgres"
	"time"
)

type Topic struct {
	serf.Worker `json:"-"`
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewTopic() *Topic {
	topic := new(Topic)
	return topic.Prep()
}

func (t *Topic) Prep() *Topic {
	t.Worker = &serf.PqWorker{
		Database: db.Get(),
		Config: serf.Configuration{
			TableName: "topics",
			Fields: []serf.Field{
				serf.Field{Pointer: &t.Id, Name: "id", UniqueIdentifier: true,
					IsSet: func(pointer interface{}) bool {
						pointerInt := *pointer.(*int64)
						return pointerInt != 0
					},
				},
				serf.Field{Pointer: &t.Name, Name: "name", Insertable: true, Updatable: true},
				serf.Field{Pointer: &t.CreatedAt, Name: "created_at"},
				serf.Field{Pointer: &t.UpdatedAt, Name: "updated_at"},
			},
		},
	}
	return t
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

func (t *Topic) consumeNextRow(rows *sql.Rows) error {
	return rows.Scan(
		&t.Id,
		&t.Name,
		&t.CreatedAt,
		&t.UpdatedAt,
	)
}
