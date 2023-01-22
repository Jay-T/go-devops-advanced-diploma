package server

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func seedUsers(userStore UserStore) error {
	err := createUser(userStore, "admin", "admin")
	if err != nil {
		return err
	}

	return createUser(userStore, "admin2", "admin2")
}

func createUser(userStore UserStore, username string, password string) error {
	user, err := NewUser(username, password)
	if err != nil {
		return err
	}

	return userStore.Save(user)
}

type User struct {
	Username string
	Passhash string
}

func NewUser(username string, password string) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("cannot hash password: %w", err)
	}

	user := &User{
		Username: username,
		Passhash: string(hashedPassword),
	}

	return user, nil
}

func (u *User) IsCorrectPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Passhash), []byte(password))
	return err == nil
}

func (u *User) Clone() *User {
	return &User{
		Username: u.Username,
		Passhash: u.Passhash,
	}
}

// ------------------------------------------------------------
// ------------------------------------------------------------

// func (u *User) authenticateUser(ctx context.Context, tokenLifeTime time.Duration) (string, error) {
// 	u.getPassHash()

// 	if u.Passhash != "20d7547437a19533e17d525e129e67308e74c88efe936639bc4decab263f3e66" {
// 		return "", NewAuthenticationError("Credentials are invalid.")
// 	}

// 	token, err := u.getToken(tokenLifeTime)
// 	if err != nil {
// 		return "", err
// 	}

// 	return token, nil
// }

// func (u *User) registerNewUser(ctx context.Context) error {
// 	u.getPassHash()

// 	// Добавление нового пользователя в базу.
// 	err := saveNewUser(u)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
