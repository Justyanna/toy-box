package repository

import (
	"database/sql"
	"fmt"
	"log"

	"app-server/internal/utilities"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func InitDB() *Database {

	db_user := utilities.GetDatabaseUser()
	db_password := utilities.GetDatabasePassword()
	db_name := utilities.GetDatabaseName()
	db_port := utilities.GetDatabasePort()
	connStr := fmt.Sprintf("postgres://%s:%s@db:%s/%s?sslmode=disable", db_user, db_password, db_port, db_name)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to db")
	return &Database{db: db}
}

func PerformMigration(d *Database) {
	driver, err := postgres.WithInstance(d.db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}
	migrationsDir := "/app/app-server/internal/repository/migrations"
	m, err := migrate.NewWithDatabaseInstance(
		"file://"+migrationsDir,
		"postgres", driver)
	if err != nil {
		log.Fatal(err)
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}
