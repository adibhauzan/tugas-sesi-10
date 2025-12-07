package repository

import (
	"context"
	"tugas-sesi-10-arsitektur-berbasis-layanan/internal/v1/domain/entity"
	"tugas-sesi-10-arsitektur-berbasis-layanan/pkg/common"

	"gorm.io/gorm"
)

type ProductTypeRepository interface {
	CreateProductType(ctx context.Context, tx *gorm.DB, productType *entity.ProductType) (*entity.ProductType, error)
	GetAll(ctx context.Context, filter *common.PaginationFilter) ([]entity.ProductType, int64, error)
	GetByID(ctx context.Context, id string) (*entity.ProductType, error)
	Update(ctx context.Context, tx *gorm.DB, data *entity.ProductType) error
	Delete(ctx context.Context, tx *gorm.DB, id string) error
}

type productTypeRepository struct {
	db *gorm.DB
}

func NewProductTypeRepository(db *gorm.DB) *productTypeRepository {
	return &productTypeRepository{db: db}
}

func (r *productTypeRepository) CreateProductType(
	ctx context.Context,
	tx *gorm.DB,
	productType *entity.ProductType,
) (*entity.ProductType, error) {

	if err := tx.WithContext(ctx).Create(productType).Error; err != nil {
		return nil, err
	}

	return productType, nil
}

func (r *productTypeRepository) GetAll(ctx context.Context, filter *common.PaginationFilter) ([]entity.ProductType, int64, error) {
	filter.Normalize()

	var results []entity.ProductType
	var total int64

	q := r.db.WithContext(ctx).Model(&entity.ProductType{})

	if filter.Search != "" {
		q = q.Where("name ILIKE ?", "%"+filter.Search+"%")
	}

	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	order := "created_at desc"
	if filter.SortBy != "" {
		order = filter.SortBy + " " + filter.SortDir
	}

	offset := (filter.Page - 1) * filter.PageSize

	if err := q.Order(order).
		Limit(filter.PageSize).
		Offset(offset).
		Find(&results).Error; err != nil {
		return nil, 0, err
	}

	return results, total, nil
}

func (r *productTypeRepository) GetByID(
	ctx context.Context,
	id string,
) (*entity.ProductType, error) {

	var data entity.ProductType

	if err := r.db.WithContext(ctx).
		Where("id = ?", id).
		First(&data).Error; err != nil {

		return nil, err
	}

	return &data, nil
}

func (r *productTypeRepository) Update(
	ctx context.Context,
	tx *gorm.DB,
	data *entity.ProductType,
) error {

	if err := tx.WithContext(ctx).
		Model(&entity.ProductType{}).
		Where("id = ?", data.ID).
		Updates(map[string]interface{}{
			"name":       data.Name,
			"updated_at": data.UpdatedAt,
		}).Error; err != nil {

		return err
	}

	return nil
}

func (r *productTypeRepository) Delete(
	ctx context.Context,
	tx *gorm.DB,
	id string,
) error {

	if err := tx.WithContext(ctx).
		Where("id = ?", id).
		Delete(&entity.ProductType{}).Error; err != nil {

		return err
	}

	return nil
}
