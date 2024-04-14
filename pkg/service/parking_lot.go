package service

import (
	"context"
	"github.com/google/uuid"
	"parking-server/pkg/model"
	"parking-server/pkg/repo"
	"parking-server/pkg/utils"
	"parking-server/pkg/valid"
)

type ParkingLotService struct {
	repo repo.PGInterface
}

func NewParkingLotService(repo repo.PGInterface) ParkingLotInterface {
	return &ParkingLotService{repo: repo}
}

type ParkingLotInterface interface {
	CreateParkingLot(ctx context.Context, req model.ParkingLotReq) (*model.ParkingLot, error)
	GetListParkingLot(ctx context.Context, req model.ListParkingLotReq) (model.ListParkingLotRes, error)
	GetOneParkingLot(ctx context.Context, id uuid.UUID) (model.ParkingLot, error)
	UpdateParkingLot(ctx context.Context, req model.ParkingLotReq) (model.ParkingLot, error)
	DeleteParkingLot(ctx context.Context, id uuid.UUID) error
}

func (s *ParkingLotService) CreateParkingLot(ctx context.Context, req model.ParkingLotReq) (*model.ParkingLot, error) {
	ParkingLot := &model.ParkingLot{
		Name:        valid.String(req.Name),
		Description: valid.String(req.Description),
		Address:     valid.String(req.Address),
		StartTime:   valid.DayTime(req.StartTime),
		EndTime:     valid.DayTime(req.EndTime),
		Lat:         valid.String(req.Lat),
		Long:        valid.String(req.Long),
		IsActive:    valid.Bool(req.IsActive),
		CompanyID:   valid.UUID(req.CompanyID),
	}

	if err := s.repo.CreateParkingLot(ctx, ParkingLot); err != nil {
		return nil, err
	}
	return ParkingLot, nil
}

func (s *ParkingLotService) GetListParkingLot(ctx context.Context, req model.ListParkingLotReq) (model.ListParkingLotRes, error) {
	return s.repo.GetListParkingLot(ctx, req)
}

func (s *ParkingLotService) GetOneParkingLot(ctx context.Context, id uuid.UUID) (model.ParkingLot, error) {
	return s.repo.GetOneParkingLot(ctx, id)
}

func (s *ParkingLotService) UpdateParkingLot(ctx context.Context, req model.ParkingLotReq) (model.ParkingLot, error) {
	ParkingLot, err := s.repo.GetOneParkingLot(ctx, valid.UUID(req.ID))
	if err != nil {
		return ParkingLot, err
	}

	utils.Sync(req, &ParkingLot)
	if err := s.repo.UpdateParkingLot(ctx, &ParkingLot); err != nil {
		return ParkingLot, err
	}

	return ParkingLot, nil
}

func (s *ParkingLotService) DeleteParkingLot(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteParkingLot(ctx, id)
}
