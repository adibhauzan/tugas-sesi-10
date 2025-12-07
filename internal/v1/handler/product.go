package handler

import (
	"context"
	"net/http"
	"strconv"
	"tugas-sesi-10-arsitektur-berbasis-layanan/internal/v1/domain/dto"
	"tugas-sesi-10-arsitektur-berbasis-layanan/internal/v1/service"
	"tugas-sesi-10-arsitektur-berbasis-layanan/pkg/common"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service service.ProductService
}

func NewProductHandler(s service.ProductService) *ProductHandler {
	return &ProductHandler{service: s}
}

func (h *ProductHandler) GetAll(ctx *gin.Context) {
	ctxWithCancel, cancel := context.WithCancel(ctx.Request.Context())
	defer cancel()

	filter := &common.PaginationFilter{}
	filter.Page, _ = strconv.Atoi(ctx.DefaultQuery("page", "1"))
	filter.PageSize, _ = strconv.Atoi(ctx.DefaultQuery("page_size", "10"))
	filter.Search = ctx.Query("search")
	filter.SortBy = ctx.DefaultQuery("sort_by", "created_at")
	filter.SortDir = ctx.DefaultQuery("sort_dir", "desc")

	res, err := h.service.GetAll(ctxWithCancel, filter)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.WebResponse{
			Code:       http.StatusInternalServerError,
			Status:     http.StatusText(http.StatusInternalServerError),
			Message:    http.StatusText(http.StatusInternalServerError),
			MessageDev: err.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, common.WebResponse{
		Code:       http.StatusOK,
		Status:     http.StatusText(http.StatusOK),
		Message:    "Success",
		MessageDev: "Success",
		Data:       res,
	})
}

func (h *ProductHandler) Create(ctx *gin.Context) {
	ctxWithCancel, cancel := context.WithCancel(ctx.Request.Context())
	defer cancel()

	var req dto.ProductCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, common.WebResponse{
			Code:       http.StatusBadRequest,
			Status:     http.StatusText(http.StatusBadRequest),
			Message:    http.StatusText(http.StatusBadRequest),
			MessageDev: err.Error(),
		})
		ctx.Abort()
		return
	}

	res, err := h.service.Create(ctxWithCancel, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.WebResponse{
			Code:       http.StatusInternalServerError,
			Status:     http.StatusText(http.StatusInternalServerError),
			Message:    http.StatusText(http.StatusInternalServerError),
			MessageDev: err.Error(),
		})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, common.WebResponse{
		Code:       http.StatusCreated,
		Status:     http.StatusText(http.StatusCreated),
		Message:    "Create Product Successfully",
		MessageDev: "Create Product Successfully",
		Data:       res,
	})
}

func (h *ProductHandler) GetByID(ctx *gin.Context) {
	ctxWithCancel, cancel := context.WithCancel(ctx.Request.Context())
	defer cancel()

	id := ctx.Param("id")

	res, err := h.service.GetByID(ctxWithCancel, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, common.WebResponse{
			Code:       http.StatusNotFound,
			Status:     http.StatusText(http.StatusNotFound),
			Message:    "Product not found",
			MessageDev: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, common.WebResponse{
		Code:       http.StatusOK,
		Status:     http.StatusText(http.StatusOK),
		Message:    "Success",
		MessageDev: "Success",
		Data:       res,
	})
}

func (h *ProductHandler) Update(ctx *gin.Context) {
	ctxWithCancel, cancel := context.WithCancel(ctx.Request.Context())
	defer cancel()

	id := ctx.Param("id")

	var req dto.ProductUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, common.WebResponse{
			Code:       http.StatusBadRequest,
			Status:     http.StatusText(http.StatusBadRequest),
			Message:    http.StatusText(http.StatusBadRequest),
			MessageDev: err.Error(),
		})
		return
	}

	err := h.service.Update(ctxWithCancel, id, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.WebResponse{
			Code:       http.StatusInternalServerError,
			Status:     http.StatusText(http.StatusInternalServerError),
			Message:    http.StatusText(http.StatusInternalServerError),
			MessageDev: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, common.WebResponse{
		Code:       http.StatusOK,
		Status:     http.StatusText(http.StatusOK),
		Message:    "Updated Product Successfully",
		MessageDev: "Updated Product Successfully",
	})
}

func (h *ProductHandler) Delete(ctx *gin.Context) {
	ctxWithCancel, cancel := context.WithCancel(ctx.Request.Context())
	defer cancel()

	id := ctx.Param("id")

	err := h.service.Delete(ctxWithCancel, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, common.WebResponse{
			Code:       http.StatusInternalServerError,
			Status:     http.StatusText(http.StatusInternalServerError),
			Message:    http.StatusText(http.StatusInternalServerError),
			MessageDev: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, common.WebResponse{
		Code:       http.StatusOK,
		Status:     http.StatusText(http.StatusOK),
		Message:    "Deleted successfully",
		MessageDev: "Deleted successfully",
	})
}
