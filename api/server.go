package api

import (
	db "github.com/SabariGanesh-K/prod-mgm-go/db/sqlc"
	"github.com/SabariGanesh-K/prod-mgm-go/util"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config util.Config
	store db.Store
	router *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server ,error){
	server := &Server{
config: config,
store: store,
	}
	server.setupRouter()
	return server,nil
	
}

func (server *Server) setupRouter() {
	router:= gin.Default()
	router.POST("/users",server.createUser)
	router.POST("/users/login",server.loginUser)
	router.POST("/products",server.createProduct)
	// router.GET("/products",server.getProductByUserID)
	router.GET("/products/:id",server.getProductByProductID)
	router.POST("/products/addcompressed",server.addCompressedImagesByProductID)


	server.router=router
}


func (server *Server) Start(address string) error {

	return server.router.Run(address)
}
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
