package handlers

import (
	"github.com/praslar/lib/common"
	"gitlab.com/goxp/cloud0/ginext"
	"gitlab.com/goxp/cloud0/logger"
	"net/http"
	"parking-server/pkg/model"
	"parking-server/pkg/service"
	"parking-server/pkg/utils"
	"parking-server/pkg/valid"
)

type ParkingLotHandler struct {
	service service.ParkingLotInterface
}

func NewParkingLotHandler(service service.ParkingLotInterface) *ParkingLotHandler {
	return &ParkingLotHandler{service: service}
}

// CreateParkingLot
// @Tags		ParkingLot
// @Summary		Create ParkingLot
// @Security	ApiKeyAuth
// @Accept		json
// @Produce		json
// @Param		x-user-id		header	string					true	"user id"
// @Param		data			body	model.ParkingLotReq	true	"data"
// @Router		/api/v1/parking-lot/create [post]
func (h *ParkingLotHandler) CreateParkingLot(r *ginext.Request) (*ginext.Response, error) {
	log := logger.WithCtx(r.Context(), utils.GetCurrentCaller(h, 0))

	// check x-user-id
	_, err := utils.CurrentUser(r.GinCtx.Request)
	if err != nil {
		log.WithError(err).Error("error_401: Error when get current user")
		return nil, ginext.NewError(http.StatusBadRequest, utils.MessageError()[http.StatusUnauthorized])
	}

	// parse & check valid request
	var req model.ParkingLotReq
	if err := r.GinCtx.BindJSON(&req); err != nil {
		log.WithError(err).Error("error_400: Error when get parse req")
		return nil, ginext.NewError(http.StatusBadRequest, err.Error())
	}
	if err := common.CheckRequireValid(req); err != nil {
		log.WithError(err).Error("error_400: Fail to check require valid: ", err)
		return nil, ginext.NewError(http.StatusBadRequest, err.Error())
	}

	res, err := h.service.CreateParkingLot(r.Context(), req)
	if err != nil {
		return nil, err
	}

	return &ginext.Response{Code: http.StatusOK, Body: &ginext.GeneralBody{Data: res}}, nil

}

// GetListParkingLot
// @Tags		ParkingLot
// @Summary		Get list ParkingLot
// @Security	ApiKeyAuth
// @Accept		json
// @Produce		json
// @Param		x-user-id		header		string					true	"user_id"
// @Param		data			query		model.ListParkingLotReq		true	"data"
// @Success		200				{object}	model.ListParkingLotRes
// @Router		/api/v1/parking-lot/get-list [get]
func (h *ParkingLotHandler) GetListParkingLot(r *ginext.Request) (*ginext.Response, error) {
	log := logger.WithCtx(r.Context(), utils.GetCurrentCaller(h, 0))

	// check x-user-id
	_, err := utils.CurrentUser(r.GinCtx.Request)
	if err != nil {
		log.WithError(err).Error("error_401: Error when get current user")
		return nil, ginext.NewError(http.StatusBadRequest, utils.MessageError()[http.StatusUnauthorized])
	}

	// parse & check valid request
	var req model.ListParkingLotReq
	if err := r.GinCtx.BindQuery(&req); err != nil {
		log.WithError(err).Error("error_400: Error when get parse req")
		return nil, ginext.NewError(http.StatusBadRequest, err.Error())
	}
	if err := common.CheckRequireValid(req); err != nil {
		log.WithError(err).Error("error_400: Fail to check require valid: ", err)
		return nil, ginext.NewError(http.StatusBadRequest, err.Error())
	}

	res, err := h.service.GetListParkingLot(r.Context(), req)
	if err != nil {
		return nil, err
	}

	return &ginext.Response{Code: http.StatusOK, Body: &ginext.GeneralBody{
		Data: res.Data,
		Meta: res.Meta,
	}}, nil
}

