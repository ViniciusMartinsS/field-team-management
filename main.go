package main

import (
	"fmt"
	"github.com/ViniciusMartinss/field-team-management/infrastructure/api"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello World!")

	r := gin.Default()

	api.CreateTaskRoutes(r)

	err := r.Run()
	if err != nil {
		panic(err)
	} // listen and serve on 0.0.0.0:8080
}
