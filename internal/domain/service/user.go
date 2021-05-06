package service

import (
	"context"
	"ddd_demo/internal/domain/entity"
	"ddd_demo/internal/domain/repository"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewUserService)

type UserService struct {
	repo repository.UserRepo
}

func NewUserService(repo repository.UserRepo) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUser(ctx context.Context, userID int64) (ds *entity.UserEntity, err error) {
	ds, err = s.repo.GetUser(ctx, userID)
	return
}
