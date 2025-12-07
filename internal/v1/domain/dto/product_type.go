package dto

import (
	"time"
	"tugas-sesi-10-arsitektur-berbasis-layanan/pkg/common"
)

// Request Body
type ProductTypeCreateRequest struct {
	Name string `json:"name" binding:"required"`
}

type ProductTypeUpdateRequest struct {
	Name string `json:"name" binding:"required"`
}

// Response Body
type ProductTypeResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ProductTypeListResponse struct {
	Data       []ProductTypeResponse `json:"product_type"`
	Pagination common.PaginationMeta `json:"pagination"`
}
