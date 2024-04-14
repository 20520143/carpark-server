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

type ParkingSlotHandler struct {
	service service.ParkingSlotInterface
}

func NewParkingSlotHandler(service service.ParkingSlotInterface) *ParkingSlotHandler {
	return &ParkingSlotHandler{service: service}
}

// CreateParkingSlot
// @Tags		ParkingSlot
// @Summary		Create ParkingSlot
// @Security	ApiKeyAuth
// @Accept		json
// @Produce		json
// @Param		x-user-id		header	string					true	"user id"
// @Param		data			body	model.ParkingSlotReq	true	"data"
// @Router		/api/v1/parking-slot/create [post]
func (h *ParkingSlotHandler) CreateParkingSlot(r *ginext.Request) (*ginext.Response, error) {
	log := logger.WithCtx(r.Context(), utils.GetCurrentCaller(h, 0))

	// check x-user-id
	_, err := utils.CurrentUser(r.GinCtx.Request)
	if err != nil {
		log.WithError(err).Error("error_401: Error when get current user")
		return nil, ginext.NewError(http.StatusBadRequest, utils.MessageError()[http.StatusUnauthorized])
	}

	// parse & check valid request
	var req model.ParkingSlotReq
	if err := r.GinCtx.BindJSON(&req); err != nil {
		log.WithError(err).Error("error_400: Error when get parse req")
		return nil, ginext.NewError(http.StatusBadRequest, err.Error())
	}
	if err := common.CheckRequireValid(req); err != nil {
		log.WithError(err).Error("error_400: Fail to check require valid: ", err)
		return nil, ginext.NewError(http.StatusBadRequest, err.Error())
	}

	res, err := h.service.CreateParkingSlot(r.Context(), req)
	if err != nil {
		return nil, err
	}

	return &ginext.Response{Code: http.StatusOK, Body: &ginext.GeneralBody{Data: res}}, nil

}

// GetListParkingSlot
// @Tags		ParkingSlot
// @Summary		Get list ParkingSlot
// @Security	ApiKeyAuth
// @Accept		json
// @Produce		json
// @Param		x-user-id		header		string					true	"user_id"
// @Param		data			query		model.ListParkingSlotReq		true	"data"
// @Success		200				{object}	model.ListParkingSlotRes
// @Router		/api/v1/parking-slot/get-list [get]
func (h *ParkingSlotHandler) GetListParkingSlot(r *ginext.Request) (*ginext.Response, error) {
	log := logger.WithCtx(r.Context(), utils.GetCurrentCaller(h, 0))

	// check x-user-id
	_, err := utils.CurrentUser(r.GinCtx.Request)
	if err != nil {
		log.WithError(err).Error("error_401: Error when get current user")
		return nil, ginext.NewError(http.StatusBadRequest, utils.MessageError()[http.StatusUnauthorized])
	}

	// parse & check valid request
	var req model.ListParkingSlotReq
	if err := r.GinCtx.BindQuery(&req); err != nil {
		log.WithError(err).Error("error_400: Error when get parse req")
		return nil, ginext.NewError(http.StatusBadRequest, err.Error())
	}
	if err := common.CheckRequireValid(req); err != nil {
		log.WithError(err).Error("error_400: Fail to check require valid: ", err)
		return nil, ginext.NewError(http.StatusBadRequest, err.Error())
	}

	res, err := h.service.GetListParkingSlot(r.Context(), req)
	if err != nil {
		return nil, err
	}

	return &ginext.Response{Code: http.StatusOK, Body: &ginext.GeneralBody{
		Data: res.Data,
		Meta: res.Meta,
	}}, nil
}

// GetOneParkingSlot
// @Tags		ParkingSlot
// @Summary		Get list ParkingSlot
// @Security	ApiKeyAuth
// @Accept		json
// @Produce		json
// @Param		x-user-id		header		string		true	"user_id"
// @Param		id				path		string		true	"id"
// @Success		200				{object}	model.ParkingSlot
// @Router		/api/v1/parking-slot/get-one/:id 	[get]
func (h *ParkingSlotHandler) GetOneParkingSlot(r *ginext.Request) (*ginext.Response, error) {
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

	res, err := h.service.GetOneParkingSlot(r.Context(), valid.UUID(id))
	if err != nil {
		return nil, err
	}

	return &ginext.Response{Code: http.StatusOK, Body: &ginext.GeneralBody{Data: res}}, nil
}

// UpdateParkingSlot
// @Tags		ParkingSlot
// @Summary		Update ParkingSlot
// @Security	ApiKeyAuth
// @Accept		json
// @Produce		json
// @Param		x-user-id		header		string				true	"user_id"
// @Param		id				path		string				true	"id"
// @Param		data			body		model.ParkingSlotReq		true	"data"
// @Success		200				{object}	model.ParkingSlot
// @Router		/api/v1/parking-slot/update/:id 	[put]
func (h *ParkingSlotHandler) UpdateParkingSlot(r *ginext.Request) (*ginext.Response, error) {
	log := logger.WithCtx(r.Context(), utils.GetCurrentCaller(h, 0))

	// check x-user-id
	_, err := utils.CurrentUser(r.GinCtx.Request)
	if err != nil {
		log.WithError(err).Error("error_401: Error when get current user")
		return nil, ginext.NewError(http.StatusBadRequest, utils.MessageError()[http.StatusUnauthorized])
	}

	// parse & check valid request
	var req model.ParkingSlotReq
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

	res, err := h.service.UpdateParkingSlot(r.Context(), req)
	if err != nil {
		return nil, err
	}

	return &ginext.Response{Code: http.StatusOK, Body: &ginext.GeneralBody{
		Data: res,
	}}, nil
}

// DeleteParkingSlot
// @Tags		ParkingSlot
// @Summary		Delete ParkingSlot
// @Security	ApiKeyAuth
// @Accept		json
// @Produce		json
// @Param		x-user-id		header		string	true	"user id"
// @Param		id				path		string	true	"id"
// @Success		200				{string}	success
// @Router		/api/v1/parking-slot/delete/:id 	[delete]
func (h *ParkingSlotHandler) DeleteParkingSlot(r *ginext.Request) (*ginext.Response, error) {
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

	err = h.service.DeleteParkingSlot(r.Context(), valid.UUID(id))
	if err != nil {
		return nil, err
	}

	return &ginext.Response{Code: http.StatusOK, Body: &ginext.GeneralBody{
		Data: "Xóa bản ghi thành công",
	}}, nil
}
