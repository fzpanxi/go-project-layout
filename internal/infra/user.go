package infra

import (
	"context"
	"ddd_demo/internal/domain/entity"
	"ddd_demo/internal/domain/repository"
)

type userRepo struct {
	data *Data
}

func NewUserRepo(data *Data) repository.UserRepo {
	return &userRepo{
		data: data,
	}
}
func (ur *userRepo) GetUser(ctx context.Context, userID int64) (*entity.UserEntity, error) {
	user, err := ur.data.db.User.Get(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &entity.UserEntity{
		UserID:    user.ID,
		Mobile:    user.Mobile,
		Password:  user.Password,
		Nickname:  user.Nickname,
		Avatar:    user.Avatar,
		CreatedAt: user.CreateTime,
		UpdatedAt: user.UpdateTime,
	}, nil
}
