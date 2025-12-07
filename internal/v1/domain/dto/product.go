package dto

import (
	"time"
	"tugas-sesi-10-arsitektur-berbasis-layanan/pkg/common"
)

// Request Body
type ProductCreateRequest struct {
	Name   string  `json:"name" binding:"required"`
	Code   int     `json:"code" binding:"required"`
	TypeID string  `json:"type_id" binding:"required"`
	Price  float64 `json:"price" binding:"required"`
	Stock  int     `json:"stock" binding:"required"`
}

type ProductUpdateRequest struct {
	Name   string  `json:"name,omitempty"`
	Code   int     `json:"code,omitempty"`
	TypeID string  `json:"type_id,omitempty"`
	Price  float64 `json:"price,omitempty"`
	Stock  int     `json:"stock,omitempty"`
}

// Response Body
type ProductResponse struct {
	ID        string              `json:"id"`
	Name      string              `json:"name"`
	Code      int                 `json:"code"`
	Type      ProductTypeResponse `json:"type"`
	Price     float64             `json:"price"`
	Stock     int                 `json:"stock"`
	CreatedAt time.Time           `json:"created_at,omitempty"`
	UpdatedAt time.Time           `json:"updated_at,omitempty"`
}

type ProductListResponse struct {
	Data       []ProductResponse     `json:"product"`
	Pagination common.PaginationMeta `json:"pagination"`
}
