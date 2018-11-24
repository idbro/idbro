package bbolt

import (
	"errors"
	"time"

	"github.com/idbro/idbro/storage"
	bolt "go.etcd.io/bbolt"
)

// Storage for bbolt
type Storage struct {
	Path string
	db   *bolt.DB
}

var (
	// ErrValueTypeNotBytes is returned if the meta value is not []byte.
	ErrValueTypeNotBytes = errors.New("bbolt storage needs value to be bytes only")
)

// New to create the Storage instance
func New(path string) *Storage {
	return &Storage{
		Path: path,
		db:   nil,
	}
}

// Initialize the bolt DB
func (s *Storage) Initialize() error {
	// create a db instance with a timeout of 1 sec
	dbInstance, err := bolt.Open(s.Path, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}

	s.db = dbInstance
	return nil
}

// Close to clean up the database
func (s *Storage) Close() error {
	return s.db.Close()
}

// Insert a new record with given meta data.
// Meta value only supports []byte type.
func (s *Storage) Insert(id []byte, meta map[string]interface{}) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		// create a bucket for the given id
		bkt, err := tx.CreateBucket(id)
		if err != nil {
			// check whether it is because of id already exists
			if err == bolt.ErrBucketExists {
				return storage.ErrRecordExists
			}
			return err
		}

		// put meta data into the bucket
		for k, v := range meta {
			value, ok := v.([]byte)
			if !ok {
				return ErrValueTypeNotBytes
			}

			if err := bkt.Put([]byte(k), value); err != nil {
				return err
			}
		}

		return nil
	})
}

// Get the record in DB
func (s *Storage) Get(id []byte) (*storage.Record, error) {
	var record storage.Record

	err := s.db.View(func(tx *bolt.Tx) error {
		// create a bucket for the given id
		bkt := tx.Bucket(id)
		if bkt == nil {
			return storage.ErrRecordNotExists
		}

		meta := make(map[string]interface{}, 0)
		bkt.ForEach(func(k, v []byte) error {
			meta[string(k)] = v
			return nil
		})
		record.ID = id
		record.Meta = meta
		return nil
	})

	if err != nil {
		return nil, err
	}
	return &record, nil
}
