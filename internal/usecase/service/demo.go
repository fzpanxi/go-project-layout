package service

import (
	"ddd_demo/internal/domain/service"
	"time"
)

type DemoUseCase struct {
	ds *service.DemoService
}

type DemoResponse struct {
	Id        int64
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewDemoUseCase(ds *service.DemoService) *DemoUseCase {
	return &DemoUseCase{
		ds: ds,
	}
}
func (uc *DemoUseCase) ListDemo() []*DemoResponse {
	ps, _ := uc.ds.ListDemo()
	result := make([]*DemoResponse, 0)
	for _, p := range ps {
		result = append(result, &DemoResponse{
			Id:      p.Id,
			Content: p.Content,
		})
	}
	return result
}
