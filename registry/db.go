package registry

import (
	"fmt"
	"io/fs"

	"github.com/boltdb/bolt"
)

type Registry struct {
	*bolt.DB
}

func New(path string, mode fs.FileMode) (*Registry, error) {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open(path, mode, nil)
	if err != nil {
		return nil, fmt.Errorf("bolt db open error: %w", err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("Contracts"))
		if err != nil {
			return fmt.Errorf("create bucket: %w", err)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &Registry{db}, nil
	// TODO:
	// defer db.Close()
}
