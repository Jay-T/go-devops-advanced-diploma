package server

// TODO(): remove
// Legacy code for inmemorystore use

// import (
// 	"fmt"

// 	db "github.com/Jay-T/go-devops-advanced-diploma/db/sqlc"
// 	"golang.org/x/crypto/bcrypt"
// )

// func seedAccounts(accountStore AccountStore) error {
// 	err := createAccount(accountStore, "admin", "admin")
// 	if err != nil {
// 		return err
// 	}

// 	return createAccount(accountStore, "admin2", "admin2")
// }

// func createAccount(accountStore AccountStore, username string, password string) error {
// 	acc, err := NewAccount(username, password)
// 	if err != nil {
// 		return err
// 	}

// 	return accountStore.Save(acc)
// }

// func NewAccount(username string, password string) (*db.Account, error) {
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return nil, fmt.Errorf("cannot hash password: %w", err)
// 	}

// 	acc := &db.Account{
// 		Username: username,
// 		Passhash: string(hashedPassword),
// 	}

// 	return acc, nil
// }
