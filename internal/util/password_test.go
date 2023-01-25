package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPassword(t *testing.T) {
	pass := RandomString(12)

	hash1, err := HashPassword(pass)
	require.NoError(t, err)
	require.NotEmpty(t, hash1)

	err = CheckPassword(pass, hash1)
	require.NoError(t, err)

	wrongPass := RandomString(12)
	err = CheckPassword(wrongPass, hash1)
	require.Error(t, err)
	require.EqualError(t, err, "crypto/bcrypt: hashedPassword is not the hash of the given password")
}
