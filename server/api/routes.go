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
	api_group.GET("/auth/whoami", s.validLoginMiddleware(), s.getCurrentUser)
	
	// user routes
	api_group.GET("/users/:id", s.validLoginMiddleware(), s.getUserById) //tmp
	api_group.GET("/users/:id/ledgers", s.validLoginMiddleware(), s.getUserLedgers) //tmp

	// ledger routes
	api_group.GET("/ledgers/:id", s.validLoginMiddleware(), s.getLedgerById) //tmp
	api_group.POST("/ledgers", s.validLoginMiddleware(), s.createLedger) //tmp
	api_group.PUT("/ledgers/add_user", s.validLoginMiddleware(), s.addUserToLedger) //tmp
	api_group.DELETE("/ledgers/:id", s.validLoginMiddleware(), s.deleteLedger) //tmp

	// transaction routes
	api_group.POST("/transactions", s.validLoginMiddleware(), s.createTransaction)
	api_group.DELETE("/transactions/:id", s.validLoginMiddleware(), s.deleteTransaction)
}

func (s *Server) rootHandler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, "hello")
}