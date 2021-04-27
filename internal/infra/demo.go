package infra

import (
	"ddd_demo/internal/domain/entity"
	"ddd_demo/internal/domain/repository"
)

type demoRepo struct {
}

func NewDemoRepo() repository.DemoRepo {
	return &demoRepo{}
}
func (dr *demoRepo) ListDemo() ([]*entity.DemoEntity, error) {
	rv := make([]*entity.DemoEntity, 0)
	rv = append(rv, &entity.DemoEntity{
		Id:      111,
		Content: "ddd",
	})
	return rv, nil
}
