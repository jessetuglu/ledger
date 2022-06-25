package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	emailContextKey = "user_email"
	firstNameContextKey = "user_first_name"
	fullNameContextKey = "user_full_name"
	sessionContextKey = "sesh"
)

func requestIdInserter() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("X-Request-Id", uuid.New().String())
    	c.Next()
	}
}

func (s *Server) validLoginMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context){
		session, err := s.sessions.Get(ctx.Request, "session")
		if (err != nil){
			s.logger.Infow("Invalid Session", err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, "Try logging in again.")
			return
		}
		_, ok := session.Values["email"].(string)
		if !ok {
			s.logger.Infow("Missing email", err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, "Make sure you are signed in")
			return
		}
		ctx.Next()
	}
}
