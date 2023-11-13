package configuration

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"path/filepath"
)

type Database struct {
	connection       *sql.DB
	driver           database.Driver
	databaseName     string
	driverName       string
	connectionString string
}

func NewDatabase(databaseName, driverName, connectionString string) *Database {
	return &Database{
		databaseName:     databaseName,
		driverName:       driverName,
		connectionString: connectionString,
	}
}

func (d *Database) Connect() error {
	db, err := sql.Open(d.driverName, d.connectionString)
	if err != nil {
		return err
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return err
	}

	d.connection = db
	d.driver = driver

	return nil
}

func (d *Database) GetConnection() *sqlx.DB {
	return sqlx.NewDb(d.connection, d.driverName)
}

func (d *Database) Migrate(path string) error {
	migrations, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("file:///%s", migrations), d.databaseName, d.driver)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil {
		if err.Error() != "no change" {
			return err
		}
	}

	return nil
}
