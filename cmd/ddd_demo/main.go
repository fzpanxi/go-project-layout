package main

import (
	"ddd_demo/internal/domain/service"
	"ddd_demo/internal/infra"
	application "ddd_demo/internal/usecase/service"
	"log"
)

func main() {
	//TODO repo依赖的store
	demoRepo := infra.NewDemoRepo()
	demoService := service.NewDemoService(demoRepo)
	demoUseCase := application.NewDemoUseCase(demoService)
	for _, v := range demoUseCase.ListDemo() {
		log.Println(v.Id, v.Content)
	}
}
