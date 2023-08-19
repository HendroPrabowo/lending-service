package monitoring

import (
	"github.com/newrelic/go-agent/v3/integrations/logcontext-v2/nrlogrus"
	"github.com/sirupsen/logrus"
)

func InitLogger()  {
	nrlogrusFormatter := nrlogrus.NewFormatter(NewrelicApp, &logrus.JSONFormatter{})

	logrus.SetFormatter(nrlogrusFormatter)
}