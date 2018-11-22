package storage

// Storage interface is used to define behaviors for storage.
type Storage interface {
	// Initialize the storage
	Initialize() error

	// Close to clean up the database
	Close() error
}
