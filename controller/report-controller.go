package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/transaction-organizer/repository"
	"net/http"
	"time"
)

type ReportController struct{}

func (rc ReportController) GetGroupSummaryReport(c *gin.Context) {
	fromDate, toEndDate, err := getFromDateAndToEndDate(c)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.IndentedJSON(http.StatusOK, repository.TransactionRepository{}.GetGroupSummaryReport(fromDate, toEndDate))
}

func (rc ReportController) GetSummaryReportByTransactionType(c *gin.Context) {
	fromDate, toEndDate, err := getFromDateAndToEndDate(c)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	c.IndentedJSON(http.StatusOK, repository.TransactionRepository{}.GetSummaryReportByTransactionType(fromDate, toEndDate))
}

func getFromDateAndToEndDate(c *gin.Context) (time.Time, time.Time, error) {
	fromDate, err := time.Parse("2006-01-02", c.Query("fromDate"))
	if err != nil {
		return fromDate, fromDate, err
	}
	toDate, err := time.Parse("2006-01-02", c.Query("toDate"))
	if err != nil {
		return fromDate, toDate, err
	}
	toEndDate := time.Date(toDate.Year(), toDate.Month(), toDate.Day(), 23, 59, 59, 0, toDate.Location())

	return fromDate, toEndDate, err
}
