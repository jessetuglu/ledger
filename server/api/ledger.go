package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jessetuglu/bill_app/server/db"
)

type createLedgerRequest struct {
	Title   string      `json:"title" binding:"required"`
	Members []uuid.UUID `json:"members" binding:"required"`
}

type addUserToLedgerRequest struct {
	ID uuid.UUID `json:"id" binding:"required"`
	User uuid.UUID `json:"user" binding:"required"`
}

func (s *Server) createLedger(ctx *gin.Context) {
	var req createLedgerRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorMessage{err.Error()})
		return
	}

	arg := db.CreateLedgerParams{
		Title:   req.Title,
		Members: req.Members,
	}

	ledger, err := s.store.CreateLedger(ctx, arg)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, ErrorMessage{err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, ledger)
}

func (s *Server) getLedgerById(ctx *gin.Context) {
	param_id, found := ctx.Params.Get("id")
	if (!found) {
		s.logger.Error("Id param not found")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, ErrorMessage{"Invalid ledger ID"})
		return
	}
	ledger_id, _ := uuid.Parse(param_id)
	ledger, err := s.store.GetLedgerById(ctx, ledger_id)
	if err != nil {
		s.logger.Error("Couldn't get user: ", err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, ErrorMessage{"Ledger not found"})
		return
	}
	ctx.JSON(http.StatusOK, ledger)
}

func (s *Server) deleteLedger(ctx *gin.Context) {
	param_id, found := ctx.Params.Get("id")
	if (!found) {
		s.logger.Error("Id param not found")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, ErrorMessage{"Invalid ledger ID"})
		return
	}
	ledger_id, _ := uuid.Parse(param_id)
	if err := s.store.DeleteLedger(ctx, ledger_id); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, ErrorMessage{err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, Message{"Ledger with id: " + ledger_id.String() + " successfully deleted."})
}

func (s *Server) addUserToLedger(ctx *gin.Context) {
	var req addUserToLedgerRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorMessage{err.Error()})
		return
	}

	arg := db.AddUserToLedgerParams{
		ID:   req.ID,
		User: req.User,
	}

	if err := s.store.AddUserToLedger(ctx, arg); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, ErrorMessage{err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, Message{"User successfully added to ledger: " + req.ID.String()})
}
