package models

import (
	db "github.com/carrot/go-base-api/db/redis"
)

type Topic struct {
	Id    int64  `json:"id"`
	Copy  string `json:"copy"`
	Asset string `json:"asset"`
}

// When only dealing with a single record, call on the struct

// AllTopics ([]Topic, error)
// (topic).Find(id) (Topic, error)
// (topic).Create() error
// (topic).Update() error
// (topic).Destroy(id) error

func AllTopics() ([]Topic, error) {
	var res []Topic

	conn := db.Get()
	defer conn.Close()

	return res, nil
}

func (m *Topic) Find(id int64) (Topic, error) {
	conn := db.Get()
	defer conn.Close()

	// Find Individual Record

	return Topic{}, nil
}

func (m *Topic) Create() error {
	conn := db.Get()
	defer conn.Close()

	// Create Record

	return nil
}

func (m *Topic) Update() error {
	conn := db.Get()
	defer conn.Close()

	// Update Record

	return nil
}

func (m *Topic) Destroy() error {
	conn := db.Get()
	defer conn.Close()

	// Delete Record

	return nil
}
