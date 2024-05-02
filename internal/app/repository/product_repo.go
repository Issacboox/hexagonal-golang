package repository

import (
	m "bam/internal/app/model"
	"mime/multipart"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	"gorm.io/gorm"
)

type IProductRepository interface {
	CreateProduct(prod *m.Product) error
	FindProductByID(id uint) (*m.Product, error)
	UpdateProduct(prod *m.Product) error
	DeleteProduct(id uint) error
	FindProductByName(name string) ([]*m.Product, error)
	FindProducts() ([]*m.Product, error)
	InsertProductsFromExcel(file *multipart.FileHeader) error
}

type ProductRepository struct {
	db             *gorm.DB
	ProductActions IProductRepository
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) CreateProduct(prod *m.Product) error {
	return r.db.Create(prod).Error
}

func (r *ProductRepository) FindProductByID(id uint) (*m.Product, error) {
	var prod m.Product
	result := r.db.First(&prod, id)
	return &prod, result.Error
}

func (r *ProductRepository) UpdateProduct(prod *m.Product) error {
	return r.db.Save(prod).Error
}

func (r *ProductRepository) DeleteProduct(id uint) error {
	return r.db.Delete(&m.Product{}, id).Error
}

func (r *ProductRepository) FindProductByName(name string) ([]*m.Product, error) {
	var prods []*m.Product
	result := r.db.Where("name = ?", name).Find(&prods)
	if result.Error != nil {
		return nil, result.Error
	}
	return prods, nil
}

func (r *ProductRepository) FindProducts() ([]*m.Product, error) {
	var prods []*m.Product
	result := r.db.Find(&prods)
	return prods, result.Error
}
func (r *ProductRepository) InsertProductsFromExcel(file *multipart.FileHeader) error {
	// Read Excel file and insert data into the database
	// Example code using excelize to read Excel file:
	f, err := excelize.OpenFile(file.Filename)
	if err != nil {
		return err
	}

	// Start a transaction
	tx := r.db.Begin()

	rows := f.GetRows("Sheet1")
	for _, row := range rows {
		name := row[0]
		priceStr := row[1]
		quantityStr := row[2]

		// Convert price and quantity strings to int
		price, err := strconv.Atoi(priceStr)
		if err != nil {
			tx.Rollback()
			return err
		}
		quantity, err := strconv.Atoi(quantityStr)
		if err != nil {
			tx.Rollback()
			return err
		}

		prod := &m.Product{Name: name, Price: price, Quantity: quantity}
		if err := tx.Create(prod).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// Commit the transaction
	return tx.Commit().Error
}
