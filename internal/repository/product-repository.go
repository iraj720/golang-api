package repository

import (
	"context"
	"snapp/internal/entity"

	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

type IProductRepository interface {
	GetProduct(ctx context.Context, p *entity.Product) entity.Product
	FilterProducts(ctx context.Context, p *entity.Product) []entity.Product
	GetProductById(ctx context.Context, id uint) entity.Product
	InsertProduct(ctx context.Context, p *entity.Product)
	UpdateProduct(ctx context.Context, p *entity.Product)
}

func NewProductRepository(db *gorm.DB) IProductRepository {
	return &productRepository{db: db}
}

func (pr *productRepository) GetProductById(ctx context.Context, id uint) entity.Product {
	return entity.Product{}
}

func (pr *productRepository) UpdateProduct(ctx context.Context, p *entity.Product) {
	pr.db.WithContext(ctx).Save(p)
}

func (pr *productRepository) InsertProduct(ctx context.Context, p *entity.Product) {
	pr.db.WithContext(ctx).Create(p)
}

func (pr *productRepository) FilterProducts(ctx context.Context, p *entity.Product) []entity.Product {
	var products []entity.Product
	pr.db.Where(p).Find(&products)
	return products
}

func (pr *productRepository) GetProduct(ctx context.Context, p *entity.Product) entity.Product {
	var product entity.Product
	pr.db.Where(p).First(&product)
	return product
}
