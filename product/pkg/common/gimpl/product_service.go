package gimpl

import (
	"__product/ent"
	productpb "__product/gen/pb/v1"
	"context"
)

type productService struct {
	productpb.UnimplementedProductServiceServer

	ent *ent.Client
}

func NewProductService(e *ent.Client) *productService {
	return &productService{
		ent: e,
	}
}

func (s *productService) CreateProduct(ctx context.Context, req *productpb.CreateProductRequest) (*productpb.CreateProductResponse, error) {
	panic("implement me")
}

func (s *productService) GetProduct(ctx context.Context, req *productpb.GetProductRequest) (*productpb.GetProductResponse, error) {
	panic("implement me")
}

func (s *productService) ListProducts(ctx context.Context, req *productpb.ListProductRequest) (*productpb.ListProductResponse, error) {
	panic("implement me")
}

func (s *productService) UpdateProduct(ctx context.Context, req *productpb.UpdateProductRequest) (*productpb.UpdateProductResponse, error) {
	panic("implement me")
}

func (s *productService) DeleteProduct(ctx context.Context, req *productpb.DeleteProductRequest) (*productpb.DeleteProductResponse, error) {
	panic("implement me")
}
