package logger

//logging setup

import "github.com/sirupsen/logrus"

var Log = logrus.New()

func Init() {
	Log.SetFormatter(&logrus.JSONFormatter{})
}
