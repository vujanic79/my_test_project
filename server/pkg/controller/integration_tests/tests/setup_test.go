package tests

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/pkg/errors"
	"github.com/vujanic79/golang-react-todo-app/pkg/logger"
	"log"
	"net/http"
	"os"
	"testing"
	"time"
)

var (
	network *dockertest.Network
)

func TestMain(m *testing.M) {
	l := logger.Get()
	pool, err := dockertest.NewPool("")
	if err != nil {
		l.Fatal().Stack().Err(errors.WithStack(err)).
			Msgf("could not create docker pool: %v", err)
	}

	network, err = pool.CreateNetwork("test-network")
	if err != nil {
		l.Fatal().Stack().Err(errors.WithStack(err)).
			Msgf("Could not create docker network: %v", err)
	}

	pg, err := deployPostgres(pool)
	if err != nil {
		l.Fatal().Stack().Err(errors.WithStack(err)).
			Msgf("Could not start docker postgres: %v", err)
	}

	api, err := deployApi(pool)
	if err != nil {
		l.Fatal().Stack().Err(errors.WithStack(err)).
			Msgf("Could not start docker api: %v", err)
	}

	resources := []*dockertest.Resource{
		pg,
		api,
	}

	exitCode := m.Run()

	err = TearDown(pool, resources)
	if err != nil {
		log.Fatalf("Could not purge resource: %v", err)
	}

	os.Exit(exitCode)
}

func deployPostgres(pool *dockertest.Pool) (*dockertest.Resource, error) {
	l := logger.Get()
	randPass := uuid.New().String()
	randUser := uuid.New().String()
	pg, err := pool.BuildAndRunWithBuildOptions(&dockertest.BuildOptions{
		ContextDir: "../",
		Dockerfile: "./Dockerfile_test_db",
	}, &dockertest.RunOptions{
		Name: "golang-react-todo-app-db-test",
		Env: []string{
			fmt.Sprintf("POSTGRES_PASSWORD=%s", randPass),
			fmt.Sprintf("POSTGRES_USER=%s", randUser),
			"POSTGRES_DB=test_db",
		},
		Networks: []*dockertest.Network{network},
	}, func(cfg *docker.HostConfig) {
		cfg.AutoRemove = true
		cfg.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})

	if err != nil {
		return nil, fmt.Errorf("could not start pg: %v", err)
	}

	port := pg.GetPort("5432/tcp")
	if err := setEnvVars(randUser, randPass); err != nil {
		return nil, err
	}

	dbUrl := fmt.Sprintf("%s://%s:%s@127.0.0.1:%s/%s?sslmode=%s",
		os.Getenv("DB_DRIVER"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		port,
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSL_MODE"))

	l.Info().Msgf("Connecting to database on url: %s", dbUrl)

	if err := pg.Expire(60); err != nil {
		return nil, err
	}

	pool.MaxWait = 60 * time.Second
	if err := pool.Retry(func() error {
		db, err := sql.Open(os.Getenv("DB_DRIVER"), dbUrl)
		if err != nil {
			l.Info().Msg("Waiting for postgres...")
			return err
		}
		return db.Ping()
	}); err != nil {
		l.Fatal().Stack().Err(errors.WithStack(err)).Msg("Could not connect to docker db")
	}

	return pg, nil
}

func setEnvVars(randUser string, randPass string) (err error) {
	err = os.Setenv("DB_DRIVER", "postgres")
	if err != nil {
		return err
	}
	err = os.Setenv("DB_USER", randUser)
	if err != nil {
		return err
	}
	err = os.Setenv("DB_PASSWORD", randPass)
	if err != nil {
		return err
	}
	err = os.Setenv("DB_HOST", "golang-react-todo-app-db-test")
	if err != nil {
		return err
	}
	err = os.Setenv("DB_PORT", "5432")
	if err != nil {
		return err
	}
	err = os.Setenv("DB_NAME", "test_db")
	if err != nil {
		return err
	}
	err = os.Setenv("DB_SSL_MODE", "disable")
	if err != nil {
		return err
	}
	return nil
}

func deployApi(pool *dockertest.Pool) (api *dockertest.Resource, err error) {
	api, err = pool.BuildAndRunWithBuildOptions(&dockertest.BuildOptions{
		ContextDir: "../../../..",
		Dockerfile: "./pkg/controller/integration_tests/Dockerfile_test_api",
	}, &dockertest.RunOptions{
		Name:         "golang-react-todo-app-api-test",
		ExposedPorts: []string{"8000"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"8000/tcp": {{HostIP: "127.0.0.1", HostPort: "8000"}},
		},
		Env: []string{
			"PORT=8000",
			"APP_ENV=development",
			"LOG_LEVEL=debug",
			"HOST=0.0.0.0",
			fmt.Sprintf("DB_DRIVER=%s", os.Getenv("DB_DRIVER")),
			fmt.Sprintf("DB_USER=%s", os.Getenv("DB_USER")),
			fmt.Sprintf("DB_PASSWORD=%s", os.Getenv("DB_PASSWORD")),
			fmt.Sprintf("DB_HOST=%s", os.Getenv("DB_HOST")),
			fmt.Sprintf("DB_PORT=%s", os.Getenv("DB_PORT")),
			fmt.Sprintf("DB_NAME=%s", os.Getenv("DB_NAME")),
			fmt.Sprintf("DB_SSL_MODE=%s", os.Getenv("DB_SSL_MODE")),
		},
		Networks: []*dockertest.Network{
			network,
		},
	})

	if err != nil {
		return nil, fmt.Errorf("could not start api: %v", err)
	}

	err = api.Expire(60)
	if err != nil {
		return nil, err
	}

	err = os.Setenv("PORT", "8000")
	if err != nil {
		return nil, err
	}

	if err = pool.Retry(func() error {
		_, err := http.Get(fmt.Sprintf("http://127.0.0.1:%s/todo/healthz", os.Getenv("PORT")))
		if err != nil {
			fmt.Println("Waiting for API...", err)
			return err
		}

		return nil
	}); err != nil {
		return nil, fmt.Errorf("could not start docker api: %v", err)
	}

	return api, nil
}

func TearDown(pool *dockertest.Pool, resources []*dockertest.Resource) error {
	for _, resource := range resources {
		if err := pool.Purge(resource); err != nil {
			return fmt.Errorf("could not purge resource: %v", err)
		}
	}

	if err := pool.RemoveNetwork(network); err != nil {
		return fmt.Errorf("could not remove network: %v", err)
	}

	return nil
}
