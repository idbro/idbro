package bbolt

import (
	bolt "go.etcd.io/bbolt"
)

// Storage for bbolt
type Storage struct {
	Path string
	db   *bolt.DB
}

// New to create the Storage instance
func New(path string) *Storage {
	dbInstance, err := bolt.Open(path, 0600, nil)
	if err != nil {
		return nil
	}

	return &Storage{
		Path: path,
		db:   dbInstance,
	}
}

// Initialize the bolt DB
func (s *Storage) Initialize() error {
	return nil
}

// Close to clean up the database
func (s *Storage) Close() error {
	return nil
}
