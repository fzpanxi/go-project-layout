package service

import (
	"context"
	"ddd_demo/internal/domain/entity"
	"ddd_demo/internal/domain/service"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewUserUseCase)

type UserUseCase struct {
	ds *service.UserService
}

//type UserResponse struct {
//	UserID    int64
//	Nickname  string
//	CreatedAt time.Time
//	UpdatedAt time.Time
//}

func NewUserUseCase(ds *service.UserService) *UserUseCase {
	return &UserUseCase{
		ds: ds,
	}
}
func (uc *UserUseCase) GetUser(ctx context.Context, userID int64) (*entity.UserEntity, error) {
	return uc.ds.GetUser(ctx, userID)
}
