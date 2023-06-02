package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/transaction-organizer/repository"
	"net/http"
	"time"
)

type ReportController struct{}

func (rc ReportController) GetGroupSummaryReport(c *gin.Context) {
	fromDate, err := time.Parse("2006-01-02", c.Query("fromDate"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	toDate, err := time.Parse("2006-01-02", c.Query("toDate"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	toEndDate := time.Date(toDate.Year(), toDate.Month(), toDate.Day(), 23, 59, 59, 0, toDate.Location())

	c.IndentedJSON(http.StatusOK, repository.TransactionRepository{}.GetGroupSummaryReport(fromDate, toEndDate))
}
