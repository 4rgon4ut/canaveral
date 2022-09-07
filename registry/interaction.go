package registry

import (
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/ethereum/go-ethereum/common"
)

// Adds contracts <name> : <address> pair to the persistent registry
func (r *Registry) AddContract(name string, addr common.Address) error {
	err := r.Update((func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Contracts"))
		if err := b.Put([]byte(name), []byte(addr.Hex())); err != nil {
			return fmt.Errorf("write transaction error: %w", err)
		}
		return nil
	}))
	return err
}

// Gets contract address from persistent registry by name
func (r *Registry) GetAddress(name string) (addr string, err error) {
	err = r.View((func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Contracts"))
		a := b.Get([]byte(name))
		if a == nil {
			return fmt.Errorf("no such contract")
		}
		addr = string(a)
		return nil
	}))
	return
}
