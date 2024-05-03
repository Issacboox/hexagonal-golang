package handler

import (
	"fmt"
	"mime/multipart"

	"github.com/gofiber/fiber/v2"
)

type ExcelActions interface {
	ReadExcel(file *multipart.FileHeader) ([]string, error)
	ExportDataToExcel() (string, error)
}

type ExcelHandler struct {
	excelService ExcelActions
}

func NewExcelHandler(service ExcelActions) *ExcelHandler {
	return &ExcelHandler{excelService: service}
}

func (h *ExcelHandler) ReadExcel(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	rows, err := h.excelService.ReadExcel(file)
	if err != nil {
		return err
	}

	// Send the data back to the user
	return c.JSON(rows)
}

func (h *ExcelHandler) ExportDataToExcel(c *fiber.Ctx) error {
	// Call the service to export data to Excel
	excelFileName, err := h.excelService.ExportDataToExcel()
	if err != nil {
		// Log the error
		fmt.Println("Error exporting data to Excel:", err)
		return err
	}

	// Return the Excel file to the client for download
	return c.SendFile(excelFileName)
}
