package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	idContextKey = "user_id"
	emailContextKey = "user_email"
	firstNameContextKey = "user_first_name"
	lastNameContextKey = "user_last_name"
	sessionContextKey = "session"
)

func requestIdInserter() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("X-Request-Id", uuid.New().String())
    	c.Next()
	}
}

func (s *Server) validLoginMiddleware() gin.HandlerFunc{
	return func(ctx *gin.Context){
		session, err := s.sessions.Get(ctx.Request, "session")
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, ErrorMessage{"Couldn't find a valid session for you."})
			return
		}

		email, ok := session.Values["email"].(string)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, ErrorMessage{"Couldn't find a valid session for you."})
			return
		}

		first_name, ok := session.Values["first_name"].(string)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, ErrorMessage{"Couldn't find a valid session for you."})
			return
		}

		last_name, ok := session.Values["last_name"].(string)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, ErrorMessage{"Couldn't find a valid session for you."})
			return
		}

		id, ok := session.Values["id"].(string)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, ErrorMessage{"Couldn't find a valid session for you."})
			return
		}

		ctx.Set(emailContextKey, email)
		ctx.Set(idContextKey, id)
		ctx.Set(firstNameContextKey, first_name)
		ctx.Set(lastNameContextKey, last_name)
		
		ctx.Next()
	}
}
