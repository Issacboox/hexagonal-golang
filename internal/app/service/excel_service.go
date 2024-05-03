package service

import (
	// "log"

	"fmt"
	"strings"

	// "fmt"
	"mime/multipart"

	m "bam/internal/app/model"
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
	var products []*m.Product
	if err := s.repo.GetExcelData(&products).Error; err != nil {
		return "", err
	}

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
	for i, product := range products {
		row := i + 2 // Excel rows start from 1, but slice index starts from 0
		file.SetCellValue("Sheet1", fmt.Sprintf("A%d", row), product.ID)
		file.SetCellValue("Sheet1", fmt.Sprintf("B%d", row), product.Name)
		file.SetCellValue("Sheet1", fmt.Sprintf("C%d", row), product.Price)
		file.SetCellValue("Sheet1", fmt.Sprintf("D%d", row), product.Quantity)
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