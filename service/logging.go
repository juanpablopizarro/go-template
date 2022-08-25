package service

import (
	"time"

	"github.com/go-kit/log"
)

type LoggingMiddleware struct {
	logger log.Logger
	next   Service
}

func NewLoggingMiddleware(logger log.Logger, next Service) LoggingMiddleware {
	return LoggingMiddleware{logger, next}
}

func (mw LoggingMiddleware) Hash(s string) (hash string, err error) {
	defer func(t time.Time) {
		mw.logger.Log(
			"method", "hash",
			"input", s,
			"err", err,
			"took", time.Since(t),
		)
	}(time.Now())
	hash, err = mw.next.Hash(s)
	return
}
