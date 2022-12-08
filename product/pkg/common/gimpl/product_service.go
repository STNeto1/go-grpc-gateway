package gimpl

import (
	"__product/ent"
	"__product/ent/product"
	productpb "__product/gen/pb/v1"
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	_, err := s.ent.Product.Create().
		SetTitle(req.Title).
		SetDescription(req.Description).
		SetPrice(int(req.Price * 100)).
		SetUserID(uuid.MustParse(req.UserId)).
		SetUserName(req.UserName).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return &productpb.CreateProductResponse{
		Success: true,
	}, nil
}

func (s *productService) GetProduct(ctx context.Context, req *productpb.GetProductRequest) (*productpb.GetProductResponse, error) {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid id")
	}

	prd, err := s.ent.Product.Query().Where(product.ID(id)).First(ctx)
	if err != nil {
		return nil, status.Error(codes.NotFound, "prodct not found")
	}

	return &productpb.GetProductResponse{
		Product: &productpb.Product{
			Id:          prd.ID.String(),
			Title:       prd.Title,
			Description: prd.Description,
			Price:       float32(prd.Price / 100),
			UserId:      prd.UserID.String(),
			UserName:    prd.UserName,
		},
	}, nil
}

func (s *productService) ListProduct(ctx context.Context, req *productpb.ListProductRequest) (*productpb.ListProductResponse, error) {

	products, err := s.ent.Product.Query().Where(func(s *sql.Selector) {
		if req.GetTerm() != "" {
			s.Where(sql.Like(product.FieldTitle, "%"+req.GetTerm()+"%"))
		}
	}).All(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	res := productpb.ListProductResponse{}

	for _, prd := range products {
		res.Products = append(res.Products, &productpb.Product{
			Id:          prd.ID.String(),
			Title:       prd.Title,
			Description: prd.Description,
			Price:       float32(prd.Price / 100),
			UserId:      prd.UserID.String(),
			UserName:    prd.UserName,
		})
	}

	return &res, nil
}

func (s *productService) UpdateProduct(ctx context.Context, req *productpb.UpdateProductRequest) (*productpb.UpdateProductResponse, error) {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid id")
	}

	usrId, err := uuid.Parse(req.GetUserId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "product id")
	}

	prd, err := s.ent.Product.Update().
		Where(product.ID(id), product.UserID(usrId)).
		SetTitle(req.GetTitle()).
		SetDescription(req.GetDescription()).
		SetPrice(int(req.GetPrice() * 100)).
		Save(ctx)
	if err != nil {
		return nil, status.Error(codes.NotFound, "unable to update")
	}

	if prd == 0 {
		return &productpb.UpdateProductResponse{
			Success: false,
		}, nil
	}

	return &productpb.UpdateProductResponse{
		Success: true,
	}, nil
}

func (s *productService) DeleteProduct(ctx context.Context, req *productpb.DeleteProductRequest) (*productpb.DeleteProductResponse, error) {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid id")
	}

	usrId, err := uuid.Parse(req.GetUserId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "product id")
	}

	prd, err := s.ent.Product.
		Delete().
		Where(product.ID(id), product.UserID(usrId)).
		Exec(ctx)
	if err != nil {
		return nil, status.Error(codes.NotFound, "unable to update")
	}

	if prd == 0 {
		return &productpb.DeleteProductResponse{
			Success: false,
		}, nil
	}

	return &productpb.DeleteProductResponse{
		Success: true,
	}, nil
}
