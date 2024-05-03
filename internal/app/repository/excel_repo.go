package repository

import (
	"mime/multipart"

	"github.com/360EntSecGroup-Skylar/excelize"
	"gorm.io/gorm"
)

type IExcelRepository interface {
	ReadExcel(file *multipart.FileHeader) ([][]string, error)
    GetExcelData(data interface{}) *gorm.DB
}

type ExcelRepository struct {
	db           *gorm.DB
	ExcelActions IExcelRepository
}

func NewExcelRepository(db *gorm.DB) *ExcelRepository {
	return &ExcelRepository{db: db}
}

func (er *ExcelRepository) ReadExcel(file *multipart.FileHeader) ([][]string, error) {
    f, err := file.Open()
    if err != nil {
        return nil, err
    }
    defer f.Close()

    xlsx, err := excelize.OpenReader(f)
    if err != nil {
        return nil, err
    }

    // อ่านข้อมูล Excel
    rows := xlsx.GetRows("Sheet1")

    return rows, nil
}


func (er *ExcelRepository) GetExcelData(data interface{}) *gorm.DB {
	return er.db.Find(data)
}