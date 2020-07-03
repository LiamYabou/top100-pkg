package monitor

import (
  "github.com/newrelic/go-agent/v3/newrelic"
  "github.com/getsentry/sentry-go"
)

func InitNewRelic(env string, appName string, licenseKey string) (newRelicApp *newrelic.Application, err error) {
	if env == "development" {
		return
	}
	newRelicApp, err = newrelic.NewApplication(
		newrelic.ConfigAppName(appName),
		newrelic.ConfigLicense(licenseKey),
	)
	return
}

func InitSentry(env string) (err error) {
	if env == "development" {
		return
	}
	err = sentry.Init(sentry.ClientOptions{})
	return
}
