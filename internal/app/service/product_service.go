package service

import (
	"mime/multipart"
	"strconv"

	m "bam/internal/app/model"
	r "bam/internal/app/repository"

	"github.com/360EntSecGroup-Skylar/excelize"
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

func (s *ProductService) InsertProductsFromExcel(file *multipart.FileHeader) error {
	// Open a reader from the file's stream
	reader, err := file.Open()
	if err != nil {
		return err
	}
	defer reader.Close()

	// Process the file directly without opening it from the filesystem
	f, err := excelize.OpenReader(reader)
	if err != nil {
		return err
	}

	// Get all rows from the Sheet1
	rows := f.GetRows("Sheet1")

	// Process each row starting from the second row
	for i, row := range rows {
		// Skip the first row
		if i == 0 {
			continue
		}

		name := row[0]
		price, err := strconv.Atoi(row[1])
		if err != nil {
			return err
		}
		quantity, err := strconv.Atoi(row[2])
		if err != nil {
			return err
		}

		// Create a new product and insert it
		prod := &m.Product{Name: name, Price: price, Quantity: quantity}
		if err := s.repo.CreateProduct(prod); err != nil {
			return err
		}
	}

	return nil
}

// Implement the ReadExcelFile method in your ProductService struct
func (s *ProductService) ReadExcelFile(file *multipart.FileHeader) ([][]string, error) {
	f, err := excelize.OpenFile(file.Filename)
	if err != nil {
		return nil, err
	}

	rows := f.GetRows("Sheet1")
	return rows, nil
}
