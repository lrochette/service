package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"

	// need for fs migrations
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/jmoiron/sqlx"

	// postgres driver
	_ "github.com/lib/pq"
)

// Config holds db config
type Config struct {
	Driver     string
	Username   string
	Password   string
	Host       string
	DBName     string
	SSLMode    string
	MaxRetries int
	Port       int
}

// Database holds methods to interact with the db
type Database interface {
	Authors
}

type queries struct {
	authorsQueries
}

type database struct {
	db *sqlx.DB
}

// New inits the db
func New(config *Config) (Database, error) {
	db, err := connectToDB(config)
	if err != nil {
		return nil, err
	}
	fmt.Println("Successfully connected to db")

	err = migrateDB(config, db)
	if err != nil {
		return nil, err
	}
	fmt.Println("Successfully ran migrations")

	sqlxDB := sqlx.NewDb(db, config.Driver)

	baseDB := database{
		db: sqlxDB,
	}

	return &queries{
		authorsQueries{baseDB},
	}, nil
}

func makeConnectionString(config *Config) string {
	dbURL := "postgresql://%s:%s@%s:%d/%s?sslmode=%s"
	return fmt.Sprintf(
		dbURL,
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.DBName,
		config.SSLMode)
}

func connectToDB(config *Config) (*sql.DB, error) {
	dbURL := makeConnectionString(config)
	// retry connection while db is booting
	for retries := 0; retries < config.MaxRetries; retries++ {

		db, _ := sql.Open(config.Driver, dbURL)
		err := db.Ping()
		if err == nil {
			return db, nil
		}

		fmt.Printf("Retrying db connection %d/%d times: %s",
			retries, config.MaxRetries, err)

		time.Sleep(1 * time.Second)
	}

	return sql.Open(config.Driver, dbURL)
}

func migrateDB(config *Config, db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}

	migrations, err := migrate.NewWithDatabaseInstance(
		"file://./migrations",
		config.Driver,
		driver,
	)
	if err != nil {
		return err
	}

	if err := migrations.Up(); err != migrate.ErrNoChange {
		return err
	}

	return nil
}
