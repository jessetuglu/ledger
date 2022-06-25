package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) getUserById(ctx *gin.Context) {
	var req byUUIDRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		s.logger.Errorw("Couldn't parse request: ", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, ErrorMessage{err.Error()})
		return
	}
	user, err := s.store.GetUserById(ctx, req.Id)
	if err != nil {
		s.logger.Errorw("Couldn't get user ", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, ErrorMessage{err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}
