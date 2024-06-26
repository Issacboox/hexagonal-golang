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
	InsertProductsFromExcel(file *multipart.FileHeader) ([]*m.Product, error)
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
	result := r.db.Where("name LIKE ?", "%"+name+"%").Find(&prods)
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

func (r *ProductRepository) InsertProductsFromExcel(file *multipart.FileHeader) ([]*m.Product, error) {
	// Start a transaction
	tx := r.db.Begin()

	var insertedProducts []*m.Product
	var duplicatedProducts []*m.Product

	f, err := excelize.OpenFile(file.Filename)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	rows := f.GetRows("Sheet1")
	for _, row := range rows {
		name := row[0]
		priceStr := row[1]
		quantityStr := row[2]

		// Convert price and quantity strings to int
		price, err := strconv.Atoi(priceStr)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		quantity, err := strconv.Atoi(quantityStr)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		// Check if the product already exists
		var existingProd m.Product
		result := tx.Where("name = ?", name).First(&existingProd)
		if result.Error == nil {
			// Product already exists, add it to duplicatedProducts
			duplicatedProducts = append(duplicatedProducts, &existingProd)
			continue
		}

		// Product does not exist, create a new one
		prod := &m.Product{Name: name, Price: price, Quantity: quantity}
		if err := tx.Create(prod).Error; err != nil {
			tx.Rollback()
			return nil, err
		}

		insertedProducts = append(insertedProducts, prod)
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return duplicatedProducts, nil
}
