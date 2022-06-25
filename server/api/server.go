package api

import (
	"database/sql"
	"os"

	"github.com/antonlindstrom/pgstore"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/jessetuglu/bill_app/server/db"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

type Server struct {
	store       *db.Store
	Router      *gin.Engine
	logger      *zap.SugaredLogger
	sessions    *pgstore.PGStore
	oauthConfig oauth2.Config

	serverBaseUrl string
	clientBaseUrl string
}

func NewServer(logger *zap.SugaredLogger, conn *sql.DB, config oauth2.Config) *Server {
	var s Server
	var err error
	s.logger = logger
	s.store = db.NewStore(conn)
	s.oauthConfig = config

	s.sessions, err = pgstore.NewPGStoreFromPool(conn, []byte(os.Getenv("SESSIONS_KEY")))
	if err != nil {
		s.logger.Fatalw("Unable to initialize sessions db", "Error:", err)
	}

	s.sessions.Options = &sessions.Options{
		HttpOnly: true,
		Secure:   true,
		MaxAge:   60 * 60 * 24 * 30, // 30 days
		Path:     "/",
	}

	s.Router = gin.Default()

	if os.Getenv("MODE") == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	s.Router.Use(gin.Recovery())
	s.loadRoutes()

	return &s
}
