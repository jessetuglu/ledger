package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (s *Server) getUserById(ctx *gin.Context){
	param_id, found := ctx.Params.Get("id")
	if (!found) {
		s.logger.Errorw("Id param not found")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, ErrorMessage{"No user id found"})
		return
	}
	user_id, err := uuid.Parse(param_id)
	user, err := s.store.GetUserById(ctx, user_id)
	if err != nil {
		s.logger.Errorw("Couldn't get user: ", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, ErrorMessage{"User not found"})
		return
	}
	ctx.JSON(http.StatusOK, user)
}
