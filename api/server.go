package api

import (
	"api.namegame.com/api/routes"
	"github.com/gin-gonic/gin"
)

//Server responsible for manage API Rest bindings
type Server struct {
	Routes []routes.BaseRoute
}

//SetupRoutes init API endpoints
func (s *Server) SetupRoutes() *gin.Engine {
	router := gin.Default()
	for _, route := range s.Routes {
		route.Setup(router)
	}
	return router
}
