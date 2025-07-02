package sentry

import "github.com/getsentry/sentry-go"

func InitSentry(dsn string, env string) error {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:           dsn,
		Environment:   env,
		EnableTracing: true,
		Debug:         true,
		EnableLogs:    true,
	})

	if err != nil {
		return err
	}

	return nil
}
