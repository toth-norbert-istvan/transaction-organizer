package service

import (
	"fmt"
	"github.com/transaction-organizer/domain"
	"github.com/xuri/excelize/v2"
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

type KhFileParserService struct{}

func (kfps KhFileParserService) GetTransactionsFromExcelFile(file *multipart.FileHeader) []domain.Transaction {
	f, err := file.Open()
	if err != nil {
		fmt.Println("Error during excel file access:", err)
		return nil
	}

	excelFile, err := excelize.OpenReader(f)

	cellIndex := 2
	var newTransactions []domain.Transaction
	for {
		partnerName := kfps.getCellStringValue(excelFile, fmt.Sprintf("%s%d", PARTNER_NAME_COLUMN, cellIndex))
		if len(partnerName) > 0 {
			newTransaction := domain.Transaction{
				Partner: partnerName,
				Amount:  kfps.getCellFloatValue(excelFile, fmt.Sprintf("%s%d", AMOUNT_COLUMN, cellIndex)),
				Date:    kfps.getCellTimeValue(excelFile, fmt.Sprintf("%s%d", TRANSACTION_DATE_COLUMN, cellIndex)),
			}
			newTransactions = append(newTransactions, newTransaction)
			cellIndex++
		} else {
			break
		}
	}

	return newTransactions
}

func (kfps KhFileParserService) getCellStringValue(excelFile *excelize.File, cell string) string {
	cellValue, _ := excelFile.GetCellValue(EXCEL_SHEET_NAME, cell)
	return strings.TrimSpace(cellValue)
}

func (kfps KhFileParserService) getCellFloatValue(excelFile *excelize.File, cell string) float64 {
	cellValue, _ := excelFile.GetCellValue(EXCEL_SHEET_NAME, cell)
	floatValue, _ := strconv.ParseFloat(cellValue, 64)
	return floatValue
}

func (kfps KhFileParserService) getCellTimeValue(excelFile *excelize.File, cell string) time.Time {
	cellValue, _ := excelFile.GetCellValue(EXCEL_SHEET_NAME, cell, excelize.Options{RawCellValue: true})
	var days, _ = strconv.Atoi(cellValue)
	return excelEpoch.Add(time.Second * time.Duration(days*86400))
}
