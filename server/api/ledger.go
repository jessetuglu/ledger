package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jessetuglu/bill_app/server/db"
)

type createLedgerRequest struct {
	Title string `json:"title" binding:"required"`
	Members []uuid.UUID `json:"members" binding:"required"`
}

func (s *Server) createLedger(ctx *gin.Context){
	var req createLedgerRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	arg := db.CreateLedgerParams{
		Title: req.Title,
		Members: req.Members,
	}

	ledger, err := s.store.CreateLedger(ctx, arg)
	if (err != nil){
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}
	ctx.JSON(http.StatusOK, ledger)
}