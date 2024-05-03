package service

import (
	// "log"

	"database/sql"
	"fmt"
	"strings"

	// "fmt"
	"mime/multipart"

	r "bam/internal/app/repository"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type ExcelService struct {
	repo r.IExcelRepository
}

func NewExcelService(repo r.IExcelRepository) *ExcelService {
	return &ExcelService{repo: repo}
}

// Implement the ReadExcelFile method in your ProductService struct
func (s *ExcelService) ReadExcel(file *multipart.FileHeader) ([]string, error) {
	rows, err := s.repo.ReadExcel(file)
	if err != nil {
		return nil, err
	}

	var result []string
	for _, row := range rows {
		result = append(result, strings.Join(row, ","))
	}

	return result, nil
}

func (s *ExcelService) ExportDataToExcel() (string, error) {
	// Query data from the database
	var (
		rows *sql.Rows // Assuming sql.Rows for database result (might need adjustment)
		err  error
	)

	// Replace with your actual logic to query data based on needs
	// This assumes GetExcelData returns a queryable object
	rows, err = s.repo.GetExcelData(nil).Rows()
	if err != nil {
		return "", err
	}
	defer rows.Close()

	// Create a new Excel file
	file := excelize.NewFile()

	// Add a new sheet
	index := file.NewSheet("Sheet1")

	// Set the header row
	file.SetCellValue("Sheet1", "A1", "ID")
	file.SetCellValue("Sheet1", "B1", "Name")
	file.SetCellValue("Sheet1", "C1", "Price")
	file.SetCellValue("Sheet1", "D1", "Quantity")

	// Fill in the data rows
	row := 2
	for rows.Next() {
		// Scan data into appropriate variables based on your data structure
		var id int
		var name string
		var price int
		var quantity int
		err := rows.Scan(&id, &name, &price, &quantity)
		if err != nil {
			return "", err
		}

		file.SetCellValue("Sheet1", fmt.Sprintf("A%d", row), id)
		file.SetCellValue("Sheet1", fmt.Sprintf("B%d", row), name)
		file.SetCellValue("Sheet1", fmt.Sprintf("C%d", row), price)
		file.SetCellValue("Sheet1", fmt.Sprintf("D%d", row), quantity)
		row++
	}

	// Set active sheet of the workbook
	file.SetActiveSheet(index)

	// Save the Excel file
	excelFileName := "products.xlsx"
	if err := file.SaveAs(excelFileName); err != nil {
		return "", err
	}

	return excelFileName, nil
}
