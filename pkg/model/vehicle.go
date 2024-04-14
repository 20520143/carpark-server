package model

import (
	"github.com/google/uuid"
	"gitlab.com/goxp/cloud0/ginext"
)

type Vehicle struct {
	BaseModel
	Name   string    `json:"name"`
	Number string    `json:"number"`
	Type   string    `json:"type"`
	UserID uuid.UUID `json:"user_id" gorm:"not null"`
}

func (Vehicle) TableName() string {
	return "parking_slot"
}

type VehicleReq struct {
	ID     *uuid.UUID `json:"id"`
	Name   *string    `json:"name"`
	Number *string    `json:"number"`
	Type   *string    `json:"type"`
	UserID *uuid.UUID `json:"user_id"`
}

type ListVehicleReq struct {
	UserID   *uuid.UUID `json:"user_id" form:"block_id"`
	Type     *string    `json:"type" form:"type"`
	Sort     string     `json:"sort" form:"sort"`
	Page     int        `json:"page" form:"page"`
	PageSize int        `json:"page_size" form:"page_size"`
}

type ListVehicleRes struct {
	Data []Vehicle       `json:"data,omitempty"`
	Meta ginext.BodyMeta `json:"meta" swaggertype:"object"`
}
