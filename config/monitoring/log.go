package monitoring

import (
	"github.com/sirupsen/logrus"
)

func InitLogger() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
}
