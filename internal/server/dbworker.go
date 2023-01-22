package server

type DBError struct {
	s string
}

func (ae *DBError) Error() string {
	return ae.s
}

func NewDBError(text string) error {
	return &DBError{text}
}

func saveNewUser(u *User) error {
	return NewDBError("User registration is not implemented in DB yet.")
}
