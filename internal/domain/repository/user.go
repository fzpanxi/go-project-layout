package repository

import (
	"context"
	"ddd_demo/internal/domain/entity"
)

type UserRepo interface {
	GetUser(ctx context.Context, userID int64) (*entity.UserEntity, error)
}
