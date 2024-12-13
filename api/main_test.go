package api

import (
	"os"
	"testing"

	db "github.com/SabariGanesh-K/prod-mgm-go/db/sqlc"
	"github.com/SabariGanesh-K/prod-mgm-go/util"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)


func newTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{
		
	}
	

	
	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}

func newBenchMarkTestServer(b *testing.B,store db.Store) *Server{
	config := util.Config{
		
	}
	

	
	server, err := NewServer(config, store)
	require.NoError(b, err)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
