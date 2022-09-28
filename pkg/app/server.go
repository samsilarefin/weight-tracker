package app

import (
	"log"
	"weight-tracker/pkg/api"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router        *gin.Engine
	userService   api.UserService
	weightService api.WeightService
}

func NewServer(router *gin.Engine, userService api.UserService, weightService api.WeightService) *Server {
	return &Server{
		router:        router,
		userService:   userService,
		weightService: weightService,
	}
}

func (s *Server) Run() error {
	// run function that initializes the routes
	r := s.Routes()

	// run the server through the router
	err := r.Run()

	if err != nil {
		log.Printf("Server - there was an error calling Run on router: %v", err)
		return err
	}

	return nil
}
