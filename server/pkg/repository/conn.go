package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/vujanic79/golang-react-todo-app/pkg/internal/database"
	"github.com/vujanic79/golang-react-todo-app/pkg/logger"
	"os"
)

func GetPostgreSQLConnection() (db *database.Queries) {
	l := logger.Get()

	driver := os.Getenv("DB_DRIVER")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")
	sslMode := os.Getenv("DB_SSL_MODE")

	dbUrl := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=%s", driver, user, password, host, port, name, sslMode)
	l.Info().Str("dbUrl", dbUrl).Msg("Connecting to database")
	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		l.Fatal().Stack().Err(errors.WithStack(err)).
			Dict("connectionParams", zerolog.Dict().
				Str("driver", driver).
				Str("host", host).
				Str("port", port).
				Str("name", name).
				Str("sslmode", sslMode)).
			Msg("Database connection error")
	}
	db = database.New(conn)
	l.Info().Msg("Database connection established")
	return db
}