// GetOneParkingLot
// @Tags		ParkingLot
// @Summary		Get list ParkingLot
// @Security	ApiKeyAuth
// @Accept		json
// @Produce		json
// @Param		x-user-id		header		string		true	"user_id"
// @Param		id				path		string		true	"id"
// @Success		200				{object}	model.ParkingLot
// @Router		/api/v1/parking-lot/get-one/:id 	[get]
func (h *ParkingLotHandler) GetOneParkingLot(r *ginext.Request) (*ginext.Response, error) {
	log := logger.WithCtx(r.Context(), utils.GetCurrentCaller(h, 0))

	// check x-user-id
	_, err := utils.CurrentUser(r.GinCtx.Request)
	if err != nil {
		log.WithError(err).Error("error_401: Error when get current user")
		return nil, ginext.NewError(http.StatusBadRequest, utils.MessageError()[http.StatusUnauthorized])
	}

	// parse id
	id := utils.ParseIDFromUri(r.GinCtx)
	if id == nil {
		log.Error("error_400: Wrong id ")
		return nil, ginext.NewError(http.StatusBadRequest, "Wrong id")
	}

	res, err := h.service.GetOneParkingLot(r.Context(), valid.UUID(id))
	if err != nil {
		return nil, err
	}

	return &ginext.Response{Code: http.StatusOK, Body: &ginext.GeneralBody{Data: res}}, nil
}

// UpdateParkingLot
// @Tags		ParkingLot
// @Summary		Update ParkingLot
// @Security	ApiKeyAuth
// @Accept		json
// @Produce		json
// @Param		x-user-id		header		string				true	"user_id"
// @Param		id				path		string				true	"id"
// @Param		data			body		model.ParkingLotReq		true	"data"
// @Success		200				{object}	model.ParkingLot
// @Router		/api/v1/parking-lot/update/:id 	[put]
func (h *ParkingLotHandler) UpdateParkingLot(r *ginext.Request) (*ginext.Response, error) {
	log := logger.WithCtx(r.Context(), utils.GetCurrentCaller(h, 0))

	// check x-user-id
	_, err := utils.CurrentUser(r.GinCtx.Request)
	if err != nil {
		log.WithError(err).Error("error_401: Error when get current user")
		return nil, ginext.NewError(http.StatusBadRequest, utils.MessageError()[http.StatusUnauthorized])
	}

	// parse & check valid request
	var req model.ParkingLotReq
	if err := r.GinCtx.BindJSON(&req); err != nil {
		log.WithError(err).Error("error_400: Error when get parse req")
		return nil, ginext.NewError(http.StatusBadRequest, err.Error())
	}
	if err := common.CheckRequireValid(req); err != nil {
		log.WithError(err).Error("error_400: Fail to check require valid: ", err)
		return nil, ginext.NewError(http.StatusBadRequest, err.Error())
	}

	// parse id
	req.ID = utils.ParseIDFromUri(r.GinCtx)
	if req.ID == nil {
		log.Error("error_400: Wrong id ")
		return nil, ginext.NewError(http.StatusBadRequest, "Wrong id")
	}

	res, err := h.service.UpdateParkingLot(r.Context(), req)
	if err != nil {
		return nil, err
	}

	return &ginext.Response{Code: http.StatusOK, Body: &ginext.GeneralBody{
		Data: res,
	}}, nil
}

// DeleteParkingLot
// @Tags		ParkingLot
// @Summary		Delete ParkingLot
// @Security	ApiKeyAuth
// @Accept		json
// @Produce		json
// @Param		x-user-id		header		string	true	"user id"
// @Param		id				path		string	true	"id"
// @Success		200				{string}	success
// @Router		/api/v1/parking-lot/delete/:id 	[delete]
func (h *ParkingLotHandler) DeleteParkingLot(r *ginext.Request) (*ginext.Response, error) {
	log := logger.WithCtx(r.Context(), utils.GetCurrentCaller(h, 0))

	// check x-user-id
	_, err := utils.CurrentUser(r.GinCtx.Request)
	if err != nil {
		log.WithError(err).Error("error_401: Error when get current user")
		return nil, ginext.NewError(http.StatusBadRequest, utils.MessageError()[http.StatusUnauthorized])
	}

	// parse id
	id := utils.ParseIDFromUri(r.GinCtx)
	if id == nil {
		log.Error("error_400: Wrong id ")
		return nil, ginext.NewError(http.StatusBadRequest, "ID không hợp lệ")
	}

	err = h.service.DeleteParkingLot(r.Context(), valid.UUID(id))
	if err != nil {
		return nil, err
	}

	return &ginext.Response{Code: http.StatusOK, Body: &ginext.GeneralBody{
		Data: "Xóa bản ghi thành công",
	}}, nil
}
