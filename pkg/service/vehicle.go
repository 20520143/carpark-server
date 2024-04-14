package service

import (
	"context"
	"github.com/google/uuid"
	"parking-server/pkg/model"
	"parking-server/pkg/repo"
	"parking-server/pkg/utils"
	"parking-server/pkg/valid"
)

type VehicleService struct {
	repo repo.PGInterface
}

func NewVehicleService(repo repo.PGInterface) VehicleInterface {
	return &VehicleService{repo: repo}
}

type VehicleInterface interface {
	CreateVehicle(ctx context.Context, req model.VehicleReq) (*model.Vehicle, error)
	GetListVehicle(ctx context.Context, req model.ListVehicleReq) (model.ListVehicleRes, error)
	GetOneVehicle(ctx context.Context, id uuid.UUID) (model.Vehicle, error)
	UpdateVehicle(ctx context.Context, req model.VehicleReq) (model.Vehicle, error)
	DeleteVehicle(ctx context.Context, id uuid.UUID) error
}

func (s *VehicleService) CreateVehicle(ctx context.Context, req model.VehicleReq) (*model.Vehicle, error) {
	Vehicle := &model.Vehicle{
		Name:   valid.String(req.Name),
		Number: valid.String(req.Number),
		Type:   valid.String(req.Type),
		UserID: valid.UUID(req.UserID),
	}

	if err := s.repo.CreateVehicle(ctx, Vehicle); err != nil {
		return nil, err
	}
	return Vehicle, nil
}

func (s *VehicleService) GetListVehicle(ctx context.Context, req model.ListVehicleReq) (model.ListVehicleRes, error) {
	return s.repo.GetListVehicle(ctx, req)
}

func (s *VehicleService) GetOneVehicle(ctx context.Context, id uuid.UUID) (model.Vehicle, error) {
	return s.repo.GetOneVehicle(ctx, id)
}

func (s *VehicleService) UpdateVehicle(ctx context.Context, req model.VehicleReq) (model.Vehicle, error) {
	Vehicle, err := s.repo.GetOneVehicle(ctx, valid.UUID(req.ID))
	if err != nil {
		return Vehicle, err
	}

	utils.Sync(req, &Vehicle)
	if err := s.repo.UpdateVehicle(ctx, &Vehicle); err != nil {
		return Vehicle, err
	}

	return Vehicle, nil
}

func (s *VehicleService) DeleteVehicle(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteVehicle(ctx, id)
}
