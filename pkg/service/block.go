package service

import (
	"context"
	"github.com/google/uuid"
	"parking-server/pkg/model"
	"parking-server/pkg/repo"
	"parking-server/pkg/utils"
	"parking-server/pkg/valid"
)

type BlockService struct {
	repo repo.PGInterface
}

func NewBlockService(repo repo.PGInterface) BlockInterface {
	return &BlockService{repo: repo}
}

type BlockInterface interface {
	CreateBlock(ctx context.Context, req model.BlockReq) (*model.Block, error)
	GetListBlock(ctx context.Context, req model.ListBlockReq) (model.ListBlockRes, error)
	GetOneBlock(ctx context.Context, id uuid.UUID) (model.Block, error)
	UpdateBlock(ctx context.Context, req model.BlockReq) (model.Block, error)
	DeleteBlock(ctx context.Context, id uuid.UUID) error
}

func (s *BlockService) CreateBlock(ctx context.Context, req model.BlockReq) (*model.Block, error) {
	block := &model.Block{
		Code:         valid.String(req.Code),
		Description:  valid.String(req.Description),
		Slot:         valid.Int(req.Slot),
		ParkingLotID: valid.UUID(req.ParkingLotID),
	}

	if err := s.repo.CreateBlock(ctx, block); err != nil {
		return nil, err
	}
	return block, nil
}

func (s *BlockService) GetListBlock(ctx context.Context, req model.ListBlockReq) (model.ListBlockRes, error) {
	return s.repo.GetListBlock(ctx, req)
}

func (s *BlockService) GetOneBlock(ctx context.Context, id uuid.UUID) (model.Block, error) {
	return s.repo.GetOneBlock(ctx, id)
}

func (s *BlockService) UpdateBlock(ctx context.Context, req model.BlockReq) (model.Block, error) {
	block, err := s.repo.GetOneBlock(ctx, valid.UUID(req.ID))
	if err != nil {
		return block, err
	}

	utils.Sync(req, &block)
	if err := s.repo.UpdateBlock(ctx, &block); err != nil {
		return block, err
	}

	return block, nil
}

func (s *BlockService) DeleteBlock(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteBlock(ctx, id)
}
