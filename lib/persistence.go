package affix

import (
	"fmt"
	"log"

	"github.com/dgraph-io/badger"
)

type PersistenceHandler interface {
	Init() error
	Set(key string, val []byte) error
	Get(key string) ([]byte, error)
}

type BadgerKV struct {
	db *badger.DB
}

func (b *BadgerKV) Init() error {
	opts := badger.DefaultOptions
	opts.Dir = "badger"
	opts.ValueDir = "badger"
	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	b.db = db
	return nil
}

func (b *BadgerKV) Set(key string, val []byte) error {
	return b.db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(key), val)
		return err
	})

}

func (b *BadgerKV) Get(key string) (e error, val []byte) {
	e = b.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("answer"))
		if err != nil {
			return err
		}
		val, err = item.Value()
		if err != nil {
			return err
		}
		fmt.Printf("The answer is: %s\n", val)
		return nil
	})
	return
}

func main() {
	// Open the Badger database located in the /tmp/badger directory.
	// It will be created if it doesn't exist.
	opts := badger.DefaultOptions
	opts.Dir = "/tmp/badger"
	opts.ValueDir = "/tmp/badger"
	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// Your code here…
}
