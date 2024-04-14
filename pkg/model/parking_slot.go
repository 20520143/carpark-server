package model

import (
	"github.com/google/uuid"
	"gitlab.com/goxp/cloud0/ginext"
)

type ParkingSlot struct {
	BaseModel
	Name        string    `json:"name"`
	Description string    `json:"description"`
	BlockID     uuid.UUID `json:"block_id"`
}

func (ParkingSlot) TableName() string {
	return "parking_slot"
}

type ParkingSlotReq struct {
	ID          *uuid.UUID `json:"id"`
	Name        *string    `json:"name"`
	Description *string    `json:"description"`
	BlockID     *uuid.UUID `json:"block_id"`
}

type ListParkingSlotReq struct {
	BlockID  *uuid.UUID `json:"block_id" form:"block_id"`
	Sort     string     `json:"sort" form:"sort"`
	Page     int        `json:"page" form:"page"`
	PageSize int        `json:"page_size" form:"page_size"`
}

type ListParkingSlotRes struct {
	Data []ParkingSlot   `json:"data,omitempty"`
	Meta ginext.BodyMeta `json:"meta" swaggertype:"object"`
}
