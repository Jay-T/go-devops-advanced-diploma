package server

import (
	"errors"
	"sync"

	db "github.com/Jay-T/go-devops-advanced-diploma/db/sqlc"
)

var ErrAlreadyExists = errors.New("record already exists")

type AccountStore interface {
	Save(user *db.Account) error
	Find(username string) (*db.Account, error)
}

type InMemoryAccountStore struct {
	mutex    sync.RWMutex
	accounts map[string]*db.Account
}

func NewInMemoryAccountStore() *InMemoryAccountStore {
	return &InMemoryAccountStore{
		accounts: make(map[string]*db.Account),
	}
}

func (store *InMemoryAccountStore) Save(account *db.Account) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.accounts[account.Username] != nil {
		return ErrAlreadyExists
	}

	store.accounts[account.Username] = account.Clone()
	return nil
}

func (store *InMemoryAccountStore) Find(username string) (*db.Account, error) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	acc := store.accounts[username]

	if acc == nil {
		return nil, nil
	}

	return acc.Clone(), nil
}
