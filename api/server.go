package api

import (
	db "github.com/alikhanMuslim/Catalog-service/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	store  *db.Store
}

func NewServer(store db.Store) *Server {
	server := &Server{}

	server.store = &store

	router := gin.New()

	server.router = router

	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
