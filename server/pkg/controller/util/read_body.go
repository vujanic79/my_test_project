package util

import (
	"bytes"
	"github.com/pkg/errors"
	"github.com/vujanic79/golang-react-todo-app/pkg/logger"
	"io"
	"net/http"
)

func ReadBody(r *http.Request) (b []byte, err error) {
	l := logger.Get()

	b, err = io.ReadAll(r.Body)
	if err != nil {
		l.Error().Stack().Err(errors.WithStack(err)).
			Str("url", r.URL.String()).
			Str("method", r.Method).
			Msg("Error reading body")
		return nil, err
	}

	reader := io.NopCloser(bytes.NewBuffer(b))
	r.Body = reader
	return b, nil
}
