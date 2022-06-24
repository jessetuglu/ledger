package server

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)


func InitializeServer() (*Server){

}

func (s *Server) load_routes(){
	// handle auth routes
	// s.router.HandleFunc("/auth/google", )
	// handle bill routes
	sql.Tx
}

