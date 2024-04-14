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

type VehicleHandler struct {
	service service.VehicleInterface
}

func NewVehicleHandler(service service.VehicleInterface) *VehicleHandler {
	return &VehicleHandler{service: service}
}

// CreateVehicle
// @Tags		Vehicle
// @Summary		Create Vehicle
// @Security	ApiKeyAuth
// @Accept		json
// @Produce		json
// @Param		x-user-id		header	string					true	"user id"
// @Param		data			body	model.VehicleReq	true	"data"
// @Router		/api/v1/vehicle/create [post]
func (h *VehicleHandler) CreateVehicle(r *ginext.Request) (*ginext.Response, error) {
	log := logger.WithCtx(r.Context(), utils.GetCurrentCaller(h, 0))

	// check x-user-id
	_, err := utils.CurrentUser(r.GinCtx.Request)
	if err != nil {
		log.WithError(err).Error("error_401: Error when get current user")
		return nil, ginext.NewError(http.StatusBadRequest, utils.MessageError()[http.StatusUnauthorized])
	}

	// parse & check valid request
	var req model.VehicleReq
	if err := r.GinCtx.BindJSON(&req); err != nil {
		log.WithError(err).Error("error_400: Error when get parse req")
		return nil, ginext.NewError(http.StatusBadRequest, err.Error())
	}
	if err := common.CheckRequireValid(req); err != nil {
		log.WithError(err).Error("error_400: Fail to check require valid: ", err)
		return nil, ginext.NewError(http.StatusBadRequest, err.Error())
	}

	res, err := h.service.CreateVehicle(r.Context(), req)
	if err != nil {
		return nil, err
	}

	return &ginext.Response{Code: http.StatusOK, Body: &ginext.GeneralBody{Data: res}}, nil

}

// GetListVehicle
// @Tags		Vehicle
// @Summary		Get list Vehicle
// @Security	ApiKeyAuth
// @Accept		json
// @Produce		json
// @Param		x-user-id		header		string					true	"user_id"
// @Param		data			query		model.ListVehicleReq		true	"data"
// @Success		200				{object}	model.ListVehicleRes
// @Router		/api/v1/vehicle/get-list [get]
func (h *VehicleHandler) GetListVehicle(r *ginext.Request) (*ginext.Response, error) {
	log := logger.WithCtx(r.Context(), utils.GetCurrentCaller(h, 0))

	// check x-user-id
	_, err := utils.CurrentUser(r.GinCtx.Request)
	if err != nil {
		log.WithError(err).Error("error_401: Error when get current user")
		return nil, ginext.NewError(http.StatusBadRequest, utils.MessageError()[http.StatusUnauthorized])
	}

	// parse & check valid request
	var req model.ListVehicleReq
	if err := r.GinCtx.BindQuery(&req); err != nil {
		log.WithError(err).Error("error_400: Error when get parse req")
		return nil, ginext.NewError(http.StatusBadRequest, err.Error())
	}
	if err := common.CheckRequireValid(req); err != nil {
		log.WithError(err).Error("error_400: Fail to check require valid: ", err)
		return nil, ginext.NewError(http.StatusBadRequest, err.Error())
	}

	res, err := h.service.GetListVehicle(r.Context(), req)
	if err != nil {
		return nil, err
	}

	return &ginext.Response{Code: http.StatusOK, Body: &ginext.GeneralBody{
		Data: res.Data,
		Meta: res.Meta,
	}}, nil
}

// GetOneVehicle
// @Tags		Vehicle
// @Summary		Get list Vehicle
// @Security	ApiKeyAuth
// @Accept		json
// @Produce		json
// @Param		x-user-id		header		string		true	"user_id"
// @Param		id				path		string		true	"id"
// @Success		200				{object}	model.Vehicle
// @Router		/api/v1/vehicle/get-one/:id 	[get]
func (h *VehicleHandler) GetOneVehicle(r *ginext.Request) (*ginext.Response, error) {
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

	res, err := h.service.GetOneVehicle(r.Context(), valid.UUID(id))
	if err != nil {
		return nil, err
	}

	return &ginext.Response{Code: http.StatusOK, Body: &ginext.GeneralBody{Data: res}}, nil
}

// UpdateVehicle
// @Tags		Vehicle
// @Summary		Update Vehicle
// @Security	ApiKeyAuth
// @Accept		json
// @Produce		json
// @Param		x-user-id		header		string				true	"user_id"
// @Param		id				path		string				true	"id"
// @Param		data			body		model.VehicleReq		true	"data"
// @Success		200				{object}	model.Vehicle
// @Router		/api/v1/vehicle/update/:id 	[put]
func (h *VehicleHandler) UpdateVehicle(r *ginext.Request) (*ginext.Response, error) {
	log := logger.WithCtx(r.Context(), utils.GetCurrentCaller(h, 0))

	// check x-user-id
	_, err := utils.CurrentUser(r.GinCtx.Request)
	if err != nil {
		log.WithError(err).Error("error_401: Error when get current user")
		return nil, ginext.NewError(http.StatusBadRequest, utils.MessageError()[http.StatusUnauthorized])
	}

	// parse & check valid request
	var req model.VehicleReq
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

	res, err := h.service.UpdateVehicle(r.Context(), req)
	if err != nil {
		return nil, err
	}

	return &ginext.Response{Code: http.StatusOK, Body: &ginext.GeneralBody{
		Data: res,
	}}, nil
}

// DeleteVehicle
// @Tags		Vehicle
// @Summary		Delete Vehicle
// @Security	ApiKeyAuth
// @Accept		json
// @Produce		json
// @Param		x-user-id		header		string	true	"user id"
// @Param		id				path		string	true	"id"
// @Success		200				{string}	success
// @Router		/api/v1/vehicle/delete/:id 	[delete]
func (h *VehicleHandler) DeleteVehicle(r *ginext.Request) (*ginext.Response, error) {
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

	err = h.service.DeleteVehicle(r.Context(), valid.UUID(id))
	if err != nil {
		return nil, err
	}

	return &ginext.Response{Code: http.StatusOK, Body: &ginext.GeneralBody{
		Data: "Xóa bản ghi thành công",
	}}, nil
}
