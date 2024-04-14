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

type BlockHandler struct {
	service service.BlockInterface
}

func NewBlockHandler(service service.BlockInterface) *BlockHandler {
	return &BlockHandler{service: service}
}

// CreateBlock
// @Tags		Block
// @Summary		Create block
// @Security	ApiKeyAuth
// @Accept		json
// @Produce		json
// @Param		x-user-id		header	string				true	"user id"
// @Param		data			body	model.BlockReq		true	"data"
// @Router		/api/v1/block/create [post]
func (h *BlockHandler) CreateBlock(r *ginext.Request) (*ginext.Response, error) {
	log := logger.WithCtx(r.Context(), utils.GetCurrentCaller(h, 0))

	// check x-user-id
	_, err := utils.CurrentUser(r.GinCtx.Request)
	if err != nil {
		log.WithError(err).Error("error_401: Error when get current user")
		return nil, ginext.NewError(http.StatusBadRequest, utils.MessageError()[http.StatusUnauthorized])
	}

	// parse & check valid request
	var req model.BlockReq
	if err := r.GinCtx.BindJSON(&req); err != nil {
		log.WithError(err).Error("error_400: Error when get parse req")
		return nil, ginext.NewError(http.StatusBadRequest, err.Error())
	}
	if err := common.CheckRequireValid(req); err != nil {
		log.WithError(err).Error("error_400: Fail to check require valid: ", err)
		return nil, ginext.NewError(http.StatusBadRequest, err.Error())
	}

	res, err := h.service.CreateBlock(r.Context(), req)
	if err != nil {
		return nil, err
	}

	return &ginext.Response{Code: http.StatusOK, Body: &ginext.GeneralBody{Data: res}}, nil

}

// GetListBlock
// @Tags		Block
// @Summary		Get list block
// @Security	ApiKeyAuth
// @Accept		json
// @Produce		json
// @Param		x-user-id		header		string					true	"user_id"
// @Param		data			query		model.ListBlockReq		true	"data"
// @Success		200				{object}	model.ListBlockRes
// @Router		/api/v1/block/get-list [get]
func (h *BlockHandler) GetListBlock(r *ginext.Request) (*ginext.Response, error) {
	log := logger.WithCtx(r.Context(), utils.GetCurrentCaller(h, 0))

	// check x-user-id
	_, err := utils.CurrentUser(r.GinCtx.Request)
	if err != nil {
		log.WithError(err).Error("error_401: Error when get current user")
		return nil, ginext.NewError(http.StatusBadRequest, utils.MessageError()[http.StatusUnauthorized])
	}

	// parse & check valid request
	var req model.ListBlockReq
	if err := r.GinCtx.BindQuery(&req); err != nil {
		log.WithError(err).Error("error_400: Error when get parse req")
		return nil, ginext.NewError(http.StatusBadRequest, err.Error())
	}
	if err := common.CheckRequireValid(req); err != nil {
		log.WithError(err).Error("error_400: Fail to check require valid: ", err)
		return nil, ginext.NewError(http.StatusBadRequest, err.Error())
	}

	res, err := h.service.GetListBlock(r.Context(), req)
	if err != nil {
		return nil, err
	}

	return &ginext.Response{Code: http.StatusOK, Body: &ginext.GeneralBody{
		Data: res.Data,
		Meta: res.Meta,
	}}, nil
}

// GetOneBlock
// @Tags		Block
// @Summary		Get list block
// @Security	ApiKeyAuth
// @Accept		json
// @Produce		json
// @Param		x-user-id		header		string		true	"user_id"
// @Param		id				path		string		true	"id"
// @Success		200				{object}	model.Block
// @Router		/api/v1/block/get-one/:id 	[get]
func (h *BlockHandler) GetOneBlock(r *ginext.Request) (*ginext.Response, error) {
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

	res, err := h.service.GetOneBlock(r.Context(), valid.UUID(id))
	if err != nil {
		return nil, err
	}

	return &ginext.Response{Code: http.StatusOK, Body: &ginext.GeneralBody{Data: res}}, nil
}

// UpdateBlock
// @Tags		Block
// @Summary		Update block
// @Security	ApiKeyAuth
// @Accept		json
// @Produce		json
// @Param		x-user-id		header		string				true	"user_id"
// @Param		id				path		string				true	"id"
// @Param		data			body		model.BlockReq		true	"data"
// @Success		200				{object}	model.Block
// @Router		/api/v1/block/update/:id 	[put]
func (h *BlockHandler) UpdateBlock(r *ginext.Request) (*ginext.Response, error) {
	log := logger.WithCtx(r.Context(), utils.GetCurrentCaller(h, 0))

	// check x-user-id
	_, err := utils.CurrentUser(r.GinCtx.Request)
	if err != nil {
		log.WithError(err).Error("error_401: Error when get current user")
		return nil, ginext.NewError(http.StatusBadRequest, utils.MessageError()[http.StatusUnauthorized])
	}

	// parse & check valid request
	var req model.BlockReq
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

	res, err := h.service.UpdateBlock(r.Context(), req)
	if err != nil {
		return nil, err
	}

	return &ginext.Response{Code: http.StatusOK, Body: &ginext.GeneralBody{
		Data: res,
	}}, nil
}

// DeleteBlock
// @Tags		Block
// @Summary		Delete block
// @Security	ApiKeyAuth
// @Accept		json
// @Produce		json
// @Param		x-user-id		header		string	true	"user id"
// @Param		id				path		string	true	"id"
// @Success		200				{string}	success
// @Router		/api/v1/block/delete/:id 	[delete]
func (h *BlockHandler) DeleteBlock(r *ginext.Request) (*ginext.Response, error) {
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

	err = h.service.DeleteBlock(r.Context(), valid.UUID(id))
	if err != nil {
		return nil, err
	}

	return &ginext.Response{Code: http.StatusOK, Body: &ginext.GeneralBody{
		Data: "Xóa bản ghi thành công",
	}}, nil
}
