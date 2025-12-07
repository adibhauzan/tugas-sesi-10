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

type ProductTypeService interface {
	GetAll(ctx context.Context, filter *common.PaginationFilter) (dto.ProductTypeListResponse, error)
	Create(ctx context.Context, req *dto.ProductTypeCreateRequest) (dto.ProductTypeResponse, error)
	GetByID(ctx context.Context, id string) (dto.ProductTypeResponse, error)
	Update(ctx context.Context, id string, req *dto.ProductTypeUpdateRequest) error
	Delete(ctx context.Context, id string) error
}

type productTypeService struct {
	repo   repository.ProductTypeRepository
	db     *gorm.DB
	logger *logrus.Logger
}

func NewProductTypeService(repo repository.ProductTypeRepository, db *gorm.DB, logger *logrus.Logger) ProductTypeService {
	return &productTypeService{repo: repo, db: db, logger: logger}
}

func (s *productTypeService) Create(ctx context.Context, req *dto.ProductTypeCreateRequest) (dto.ProductTypeResponse, error) {
	data := &entity.ProductType{
		ID:   uuid.New().String(),
		Name: req.Name,
		// CreatedAt: time.Now(),
		// UpdatedAt: time.Now(),
	}

	var result *entity.ProductType
	err := s.db.Transaction(func(tx *gorm.DB) error {
		datas, err := s.repo.CreateProductType(ctx, tx, data)
		if err != nil {
			return err
		}
		result = datas
		return nil
	})
	if err != nil {
		return dto.ProductTypeResponse{}, err
	}

	return dto.ProductTypeResponse{
		ID:        result.ID,
		Name:      result.Name,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}, nil
}

func (s *productTypeService) GetAll(ctx context.Context, filter *common.PaginationFilter) (dto.ProductTypeListResponse, error) {
	data, total, err := s.repo.GetAll(ctx, filter)
	if err != nil {
		s.logger.WithContext(ctx).WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("failed to get product types")
		return dto.ProductTypeListResponse{}, err
	}

	res := make([]dto.ProductTypeResponse, len(data))
	for i, e := range data {
		res[i] = dto.ProductTypeResponse{ID: e.ID, Name: e.Name, CreatedAt: e.CreatedAt, UpdatedAt: e.UpdatedAt}
	}

	return dto.ProductTypeListResponse{
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

func (s *productTypeService) GetByID(ctx context.Context, id string) (dto.ProductTypeResponse, error) {
	data, err := s.repo.GetByID(ctx, id)
	if err != nil {
		s.logger.WithContext(ctx).WithField("id", id).Error("failed to get product type by id")
		return dto.ProductTypeResponse{}, err
	}

	return dto.ProductTypeResponse{
		ID:        data.ID,
		Name:      data.Name,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}, nil
}

func (s *productTypeService) Update(ctx context.Context, id string, req *dto.ProductTypeUpdateRequest) error {
	err := s.db.Transaction(func(tx *gorm.DB) error {
		existing, err := s.repo.GetByID(ctx, id)
		if err != nil {
			return err
		}

		existing.Name = req.Name
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
		}).Error("failed to update product type")
		return err
	}

	return nil
}

func (s *productTypeService) Delete(ctx context.Context, id string) error {
	err := s.db.Transaction(func(tx *gorm.DB) error {
		_, err := s.repo.GetByID(ctx, id)
		if err != nil {
			return err
		}

		if err := s.repo.Delete(ctx, tx, id); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		s.logger.WithContext(ctx).WithFields(logrus.Fields{
			"id":    id,
			"error": err.Error(),
		}).Error("failed to delete product type")

		return err
	}

	return nil
}
