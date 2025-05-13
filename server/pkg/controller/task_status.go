package controller

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/vujanic79/golang-react-todo-app/pkg/controller/util"
	"github.com/vujanic79/golang-react-todo-app/pkg/domain"
	"github.com/vujanic79/golang-react-todo-app/pkg/logger"
	"net/http"
)

type TaskStatusController struct {
	Tss domain.TaskStatusService
}

var _ domain.TaskStatusController = (*TaskStatusController)(nil)

func NewTaskStatusController(tss domain.TaskStatusService) (tsc TaskStatusController) {
	return TaskStatusController{Tss: tss}
}

func (tsc *TaskStatusController) CreateTaskStatus(w http.ResponseWriter, r *http.Request) {
	l := logger.Get()

	b, err := util.ReadBody(r)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, "Could not read user input")
		return
	}

	decoder := json.NewDecoder(r.Body)
	var params domain.CreateTaskStatusParams
	err = decoder.Decode(&params)
	if err != nil {
		l.Error().Stack().Err(errors.WithStack(err)).
			Str("url", r.URL.RequestURI()).
			Str("method", r.Method).
			Str("body", string(b)). // Raw string
			Msg("Creating task status error")
		util.RespondWithError(w, http.StatusBadRequest, "Parsing task status data from the body error")
		return
	}

	l = l.With().
		Dict("controller.params", zerolog.Dict().
			Str("func", "CreateTaskStatus").
			Dict("params", zerolog.Dict().
				Str("url", r.URL.RequestURI()).
				Str("method", r.Method).
				RawJSON("body", b))).
		Logger()
	ctx := logger.WithLogger(r.Context(), l)

	ts, err := tsc.Tss.CreateTaskStatus(ctx, params.Status)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, "Creating task status error")
		return
	}

	util.RespondWithJson(w, http.StatusCreated, ts)
}

func (tsc *TaskStatusController) GetTaskStatuses(w http.ResponseWriter, r *http.Request) {
	l := logger.Get()

	l = l.With().
		Dict("controller.params", zerolog.Dict().
			Str("func", "GetTaskStatuses").
			Dict("params", zerolog.Dict().
				Str("url", r.URL.RequestURI()).
				Str("method", r.Method))).
		Logger()
	ctx := logger.WithLogger(r.Context(), l)

	tss, err := tsc.Tss.GetTaskStatuses(ctx)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, "Getting task statuses error")
		return
	}

	util.RespondWithJson(w, http.StatusOK, tss)
}

func (tsc *TaskStatusController) GetTaskStatusByStatus(w http.ResponseWriter, r *http.Request) {
	l := logger.Get()
	status := chi.URLParam(r, "taskStatus")

	l = l.With().
		Dict("controller.params", zerolog.Dict().
			Str("func", "GetTaskStatusByStatus").
			Dict("params", zerolog.Dict().
				Str("url", r.URL.RequestURI()).
				Str("method", r.Method).
				Str("urlParam", status))).
		Logger()
	ctx := logger.WithLogger(r.Context(), l)

	ts, err := tsc.Tss.GetTaskStatusByStatus(ctx, status)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, "Getting task status error")
		return
	}

	util.RespondWithJson(w, http.StatusOK, ts)
}
