package cache

import (
	"time"

	"github.com/dgraph-io/badger/v3"
)

type BadgerCache struct {
	Conn   *badger.DB
	Prefix string
}

func (b *BadgerCache) Has(str string) (bool, error) {
	_, err := b.Get(str)
	if err != nil {
		return false, err
	}

	return true, nil
}
func (b *BadgerCache) Get(str string) (any, error) {
	var fromCache []byte

	err := b.Conn.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(str))
		if err != nil {
			return err
		}

		err = item.Value(func(val []byte) error {
			fromCache = append([]byte{}, val...)
			return nil
		})
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	decoded, err := decode(string(fromCache))
	if err != nil {
		return nil, err
	}

	item := decoded[str]

	return item, nil
}
func (b *BadgerCache) Set(str string, value any, expires ...int) error {
	entry := Entry{}
	entry[str] = value
	encoded, err := encode(entry)
	if err != nil {
		return err
	}

	newEntry := badger.NewEntry([]byte(str), encoded)

	return b.Conn.Update(func(txn *badger.Txn) error {
		if len(expires) > 0 {
			err = txn.SetEntry(newEntry.WithTTL(time.Duration(expires[0]) * time.Second))
		} else {
			err = txn.SetEntry(newEntry)
		}
		err = txn.SetEntry(newEntry)
		return err
	})
}

func (b *BadgerCache) Forget(str string) error {
	return b.Conn.Update(func(txn *badger.Txn) error {
		err := txn.Delete([]byte(str))
		return err
	})
}
func (b *BadgerCache) EmptyByMatch(str string) error
func (b *BadgerCache) Empty() error

func (b *BadgerCache) makeKey(str string) string
func (b *BadgerCache) getKeys(pattern string) ([]string, error)

func (b *BadgerCache) emptyByMatch(str string) error