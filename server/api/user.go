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
	if (ctx.Value("user_id").(string) != param_id){
		s.logger.Errorw("Unauthorized to access this user")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, ErrorMessage{"Oops. You can only access info about yourself!"})
		return
	}
	user_id, _ := uuid.Parse(param_id)
	user, err := s.store.GetUserById(ctx, user_id)
	if err != nil {
		s.logger.Errorw("Couldn't get user: ", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, ErrorMessage{"User not found"})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (s *Server) getUserLedgers(ctx *gin.Context){
	param_id, found := ctx.Params.Get("id")
	if (!found) {
		s.logger.Errorw("Id param not found")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, ErrorMessage{"No user id found"})
		return
	}
	if (ctx.Value("user_id").(string) != param_id){
		s.logger.Errorw("Unauthorized to access this user")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, ErrorMessage{"Oops. You can only access info about yourself!"})
		return
	}
	user_id, _ := uuid.Parse(param_id)
	ledgers, err := s.store.GetUserLedgers(ctx, user_id)
	if err != nil {
		s.logger.Errorw("Couldn't get user ledgers: ", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, ErrorMessage{"Couldn't get ledgers for user with id: " + param_id})
		return
	}
	ctx.JSON(http.StatusOK, ledgers)
}
