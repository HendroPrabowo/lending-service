package monitoring

import (
	"github.com/newrelic/go-agent/v3/newrelic"
	log "github.com/sirupsen/logrus"
)

var NewrelicApp *newrelic.Application

func InitNewRelic() {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("lending-service"),
		newrelic.ConfigLicense("9c4add1af0b4b4367fa33081aa2423e2FFFFNRAL"),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)
	if err != nil {
		log.Fatal(err)
	}
	NewrelicApp = app
	log.Info("newrelic connected")
}
