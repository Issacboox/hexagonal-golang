package repository

import (
	m "bam/internal/app/model"

	"gorm.io/gorm"
)

type IProductRepository interface {
	CreateProduct(prod *m.Product) error
	FindProductByID(id uint) (*m.Product, error)
	UpdateProduct(prod *m.Product) error
	DeleteProduct(id uint) error
	FindProductByName(name string) ([]*m.Product, error)
	FindProducts() ([]*m.Product, error)
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
