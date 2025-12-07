package service

import (
	"context"
	"time"
	"tugas-sesi-10-arsitektur-berbasis-layanan/internal/v1/domain/dto"
	"tugas-sesi-10-arsitektur-berbasis-layanan/internal/v1/domain/entity"
	"tugas-sesi-10-arsitektur-berbasis-layanan/internal/v1/repository"
	"tugas-sesi-10-arsitektur-berbasis-layanan/pkg/common"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProductService interface {
	GetAll(ctx context.Context, filter *common.PaginationFilter) (dto.ProductListResponse, error)
	Create(ctx context.Context, req *dto.ProductCreateRequest) (dto.ProductResponse, error)
	GetByID(ctx context.Context, id string) (dto.ProductResponse, error)
	Update(ctx context.Context, id string, req *dto.ProductUpdateRequest) error
	Delete(ctx context.Context, id string) error
}

type productService struct {
	repo            repository.ProductRepository
	productTypeRepo repository.ProductTypeRepository
	db              *gorm.DB
	logger          *logrus.Logger
}

func NewProductService(repo repository.ProductRepository, productTypeRepo repository.ProductTypeRepository, db *gorm.DB, logger *logrus.Logger) ProductService {
	return &productService{repo: repo, productTypeRepo: productTypeRepo, db: db, logger: logger}
}
func (s *productService) Create(ctx context.Context, req *dto.ProductCreateRequest) (dto.ProductResponse, error) {
	// cek product type
	productType, err := s.productTypeRepo.GetByID(ctx, req.TypeID)
	if err != nil {
		s.logger.WithContext(ctx).WithField("type_id", req.TypeID).Error("failed to get product type by id")
		return dto.ProductResponse{}, err
	}

	data := &entity.Product{
		ID:        uuid.New().String(),
		TypeID:    req.TypeID,
		Name:      req.Name,
		Code:      req.Code,
		Price:     req.Price,
		Stock:     req.Stock,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	var result *entity.Product
	err = s.db.Transaction(func(tx *gorm.DB) error {
		datas, err := s.repo.CreateProduct(ctx, tx, data)
		if err != nil {
			return err
		}
		result = datas
		return nil
	})
	if err != nil {
		s.logger.WithContext(ctx).WithField("error", err.Error()).Error("failed to create product")
		return dto.ProductResponse{}, err
	}

	return dto.ProductResponse{
		ID:   result.ID,
		Name: result.Name,
		Code: result.Code,
		Type: dto.ProductTypeResponse{
			ID:        productType.ID,
			Name:      productType.Name,
			CreatedAt: productType.CreatedAt,
			UpdatedAt: productType.UpdatedAt,
		},
		Price:     result.Price,
		Stock:     result.Stock,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}, nil
}

func (s *productService) GetAll(ctx context.Context, filter *common.PaginationFilter) (dto.ProductListResponse, error) {
	data, total, err := s.repo.GetAll(ctx, filter)
	if err != nil {
		s.logger.WithContext(ctx).WithField("error", err.Error()).Error("failed to get products")
		return dto.ProductListResponse{}, err
	}

	res := make([]dto.ProductResponse, len(data))
	for i, e := range data {
		typeResp := dto.ProductTypeResponse{}
		if e.Type.ID != "" {
			typeResp = dto.ProductTypeResponse{
				ID:        e.Type.ID,
				Name:      e.Type.Name,
				CreatedAt: e.CreatedAt,
				UpdatedAt: e.UpdatedAt,
			}
		}
		res[i] = dto.ProductResponse{
			ID:        e.ID,
			Name:      e.Name,
			Code:      e.Code,
			Price:     e.Price,
			Stock:     e.Stock,
			Type:      typeResp,
			CreatedAt: e.CreatedAt,
			UpdatedAt: e.UpdatedAt,
		}
	}

	return dto.ProductListResponse{
		Data: res,
		Pagination: common.PaginationMeta{
			Total:         total,
			Page:          filter.Page,
			PageSize:      filter.PageSize,
			SortBy:        filter.SortBy,
			SortDirection: filter.SortDir,
			Search:        filter.Search,
		},
	}, nil
}

func (s *productService) GetByID(ctx context.Context, id string) (dto.ProductResponse, error) {
	data, err := s.repo.GetByID(ctx, id)
	if err != nil {
		s.logger.WithContext(ctx).WithField("id", id).Error("failed to get product by id")
		return dto.ProductResponse{}, err
	}

	typeResp := dto.ProductTypeResponse{}
	if data.Type.ID != "" {
		typeResp = dto.ProductTypeResponse{
			ID:        data.Type.ID,
			Name:      data.Type.Name,
			CreatedAt: data.Type.CreatedAt,
			UpdatedAt: data.Type.UpdatedAt,
		}
	}

	return dto.ProductResponse{
		ID:        data.ID,
		Name:      data.Name,
		Code:      data.Code,
		Price:     data.Price,
		Stock:     data.Stock,
		Type:      typeResp,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}, nil
}

func (s *productService) Update(ctx context.Context, id string, req *dto.ProductUpdateRequest) error {
	err := s.db.Transaction(func(tx *gorm.DB) error {
		existing, err := s.repo.GetByID(ctx, id)
		if err != nil {
			return err
		}

		existing.Name = req.Name
		existing.Code = req.Code
		existing.Price = req.Price
		existing.Stock = req.Stock
		existing.TypeID = req.TypeID
		existing.UpdatedAt = time.Now()

		err = s.repo.Update(ctx, tx, existing)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		s.logger.WithContext(ctx).WithFields(logrus.Fields{
			"id":    id,
			"error": err.Error(),
		}).Error("failed to update product")
		return err
	}

	return nil
}

func (s *productService) Delete(ctx context.Context, id string) error {
	err := s.db.Transaction(func(tx *gorm.DB) error {
		_, err := s.repo.GetByID(ctx, id)
		if err != nil {
			return err
		}

		return s.repo.Delete(ctx, tx, id)
	})

	if err != nil {
		s.logger.WithContext(ctx).WithFields(logrus.Fields{
			"id":    id,
			"error": err.Error(),
		}).Error("failed to delete product")
		return err
	}

	return nil
}
