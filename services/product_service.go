package services

import (
	"github.com/KicaiBadai/gin-firebase-backend/models"
	"github.com/KicaiBadai/gin-firebase-backend/repositories"
)

type ProductService struct {
	productRepo *repositories.ProductRepository
}

func NewProductService() *ProductService {
	return &ProductService{
		productRepo: repositories.NewProductRepository(),
	}
}

// GetAll ambil semua produk dengan pagination
func (s *ProductService) GetAll(page, limit int, category string) ([]models.Product, int64, error) {
	if page <= 0 {
		page = 1
	}

	if limit <= 0 || limit > 100 {
		limit = 10
	}

	return s.productRepo.FindAll(page, limit, category)
}

// GetByID ambil produk berdasarkan ID
func (s *ProductService) GetByID(id uint) (*models.Product, error) {
	return s.productRepo.FindByID(id)
}

// Create membuat produk baru
func (s *ProductService) Create(req *models.CreateProductRequest) (*models.Product, error) {
	product := &models.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		Category:    req.Category,
		ImageURL:    req.ImageURL,
	}

	err := s.productRepo.Create(product)
	return product, err
}

// Update memperbarui produk
func (s *ProductService) Update(id uint, req *models.UpdateProductRequest) (*models.Product, error) {
	product, err := s.productRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	// Update hanya field yang dikirim (pointer nil = tidak diupdate)
	if req.Name != nil {
		product.Name = *req.Name
	}
	if req.Description != nil {
		product.Description = *req.Description
	}
	if req.Price != nil {
		product.Price = *req.Price
	}
	if req.Stock != nil {
		product.Stock = *req.Stock
	}
	if req.Category != nil {
		product.Category = *req.Category
	}
	if req.ImageURL != nil {
		product.ImageURL = *req.ImageURL
	}

	err = s.productRepo.Update(product)
	return product, err
}

// Delete menghapus produk (soft delete)
func (s *ProductService) Delete(id uint) error {
	return s.productRepo.Delete(id)
}