package model

import (
	"github.com/google/uuid"
	"gitlab.com/goxp/cloud0/ginext"
	"time"
)

type ParkingLot struct {
	BaseModel
	Name        string    `json:"name" gorm:"not null"`
	Description string    `json:"description"`
	Address     string    `json:"address"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Lat         string    `json:"lat"`
	Long        string    `json:"long"`
	IsActive    bool      `json:"is_active"`
	CompanyID   uuid.UUID `json:"company_id"`
}

func (ParkingLot) TableName() string {
	return "parking_lot"
}

type ParkingLotReq struct {
	ID          *uuid.UUID `json:"id"`
	Name        *string    `json:"name" valid:"Required"`
	Description *string    `json:"description"`
	Address     *string    `json:"address"`
	StartTime   *time.Time `json:"start_time"`
	EndTime     *time.Time `json:"end_time"`
	Lat         *string    `json:"lat"`
	Long        *string    `json:"long"`
	IsActive    *bool      `json:"is_active"`
	CompanyID   *uuid.UUID `json:"company_id"`
}

type ListParkingLotReq struct {
	Name     *string `json:"name" form:"name"`
	Lat      *string `json:"lat" form:"lat"`
	Long     *string `json:"long" form:"long"`
	IsActive *bool   `json:"is_active" form:"is_active"`
	Sort     string  `json:"sort" form:"sort"`
	Page     int     `json:"page" form:"page"`
	PageSize int     `json:"page_size" form:"page_size"`
}

type ListParkingLotRes struct {
	Data []ParkingLot    `json:"data,omitempty"`
	Meta ginext.BodyMeta `json:"meta" swaggertype:"object"`
}
