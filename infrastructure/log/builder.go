package log

import (
	"io"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type LogBuilder struct {
	e *logrus.Entry
}

const (
	reqIDField  = "req_id"
	sourceField = "source"
	errorField  = "error"
)

func newBuilder(ctx echo.Context) *LogBuilder {
	entry := logger.WithFields(logrus.Fields{"uri": ctx.Request().URL.String()})
	entry.Logger.SetFormatter(&logrus.JSONFormatter{})
	return &LogBuilder{entry}
}

func (b *LogBuilder) WithFields(flds map[string]any) *LogBuilder {
	b.e = b.e.WithFields(flds)
	return b
}

func (b *LogBuilder) WithSource(source string) *LogBuilder {
	b.e = b.e.WithField(sourceField, source)
	return b
}

func (b *LogBuilder) WithError(err error) *LogBuilder {
	b.e = b.e.WithField(errorField, err)
	return b
}

func (b *LogBuilder) WithWriter(w io.Writer) *LogBuilder {
	b.e.Logger.Out = w
	return b
}

func (b *LogBuilder) Now() *writer {
	now := time.Now()
	b.e.Time = now
	return &writer{b}
}

type writer struct{ b *LogBuilder }

func (w *writer) Infof(msg string, args ...any) {
	w.b.e.Infof(msg, args...)
}

func (w *writer) Errorf(msg string, args ...any) {
	w.b.e.Errorf(msg, args...)
}
