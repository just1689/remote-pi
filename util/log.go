package util

import "github.com/sirupsen/logrus"

func LogMsg(m string) {
	logrus.Info(">>> New message:")
	logrus.Info(m)

}
