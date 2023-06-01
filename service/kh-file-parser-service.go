package transaction_organizer_domain_service

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"gitlab.com/transaction-organizer/domain"
	"mime/multipart"
	"strconv"
	"strings"
	"time"
)

const EXCEL_SHEET_NAME = "Könyvelt tételek"
const PARTNER_NAME_COLUMN = "G"
const AMOUNT_COLUMN = "H"
const TRANSACTION_DATE_COLUMN = "A"

var excelEpoch = time.Date(1899, time.December, 30, 0, 0, 0, 0, time.UTC)

func GetTransactionsFromExcelFile(file *multipart.FileHeader) []transaction_organizer_domain.Transaction {
	f, err := file.Open()
	if err != nil {
		fmt.Println("Error during excel file access:", err)
		return nil
	}

	excelFile, err := excelize.OpenReader(f)

	cellIndex := 2
	var newTransactions []transaction_organizer_domain.Transaction
	for {
		partnerName := getCellStringValue(excelFile, fmt.Sprintf("%s%d", PARTNER_NAME_COLUMN, cellIndex))
		if len(partnerName) > 0 {
			newTransaction := transaction_organizer_domain.Transaction{
				Partner: partnerName,
				Amount:  getCellFloatValue(excelFile, fmt.Sprintf("%s%d", AMOUNT_COLUMN, cellIndex)),
				Date:    getCellTimeValue(excelFile, fmt.Sprintf("%s%d", TRANSACTION_DATE_COLUMN, cellIndex)),
			}
			newTransactions = append(newTransactions, newTransaction)
			cellIndex++
		} else {
			break
		}
	}

	return newTransactions
}

func getCellStringValue(excelFile *excelize.File, cell string) string {
	cellValue, _ := excelFile.GetCellValue(EXCEL_SHEET_NAME, cell)
	return strings.TrimSpace(cellValue)
}

func getCellFloatValue(excelFile *excelize.File, cell string) float64 {
	cellValue, _ := excelFile.GetCellValue(EXCEL_SHEET_NAME, cell)
	floatValue, _ := strconv.ParseFloat(cellValue, 64)
	return floatValue
}

func getCellTimeValue(excelFile *excelize.File, cell string) time.Time {
	cellValue, _ := excelFile.GetCellValue(EXCEL_SHEET_NAME, cell, excelize.Options{RawCellValue: true})
	var days, _ = strconv.Atoi(cellValue)
	return excelEpoch.Add(time.Second * time.Duration(days*86400))
}
