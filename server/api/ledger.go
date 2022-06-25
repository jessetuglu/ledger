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
	ID   uuid.UUID `json:"id" binding:"required"`
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
	var req byUUIDRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorMessage{err.Error()})
		return
	}

	ledger, err := s.store.GetLedgerById(ctx, req.Id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, Message{err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, ledger)
}

func (s *Server) deleteLedger(ctx *gin.Context) {
	var req byUUIDRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, Message{err.Error()})
		return
	}

	if err := s.store.DeleteLedger(ctx, req.Id); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, ErrorMessage{err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, Message{"Ledger with id: " + req.Id.String() + " successfully deleted."})
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
