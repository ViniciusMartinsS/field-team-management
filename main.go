package main

import (
	"context"
	"github.com/ViniciusMartinss/field-team-management/application/usecase"
	"github.com/ViniciusMartinss/field-team-management/configuration"
	"github.com/ViniciusMartinss/field-team-management/infrastructure/api"
	"github.com/ViniciusMartinss/field-team-management/infrastructure/encryption"
	"github.com/ViniciusMartinss/field-team-management/infrastructure/jwt"
	"github.com/ViniciusMartinss/field-team-management/infrastructure/notifier"
	"github.com/ViniciusMartinss/field-team-management/infrastructure/repository"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const (
	queueEnvKey         = "QUEUE"
	queueConnKey        = "QUEUE_CONN_STRING"
	dbEnvKey            = "DATABASE"
	dbConnKey           = "DATABASE_CONN_STRING"
	encryptionSecretKey = "ENCRYPTION_KEY"
	jwtSecretKey        = "JWT_KEY"

	databaseDriver = "mysql"
)

func main() {
	queueName := os.Getenv(queueEnvKey)
	queueConn := os.Getenv(queueConnKey)

	dbName := os.Getenv(dbEnvKey)
	dbConn := os.Getenv(dbConnKey)

	encryptionSecret := os.Getenv(encryptionSecretKey)
	jwtSecret := os.Getenv(jwtSecretKey)

	rabbitmq := configuration.NewRabbitmq(queueName, queueConn)
	brokerConn, err := rabbitmq.Connect()
	if err != nil {
		panic(err)
	}
	defer brokerConn.Close()

	brokerCh, err := rabbitmq.CreateChannelAndQueue(brokerConn)
	if err != nil {
		panic(err)
	}
	defer brokerCh.Close()

	database := configuration.NewDatabase(
		dbName,
		databaseDriver,
		dbConn,
	)

	err = database.Connect()
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

	taskNotifier := notifier.NewNotifier(queueName, brokerCh)

	encryptor, err := encryption.New(encryptionSecret)
	if err != nil {
		panic(err)
	}

	taskUsecase, err := usecase.NewTask(
		taskRepository,
		taskRepository,
		taskRepository,
		taskRepository,
		encryptor,
		taskNotifier,
	)
	if err != nil {
		panic(err)
	}

	authenticator, err := jwt.New(jwtSecret)
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

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("error to start app on: %s\n", err)
		}
	}()

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		os.Kill,
		syscall.SIGTERM,
	)

	<-ctx.Done()
	log.Printf("server shutting down")

	err = server.Shutdown(ctx)
	if err != nil {
		log.Fatalf("error to gracefully shut down, reason: %s\n", err.Error())
	}

	stop()
}
