package service

import (
	"context"
	"ddd_demo/internal/domain/service"
	"github.com/google/wire"
	"time"
)

var ProviderSet = wire.NewSet(NewUserUseCase)

type UserUseCase struct {
	ds *service.UserService
}

type UserResponse struct {
	UserID    int64
	Nickname  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUserUseCase(ds *service.UserService) *UserUseCase {
	return &UserUseCase{
		ds: ds,
	}
}
func (uc *UserUseCase) GetUser(ctx context.Context, userID int64) (*UserResponse, error) {
	user, err := uc.ds.GetUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &UserResponse{
		UserID:    user.UserID,
		Nickname:  user.Nickname,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
