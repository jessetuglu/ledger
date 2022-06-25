package api

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jessetuglu/bill_app/server/db"
)

type byIntIdRequest struct {
	Id int64 `json:"id" binding:"required"`
}

type createTransactionRequest struct {
	Ledger   uuid.UUID `json:"ledger" binding:"required"`
	Debitor  uuid.UUID `json:"debitor" binding:"required"`
	Creditor uuid.UUID `json:"creditor" binding:"required"`

	Date   time.Time      `json:"date" binding:"required" time_format:"2006-01-02"`
	Amount float64        `json:"amount" binding:"required"`
	Note   sql.NullString `json:"note" binding:"varchar"`
}

func (s *Server) createTransaction(ctx *gin.Context) {
	var req createTransactionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		s.logger.Errorw("Invalid transaction creation request", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, ErrorMessage{err.Error()})
		return
	}
	arg := db.CreateTransactionParams{
		Ledger:   req.Ledger,
		Debitor:  req.Debitor,
		Creditor: req.Creditor,
		Date:     req.Date,
		Amount:   req.Amount,
		Note:     req.Note,
	}
	transaction, err := s.store.CreateTransaction(ctx, arg)
	if err != nil {
		s.logger.Errorw("Couldn't create transaction", err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, ErrorMessage{err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, transaction)
}

func (s *Server) deleteTransaction(ctx *gin.Context) {
	var req byIntIdRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		s.logger.Errorw("Invalid transaction deletion request", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, ErrorMessage{err.Error()})
		return
	}
	err := s.store.DeleteTransaction(ctx, req.Id)
	if err != nil {
		s.logger.Errorw("Couldn't create transaction", err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, ErrorMessage{err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, Message{"Transaction with id: " + strconv.FormatInt(req.Id, 10) + " successfully deleted."})
}
