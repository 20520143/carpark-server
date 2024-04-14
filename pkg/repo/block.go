package repo

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"gitlab.com/goxp/cloud0/ginext"
	"gitlab.com/goxp/cloud0/logger"
	"gorm.io/gorm"
	"net/http"
	"parking-server/pkg/model"
	"parking-server/pkg/utils"
	"parking-server/pkg/valid"
)

func (r *RepoPG) CreateBlock(ctx context.Context, req *model.Block) error {
	log := logger.WithCtx(ctx, utils.GetCurrentCaller(r, 0))

	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()

	if err := tx.Model(&model.Block{}).Create(&req).Error; err != nil {
		log.WithError(err).Error("error_500: error when CreateBlock")
		return ginext.NewError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func (r *RepoPG) GetOneBlock(ctx context.Context, id uuid.UUID) (res model.Block, err error) {
	log := logger.WithCtx(ctx, utils.GetCurrentCaller(r, 0))

	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()

	if err = tx.Model(&model.Block{}).Where("id = ?", id).Take(&res).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.WithError(err).Error("error_404: not found")
			return res, ginext.NewError(http.StatusNotFound, err.Error())
		}
		log.WithError(err).Error("error_500: failed to GetOneBlock")
		return res, ginext.NewError(http.StatusInternalServerError, err.Error())
	}
	return res, nil
}

func (r *RepoPG) GetListBlock(ctx context.Context, req model.ListBlockReq) (res model.ListBlockRes, err error) {
	log := logger.WithCtx(ctx, utils.GetCurrentCaller(r, 0))

	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()

	tx = tx.Model(&model.Block{})

	if req.Code != nil {
		name := utils.TransformString(valid.String(req.Code), false)
		tx = tx.Where("unaccent(name) ilike ?", name+"%")
	}

	if req.Sort != "" {
		tx = tx.Order(req.Sort)
	} else {
		tx = tx.Order("created_at desc")
	}

	var total int64 = 0
	page := r.GetPage(req.Page)
	pageSize := r.GetPageSize(req.PageSize)

	if err := tx.Count(&total).Limit(pageSize).Offset(r.GetOffset(page, pageSize)).Find(&res.Data).Error; err != nil {
		log.WithError(err).Error("error_500: failed to GetListBlock")
		return res, ginext.NewError(http.StatusInternalServerError, err.Error())
	}

	if res.Meta, err = r.GetPaginationInfo("", nil, int(total), page, pageSize); err != nil {
		log.WithError(err).Error("error_500: failed to get pagination")
		return res, ginext.NewError(http.StatusInternalServerError, err.Error())
	}

	return res, nil
}

func (r *RepoPG) UpdateBlock(ctx context.Context, req *model.Block) error {
	log := logger.WithCtx(ctx, utils.GetCurrentCaller(r, 0))

	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()

	if err := tx.Model(&model.Block{}).Where("id = ?", req.ID).Save(&req).Error; err != nil {
		log.WithError(err).Error("error_500: error when UpdateBlock")
		return ginext.NewError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func (r *RepoPG) DeleteBlock(ctx context.Context, id uuid.UUID) error {
	log := logger.WithCtx(ctx, utils.GetCurrentCaller(r, 0))

	tx, cancel := r.DBWithTimeout(ctx)
	defer cancel()

	if err := tx.Where("id = ?", id).Delete(&model.Block{}).Error; err != nil {
		log.WithError(err).Error("error_500: error when DeleteBlock")
		return ginext.NewError(http.StatusInternalServerError, err.Error())
	}
	return nil
}
