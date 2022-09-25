package logger

import "go.uber.org/zap"

type Logger = *zap.SugaredLogger

func New() (Logger, error) {
	l, err  := zap.NewProduction(
		zap.AddStacktrace(zap.ErrorLevel),
		zap.WithCaller(true),
	)
	if err != nil {
		return nil, err
	}
	return l.Sugar(), nil
}