package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/Jay-T/go-devops-advanced-diploma/internal/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	hash, err := util.HashPassword(util.RandomPass())
	require.NoError(t, err)

	arg := CreateAccountParams{
		Username: util.RandomUser(),
		Passhash: hash,
	}

	ctx := context.TODO()

	acc, err := testQueries.CreateAccount(ctx, arg)

	require.NoError(t, err)
	require.Equal(t, arg.Username, acc.Username)
	require.Equal(t, arg.Passhash, acc.Passhash)

	return acc
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestBlockAccount(t *testing.T) {
	acc := createRandomAccount(t)

	err := testQueries.BlockAccount(context.TODO(), acc.Username)
	require.NoError(t, err)

	acc1, err := testQueries.GetAccount(context.TODO(), acc.Username)
	require.Equal(t, true, acc1.Blocked)
	require.NoError(t, err)
}

func TestDeleteAccount(t *testing.T) {
	acc := createRandomAccount(t)

	err := testQueries.DeleteAccount(context.TODO(), acc.Username)
	require.NoError(t, err)

	acc1, err := testQueries.GetAccount(context.TODO(), acc.Username)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, acc1)
}
