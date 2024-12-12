package api

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	db "github.com/SabariGanesh-K/prod-mgm-go/db/sqlc"
	"github.com/SabariGanesh-K/prod-mgm-go/util"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{
		// TokenSymmetricKey:   util.RandomString(32),
		// AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
