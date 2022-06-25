package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) loadRoutes(){
	s.Router.GET("/", s.rootHandler)
	api_group := s.Router.Group("/api")

	// authentication routes
	api_group.GET("/auth/login", s.googleLoginHandler)
	api_group.GET("/auth/callback", s.googleCallBackHandler)
	
	// user routes
	api_group.GET("/users/:id", s.getUserById) //tmp

	// ledger routes
	api_group.GET("/ledgers/:id", s.getLedgerById) //tmp
	api_group.POST("/ledgers", s.createLedger) //tmp
	api_group.PUT("/ledgers/:id/add_user", s.addUserToLedger) //tmp
	api_group.DELETE("/ledgers/:id", s.deleteLedger) //tmp

	// transaction routes
	api_group.POST("/transactions", s.createTransaction)
	api_group.DELETE("/transactions/:id", s.deleteTransaction)
}

func (s *Server) rootHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, "hello")
}