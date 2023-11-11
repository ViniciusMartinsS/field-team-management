package main

import (
	"database/sql"
	"fmt"
	"github.com/ViniciusMartinss/field-team-management/infrastructure/api"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/jmoiron/sqlx"
	"path/filepath"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/field?multiStatements=true&parseTime=true")
	if err != nil {
		panic(err)
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		panic(err)
	}

	migrations, err := filepath.Abs("./migrations")
	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("file:///%s", migrations), "field", driver)
	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil {
		if err.Error() != "no change" {
			panic(err)
		}
	}

	fmt.Println("Hello World!")

	r := gin.Default()

	api.CreateTaskRoutes(r, sqlx.NewDb(db, "mysql"))

	err = r.Run()
	if err != nil {
		panic(err)
	} // listen and serve on 0.0.0.0:8080
}
