package main

import (
	"context"
	config "ddd_demo/internal/conf"
	"log"
)

func main() {
	//data, cleanup, err := infra.NewData()
	//if err != nil {
	//	panic(err)
	//}
	//defer cleanup()
	//userRepo := infra.NewUserRepo(data)
	//userService := service.NewUserService(userRepo)
	//userUseCase := application.NewUserUseCase(userService)
	bc, err := config.LoadConfig("./configs/config.yaml")
	if err != nil {
		panic(err)
	}
	userUseCase, cleanup, err := initApp(bc)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	log.Println(userUseCase.GetUser(context.Background(), 0))
}
