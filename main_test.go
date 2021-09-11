package main

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func Test2Plus2Equals4(t *testing.T) {
	actual := add(2, 2)
	assert.Equal(t, 4, actual)
}

func TestConnectToDatabase(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}
	ctx := context.Background()
	db, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	require.NoError(t, err)
	defer db.Close(ctx)
	assert.NoError(t, db.Ping(ctx))
}
