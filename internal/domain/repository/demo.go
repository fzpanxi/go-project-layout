package repository

import "ddd_demo/internal/domain/entity"

type DemoRepo interface {
	ListDemo() ([]*entity.DemoEntity, error)
}
