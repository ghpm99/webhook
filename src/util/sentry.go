package util

import (
	"log"
	"time"
	"webhook/src/config"

	"github.com/getsentry/sentry-go"
)

func CaptureException(exception error, extra interface{}) error {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: config.SentryDSN,
	})
	if err != nil {
		log.Printf("sentry.Init: %s", err)
		return err
	}
	defer sentry.Flush(2 * time.Second)

	sentry.ConfigureScope(func(scope *sentry.Scope) {
		scope.SetTag("project", "Webhook")
		scope.SetExtra("extra", extra)

		sentry.CaptureException(exception)
	})

	return nil
}
