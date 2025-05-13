package util

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/vujanic79/golang-react-todo-app/pkg/logger"
	"net/http"
)

func RespondWithJson(w http.ResponseWriter, httpStatus int, response interface{}) {
	l := logger.Get()
	w.Header().Set("Content-Type", "application/json")

	b, err := json.Marshal(response)
	if err != nil {
		l.Error().Stack().Err(errors.WithStack(err)).
			Msg("Marshalling json response error")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(httpStatus)
	_, err = w.Write(b)
	if err != nil {
		l.Error().Stack().Err(errors.WithStack(err)).
			Dict("params", zerolog.Dict().
				Interface("response", response)).
			Msg("Writing json response error")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func RespondWithError(w http.ResponseWriter, httpStatus int, error string) {
	type errorResponse struct {
		Error string `json:"error"`
	}

	RespondWithJson(w, httpStatus, errorResponse{Error: error})
}
