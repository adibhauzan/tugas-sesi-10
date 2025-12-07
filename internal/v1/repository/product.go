package repository

import (
	"context"
	"strings"
	"tugas-sesi-10-arsitektur-berbasis-layanan/internal/v1/domain/entity"
	"tugas-sesi-10-arsitektur-berbasis-layanan/pkg/common"

	"gorm.io/gorm"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, tx *gorm.DB, product *entity.Product) (*entity.Product, error)
	GetAll(ctx context.Context, filter *common.PaginationFilter) ([]entity.Product, int64, error)
	GetByID(ctx context.Context, id string) (*entity.Product, error)
	Update(ctx context.Context, tx *gorm.DB, data *entity.Product) error
	Delete(ctx context.Context, tx *gorm.DB, id string) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{db: db}
}

func (r *productRepository) CreateProduct(ctx context.Context, tx *gorm.DB, product *entity.Product) (*entity.Product, error) {
	if err := tx.WithContext(ctx).Create(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (r *productRepository) GetAll(ctx context.Context, filter *common.PaginationFilter) ([]entity.Product, int64, error) {
	var (
		data      []entity.Product
		totalRows int64
	)

	query := r.db.WithContext(ctx).Model(&entity.Product{}).
		Joins("LEFT JOIN product_types pt ON pt.id = products.type_id")

	if filter.Search != "" {
		q := "%" + strings.ToLower(filter.Search) + "%"
		query = query.Where(`
			LOWER(products.name) LIKE ?
			OR LOWER(pt.name) LIKE ?
		`, q, q)
	}

	if err := query.Count(&totalRows).Error; err != nil {
		return nil, 0, err
	}

	sortBy := filter.SortBy
	sortDir := filter.SortDir
	if sortDir == "" {
		sortDir = "desc"
	}

	switch sortBy {
	case "type_name":
		query = query.Order("pt.name " + sortDir)
	case "name":
		query = query.Order("products.name " + sortDir)
	case "price":
		query = query.Order("products.price " + sortDir)
	default:
		query = query.Order("products.created_at " + sortDir)
	}

	offset := (filter.Page - 1) * filter.PageSize

	err := query.
		Preload("Type").
		Limit(filter.PageSize).
		Offset(offset).
		Find(&data).Error

	if err != nil {
		return nil, 0, err
	}

	return data, totalRows, nil
}

func (r *productRepository) GetByID(ctx context.Context, id string) (*entity.Product, error) {
	var product entity.Product
	err := r.db.WithContext(ctx).
		Preload("Type").
		First(&product, "id = ?", id).Error

	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) Update(ctx context.Context, tx *gorm.DB, data *entity.Product) error {
	if err := tx.WithContext(ctx).Model(&entity.Product{}).Where("id = ?", data.ID).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

func (r *productRepository) Delete(ctx context.Context, tx *gorm.DB, id string) error {
	if err := tx.WithContext(ctx).Delete(&entity.Product{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
