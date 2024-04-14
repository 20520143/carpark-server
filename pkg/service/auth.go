package service

import (
	"context"
	"parking-server/pkg/model"
	"parking-server/pkg/repo"
)

type AuthService struct {
	repo repo.PGInterface
}

func NewAuthService(repo repo.PGInterface) AuthServiceInterface {
	return &AuthService{repo: repo}
}

type AuthServiceInterface interface {
	Login(ctx context.Context, req model.LoginParam) (interface{}, error)
}

func (s *AuthService) Login(ctx context.Context, req model.LoginParam) (interface{}, error) {
	//user, err := s.repo.GetOneUserByPhone(ctx, valid.String(req.UserName), nil)
	//if err != nil {
	//	return nil, err
	//}

	return nil, nil
}
