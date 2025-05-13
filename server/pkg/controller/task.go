package controller

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/vujanic79/golang-react-todo-app/pkg/controller/util"
	"github.com/vujanic79/golang-react-todo-app/pkg/domain"
	"github.com/vujanic79/golang-react-todo-app/pkg/logger"
	"net/http"
)

type TaskController struct {
	Ts domain.TaskService
	Us domain.UserService
}

var _ domain.TaskController = (*TaskController)(nil)

func NewTaskController(ts domain.TaskService, us domain.UserService) (tc TaskController) {
	return TaskController{Ts: ts, Us: us}
}

func (tc *TaskController) CreateTask(
	w http.ResponseWriter,
	r *http.Request) {
	l := logger.Get()

	b, err := util.ReadBody(r)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, "Could not read user input")
		return
	}

	decoder := json.NewDecoder(r.Body)
	var params domain.CreateTaskParams
	err = decoder.Decode(&params)
	if err != nil {
		l.Error().Stack().Err(errors.WithStack(err)).
			Str("url", r.URL.RequestURI()).
			Str("method", r.Method).
			Str("body", string(b)). // Raw string
			Msg("Creating task error")
		util.RespondWithError(w, http.StatusBadRequest, "Parsing task data from the body error")
		return
	}

	l = l.With().
		Dict("controller.params", zerolog.Dict().
			Str("func", "CreateTask").
			Dict("params", zerolog.Dict().
				Str("url", r.URL.RequestURI()).
				Str("method", r.Method).
				RawJSON("body", b))).
		Logger()
	ctx := logger.WithLogger(r.Context(), l)

	userId, err := tc.Us.GetUserIdByEmail(ctx, params.UserEmail)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, "Could not complete the request")
		return
	}

	t, err := tc.Ts.CreateTask(ctx, userId, params)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, "Creating task error")
		return
	}

	util.RespondWithJson(w, http.StatusCreated, t)
}

func (tc *TaskController) DeleteTask(w http.ResponseWriter, r *http.Request) {
	l := logger.Get()

	idStr := chi.URLParam(r, "taskId")
	id, err := uuid.Parse(idStr)
	if err != nil {
		l.Error().Stack().Err(errors.WithStack(err)).
			Str("id", idStr).
			Msg("Parsing id error")
		util.RespondWithError(w, http.StatusBadRequest, "Invalid task id")
		return
	}

	l = l.With().
		Dict("controller.params", zerolog.Dict().
			Str("func", "DeleteTask").
			Dict("params", zerolog.Dict().
				Str("url", r.URL.RequestURI()).
				Str("method", r.Method).
				Str("urlParam", idStr))).
		Logger()
	ctx := logger.WithLogger(r.Context(), l)

	err = tc.Ts.DeleteTask(ctx, id)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, "Deleting task error")
		return
	}
	util.RespondWithJson(w, http.StatusOK, fmt.Sprintf("Task with id %s successfully deleted", id))
}

func (tc *TaskController) UpdateTask(w http.ResponseWriter, r *http.Request) {
	l := logger.Get()

	idStr := chi.URLParam(r, "taskId")
	id, err := uuid.Parse(idStr)
	if err != nil {
		l.Error().Stack().Err(errors.WithStack(err)).
			Str("id", idStr).
			Msg("Parsing id error")
		util.RespondWithError(w, http.StatusBadRequest, "Invalid task id")
		return
	}

	b, err := util.ReadBody(r)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, "Could not read user input")
		return
	}

	decoder := json.NewDecoder(r.Body)
	var params domain.UpdateTaskParams
	err = decoder.Decode(&params)
	if err != nil {
		l.Error().Stack().Err(errors.WithStack(err)).
			Str("url", r.URL.RequestURI()).
			Str("method", r.Method).
			Str("body", string(b)). // Raw string
			Msg("Updating task error")
		util.RespondWithError(w, http.StatusBadRequest, "Parsing task data from the body error")
		return
	}

	l = l.With().
		Dict("controller.params", zerolog.Dict().
			Str("func", "UpdateTask").
			Dict("params", zerolog.Dict().
				Str("url", r.URL.RequestURI()).
				Str("method", r.Method).
				RawJSON("body", b))).
		Logger()
	ctx := logger.WithLogger(r.Context(), l)

	params.ID = id
	t, err := tc.Ts.UpdateTask(ctx, params)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, "Updating task error")
		return
	}

	util.RespondWithJson(w, http.StatusOK, t)
}

func (tc *TaskController) GetTasksByUserId(w http.ResponseWriter, r *http.Request) {
	l := logger.Get()

	b, err := util.ReadBody(r)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, "Could not read user input")
		return
	}

	decoder := json.NewDecoder(r.Body)
	var params domain.GetTasksByUserIdParams
	err = decoder.Decode(&params)
	if err != nil {
		l.Error().Stack().Err(errors.WithStack(err)).
			Str("url", r.URL.RequestURI()).
			Str("method", r.Method).
			Str("body", string(b)). // Raw string
			Msg("Getting task by userId error")
		util.RespondWithError(w, http.StatusBadRequest, "Parsing task data from the body error")
		return
	}

	l = l.With().
		Dict("controller.GetTasksByUserId_params", zerolog.Dict().
			Str("func", "GetTasksByUserId").
			Dict("params", zerolog.Dict().
				Str("url", r.URL.RequestURI()).
				Str("method", r.Method).
				RawJSON("body", b))).
		Logger()
	ctx := logger.WithLogger(r.Context(), l)

	ts, err := tc.Ts.GetTasksByUserId(ctx, params.UserID)
	if err != nil {
		util.RespondWithError(w, http.StatusInternalServerError, "Getting tasks error")
		return
	}

	util.RespondWithJson(w, http.StatusOK, ts)
}
