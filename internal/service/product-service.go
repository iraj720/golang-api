package service

import (
	"context"
	"snapp/internal/entity"
	"snapp/internal/repository"
)

type productService struct {
	repository repository.IProductRepository
}

type IProductService interface {
	GetAllProducts(ctx context.Context) []entity.Product
	InsertProduct(ctx context.Context, p *entity.Product)
	VoteToProduct(ctx context.Context, id uint, name string, vote float32)
	GetSpecialProducts(ctx context.Context, id uint, name string, point float32) []entity.Product
}

func NewProductService(r repository.IProductRepository) IProductService {
	return &productService{repository: r}
}

func (ps *productService) GetAllProducts(ctx context.Context) []entity.Product {
	return ps.repository.FilterProducts(ctx, &entity.Product{})
}

func (ps *productService) GetSpecialProducts(ctx context.Context, id uint, name string, point float32) []entity.Product {
	p := &entity.Product{ID: id, Name: name, Point: point}
	return ps.repository.FilterProducts(ctx, p)
}

func (ps *productService) InsertProduct(ctx context.Context, p *entity.Product) {
	ps.repository.InsertProduct(ctx, p)
}

func (ps *productService) VoteToProduct(ctx context.Context, id uint, name string, vote float32) {
	p := &entity.Product{
		ID:   id,
		Name: name,
	}
	pr := ps.repository.GetProduct(ctx, p)
	newpr := voteToProduct(pr, vote)
	ps.repository.UpdateProduct(ctx, &newpr)
}

func voteToProduct(pr entity.Product, vote float32) entity.Product {
	pr.Point = ((float32(pr.VotesCount) * pr.Point) + vote) / float32(pr.VotesCount+1)
	pr.VotesCount++
	return pr
}
