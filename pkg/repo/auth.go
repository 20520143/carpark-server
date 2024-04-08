package repo

import (
	"carpark-server/pkg/model"
	"carpark-server/pkg/utils"
	"context"
	"gitlab.com/goxp/cloud0/ginext"
	"gitlab.com/goxp/cloud0/logger"
	"gorm.io/gorm"
	"net/http"
)

func (r *RepoPG) GetOneUserByPhone(ctx context.Context, phoneNumber string, tx *gorm.DB) (*model.User, error) {
	var cancel context.CancelFunc
	log := logger.WithCtx(ctx, utils.GetCurrentCaller(r, 0))
	if tx == nil {
		tx, cancel = r.DBWithTimeout(ctx)
		defer cancel()
	}
	rs := &model.User{}

	if err := tx.Model(&model.User{}).Where("phone_number = ?", phoneNumber).Take(&rs).Error; err != nil {
		log.WithError(err).Error("Error when get one user by phone number")
		return nil, ginext.NewError(http.StatusInternalServerError, "Error when get one user by phone number")
	}
	return rs, nil
}
