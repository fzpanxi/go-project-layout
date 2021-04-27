package service

import (
	"ddd_demo/internal/domain/entity"
	"ddd_demo/internal/domain/repository"
)

type DemoService struct {
	repo repository.DemoRepo
}

func NewDemoService(repo repository.DemoRepo) *DemoService {
	return &DemoService{repo: repo}
}

func (s *DemoService) ListDemo() (ds []*entity.DemoEntity, err error) {
	ds, err = s.repo.ListDemo()
	return
}
