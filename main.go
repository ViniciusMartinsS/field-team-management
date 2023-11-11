package main

import (
	"fmt"
	"github.com/ViniciusMartinss/field-team-management/configuration"
	"github.com/ViniciusMartinss/field-team-management/infrastructure/api"
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

	fmt.Println("Hello World!")

	r := gin.Default()

	api.CreateTaskRoutes(r, db)

	err = r.Run()
	if err != nil {
		panic(err)
	} // listen and serve on 0.0.0.0:8080
}
