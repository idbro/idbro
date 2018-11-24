package storage

import "errors"

type (
	// Record is a piece of information with ID and Meta data.
	Record struct {
		ID   []byte
		Meta map[string]interface{}
	}
)

var (
	// ErrRecordExists is returned when creating a record that already exists.
	ErrRecordExists = errors.New("record already exists")

	// ErrRecordNotExists is returned when trying to access a record which not exists.
	ErrRecordNotExists = errors.New("record not exists")
)

// Storage interface is used to define behaviors for storage.
type Storage interface {
	// Initialize the storage. Used for creating the db file, ping db server and related initialization jobs.
	Initialize() error

	// Close to clean up the database, connections and so on.
	Close() error

	// Insert a new record with given meta data.
	Insert(id []byte, meta map[string]interface{}) error

	// Get a record with its ID
	Get(id []byte) (*Record, error)
}
