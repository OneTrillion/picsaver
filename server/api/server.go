package api

import (
	"picsaver/server/db/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct {
    store db.Store
    router *gin.Engine
}

func NewServer(store db.Store) *Server {
    server := &Server{store: store}
    router := gin.Default()

    // Routes
    router.GET("/google_login", server.GoogleLogin)
    router.GET("/google_callback", server.GoogleCallback)

    server.router = router
    return server
}

func (server *Server) Start(address string) error {
    return server.router.Run(address) 
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

