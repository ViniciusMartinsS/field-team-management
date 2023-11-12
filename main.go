package main

import (
	"github.com/ViniciusMartinss/field-team-management/application/usecase"
	"github.com/ViniciusMartinss/field-team-management/configuration"
	"github.com/ViniciusMartinss/field-team-management/infrastructure/api"
	"github.com/ViniciusMartinss/field-team-management/infrastructure/encryption"
	"github.com/ViniciusMartinss/field-team-management/infrastructure/jwt"
	"github.com/ViniciusMartinss/field-team-management/infrastructure/repository"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	database := configuration.NewDatabase(
		"field",
		"mysql",
		"root:root@tcp(localhost:3306)/field?multiStatements=true&parseTime=true",
	)

	err := database.Connect()
	if err != nil {
		panic(err)
	}

	err = database.Migrate("./migrations")
	if err != nil {
		panic(err)
	}

	db := database.GetConnection()

	userRepository, err := repository.NewUser(db)
	if err != nil {
		panic(err)
	}

	taskRepository, err := repository.NewTask(db)
	if err != nil {
		panic(err)
	}

	encryptor, err := encryption.New("123456789123456789123456")
	if err != nil {
		panic(err)
	}

	taskUsecase, err := usecase.NewTask(
		taskRepository,
		taskRepository,
		taskRepository,
		taskRepository,
		encryptor,
	)
	if err != nil {
		panic(err)
	}

	authenticator, err := jwt.New("my_secret_key")
	if err != nil {
		panic(err)
	}

	authUsecase, err := usecase.NewAuth(authenticator, userRepository)
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	taskRouter, err := api.NewTask(r, authenticator, taskUsecase)
	if err != nil {
		panic(err)
	}
	taskRouter.CreateRouter()

	authRouter, err := api.NewAuth(r, authUsecase)
	if err != nil {
		panic(err)
	}
	authRouter.CreateRouter()

	err = r.Run()
	if err != nil {
		panic(err)
	} // listen and serve on 0.0.0.0:8080
}
