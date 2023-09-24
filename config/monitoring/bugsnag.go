package monitoring

import (
	"os"

	"github.com/bugsnag/bugsnag-go/v2"
	log "github.com/sirupsen/logrus"
)

func InitBugsnag() {
	apiKey := os.Getenv("BUGSNAG_API_KEY")
	if apiKey == "" {
		apiKey = "57489b4e6b66d065120271d6643449cc"
	}

	releaseStage := os.Getenv("BUGSNAG_RELEASE_STAGE")
	if releaseStage == "" {
		releaseStage = "development"
	}

	bugsnag.Configure(bugsnag.Configuration{
		APIKey:       releaseStage,
		ReleaseStage: "development",
	})

	log.Info("bugsnag connected")
}
