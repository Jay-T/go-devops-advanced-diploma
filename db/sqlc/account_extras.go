package db

import "golang.org/x/crypto/bcrypt"

func (a *Account) Clone() *Account {
	return &Account{
		Username: a.Username,
		Passhash: a.Passhash,
	}
}

func (a *Account) IsCorrectPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(a.Passhash), []byte(password))
	return err == nil
}
