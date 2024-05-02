package service

import (
	// "errors"
	m "bam/internal/app/model"
	r "bam/internal/app/repository"
)

type ProductService struct {
	repo r.IProductRepository
}

func NewProductService(repo r.IProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetProductByID(id uint) (*m.Product, error) {
	return s.repo.FindProductByID(id)
}

func (s *ProductService) CreateProduct(prod *m.Product) error {
	return s.repo.CreateProduct(prod)
}

func (s *ProductService) UpdateProduct(prod *m.Product) error {
	return s.repo.UpdateProduct(prod)
}

func (s *ProductService) DeleteProduct(id uint) error {
	return s.repo.DeleteProduct(id)
}

func (s *ProductService) GetProducts() ([]*m.Product, error) {
	return s.repo.FindProducts()
}

func (s *ProductService) FindProductByName(name string) ([]*m.Product, error) {
	return s.repo.FindProductByName(name)
}

func (s *ProductService) InsertProductsFromExcel(file string) error {
    return s.repo.InsertProductsFromExcel(file)
}
