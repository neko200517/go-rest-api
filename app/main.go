package main

import (
	"go-echo/controller"
	"go-echo/db"
	"go-echo/repository"
	"go-echo/router"
	"go-echo/usecase"
	"go-echo/validator"
)

func main() {
	// dbインスタンスを作成
	db := db.NewDB()

	// 依存性の抽入（DI）

	// バリデーション関係のインスタンスを生成
	userValidator := validator.NewUserValidator()
	taskValidator := validator.NewTaskValidator()

	// user関係のインスタンスを生成
	userRepository := repository.NewUserRepository(db)                   // db → repository
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator) // repository → usecase
	userController := controller.NewUserController(userUsecase)          // usecase → controller

	// task関係のインスタンスを生成
	taskRepository := repository.NewTaskRepository(db)                   // db → repository
	taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator) // repository → usecase
	taskController := controller.NewTaskController(taskUsecase)          // usecase → controller

	// routerコンストラクタの起動
	e := router.NewRouter(userController, taskController)

	// サーバーの起動（+エラーのロギング、強制終了、ポート8080）
	e.Logger.Fatal(e.Start(":8080"))
}
