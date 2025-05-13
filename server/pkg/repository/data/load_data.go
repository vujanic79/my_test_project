package data

import (
	"context"
	"encoding/csv"
	"github.com/gocarina/gocsv"
	"github.com/pkg/errors"
	"github.com/vujanic79/golang-react-todo-app/pkg/internal/database"
	"github.com/vujanic79/golang-react-todo-app/pkg/logger"
	"io"
	"log"
	"os"
	"strings"
)

type TaskStatus struct {
	Status string `csv:"status"`
}

func LoadDataToDatabase(dbQueries *database.Queries, filePath string) {
	l := logger.Get()

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		l.Fatal().Stack().Err(errors.WithStack(err)).
			Msg("Reading CSV file error")
	}

	if len(rows) <= 1 {
		l.Info().Msg("CSV is empty or header-only")
		return
	}

	csvData := rows[1:] // Skip header

	taskStatuses, err := dbQueries.GetTaskStatuses(context.Background())
	if err != nil {
		l.Fatal().Stack().Err(errors.WithStack(err)).
			Msg("Getting task statuses error")
	}

	if len(csvData) <= len(taskStatuses) {
		l.Info().Msg("CSV data already loaded. Skipping insert.")
		return
	}

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		l.Fatal().Stack().Err(errors.WithStack(err)).Msg("Returning pointer to the start of the file error")
	}

	l.Info().Str("Ciro", "Ferrara").Msg("Loading task statuses")
	var entries []TaskStatus
	err = gocsv.Unmarshal(file, &entries)
	if err != nil {
		l.Fatal().Stack().Err(errors.WithStack(err)).
			Str("file", filePath).
			Msg("Unmarshalling CSV file error")
	}

	for _, entry := range entries {
		_, err := dbQueries.CreateTaskStatus(context.Background(), entry.Status)
		if err != nil && !strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			l.Fatal().Stack().Err(errors.WithStack(err)).
				Str("status", entry.Status).
				Msg("Creating task status error")
		}
	}
}
