package monitor

import (
  "github.com/getsentry/sentry-go"
)

func InitSentry(env string) (err error) {
	if env == "development" {
		return
	}
	err = sentry.Init(sentry.ClientOptions{})
	return
}
