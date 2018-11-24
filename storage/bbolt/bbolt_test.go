package bbolt

import (
	"os"
	"testing"

	"github.com/idbro/idbro/storage"
	"github.com/stretchr/testify/assert"
)

const (
	path = "./db.bolt"
)

func TestBBoltDBOperations(t *testing.T) {
	boltStorage := New(path)

	// Assert to Storage interface and check whether the interface is implemented.
	s, ok := interface{}(boltStorage).(storage.Storage)
	assert.True(t, ok, "boltStorage should implement storage.Storage interface")

	// Initialize
	assert.NoError(t, s.Initialize(), "should not have error when initializing the db")

	record1ID := []byte("record1")

	record1MetaKey1 := "name"
	record1MetaValue1 := []byte("I am record1")
	record1MetaKey2 := "email"
	record1MetaValue2 := []byte("record1@idbro.email")

	record1Meta := make(map[string]interface{}, 0)
	record1Meta[record1MetaKey1] = record1MetaValue1
	record1Meta[record1MetaKey2] = record1MetaValue2

	// insert record1
	assert.NoError(t, s.Insert(record1ID, record1Meta), "should not have error when inserting record1")

	// insert record1 again
	assert.Equal(t, storage.ErrRecordExists, s.Insert(record1ID, record1Meta),
		"insert record1 again will have ErrRecordExists")

	// Test insert record with a bad Meta data that has not []byte value
	recordBadID := []byte("recordBad")
	recordBadMeta := make(map[string]interface{}, 0)
	recordBadMeta["key1"] = 12
	assert.Equal(t, ErrValueTypeNotBytes, s.Insert(recordBadID, recordBadMeta),
		"insert record with bad Meta data having not []byte value will have ErrValueTypeNotBytes")

	// Get back record1
	recordResult, err := s.Get(record1ID)
	assert.NoError(t, err, "should not have error when getting back record1")
	assert.EqualValues(t, record1ID, recordResult.ID, "record1ID should match")
	assert.EqualValues(t, record1Meta, recordResult.Meta, "record1 Meta data should match")

	// Get back the bad record
	recordResult, err = s.Get(recordBadID)
	assert.Equal(t, storage.ErrRecordNotExists, err, "the bad record should not exist")
	assert.Nil(t, recordResult, "the bad record result should be nil")

	// Close and delete the db file.
	assert.NoError(t, s.Close(), "should not have error when closing the db")
	assert.NoError(t, os.Remove(path), "should not have error when removing the db file")
}
